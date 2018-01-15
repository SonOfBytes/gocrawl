package server

import (
	"log"
	"net/http"

	"io"

	"fmt"

	"net/url"

	"context"

	"strings"

	"golang.org/x/net/html"
)

const maxDepth = 20

// process will be launched as a go routine so must be self sufficient
func (s *Server) process(url string, depth int, job string) {
	if depth > maxDepth {
		log.Printf("Depth exceeded for http.Get [%d]: %s", depth, url)
		return
	}

	log.Printf("http.Get [%d]: %s", depth, url)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("http.Get Error (%s): %s", url, err)
		// do something based on the error?
		// 404 discard
		// 500 queue again?
		return
	}
	defer resp.Body.Close()

	links, err := parseHTLMReaderForLinks(url, resp.Body)
	if err != nil {
		// handle errors based on business need
		return
	}

	session := s.getSession()
	// Store the links for later retrieval
	err = s.storeService.Submit(context.Background(), session, url, links)
	if err != nil {
		log.Printf("store URL crawl results error: %s", err)
		return
	}

	for _, link := range links {
		err := s.queueService.Submit(context.Background(), session, link, depth+1, job)
		if err != nil {
			//log.Printf("Submit Notice [%s]: %s", link, err)
			continue
		}
	}
}

func parseHTLMReaderForLinks(url string, rc io.ReadCloser) ([]string, error) {
	links := make([]string, 0)

	doc, err := html.Parse(rc)
	if err != nil {
		// need to understand what the business need on error is in order to handle correctly
		log.Printf("http.Parse Error (%s): %s", url, err)
		return links, err
	}

	var f func(*html.Node, string, *[]string)
	f = func(n *html.Node, url string, links *[]string) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, v := range n.Attr {
				if v.Key == "href" {
					fixurl, err := fixRelativeURL(url, v.Val)
					if err != nil {
						log.Printf("fix url error (%s + %s): %s", url, v.Val, err)
						continue
					}
					ok, err := urlInDomain(url, fixurl)
					if err != nil {
						log.Printf("url in domain error (%s + %s): %s", url, v.Val, err)
					}
					if !ok {
						continue
					}
					*links = append(*links, fixurl)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, url, links)
		}
	}
	f(doc, url, &links)

	// create unique links
	u := make([]string, 0)
	m := make(map[string]int)
	for _, v := range links {
		if _, ok := m[v]; !ok {
			u = append(u, v)
		}
		m[v] = 1
	}
	return u, nil
}

func fixRelativeURL(baseurl, parseurl string) (string, error) {
	base, err := url.Parse(baseurl)
	if err != nil {
		return "", fmt.Errorf("Parse base url Error : %s", err)
	}
	u, err := url.Parse(parseurl)
	if err != nil {
		return "", fmt.Errorf("Parse url Error : %s", err)
	}
	return base.ResolveReference(u).String(), nil
}

func urlInDomain(baseurl, parseurl string) (bool, error) {
	// assumes absolute urls
	b, err := url.Parse(baseurl)
	if err != nil {
		return false, fmt.Errorf("Parse base url Error : %s", err)
	}
	u, err := url.Parse(parseurl)
	if err != nil {
		return false, fmt.Errorf("Parse url Error : %s", err)
	}

	return baseDomain(b.Host) == baseDomain(u.Host), nil
}

func baseDomain(h string) string {
	h = strings.TrimSpace(h)
	hp := strings.Split(h, ".")
	if len(hp) < 2 {
		return "<<Invalid>>"
	}
	return strings.Join(hp[len(hp)-2:], ".")
}
