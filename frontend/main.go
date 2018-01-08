package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"

	"time"

	"html"
	"net/url"

	"fmt"

	"strings"

	apiclient "github.com/sonofbytes/gocrawl/api/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DefaultPort is the default port to use if once is not specified by the SERVER_PORT environment variable
const DefaultPort = "8888"
const cookieSessionName = "gocrawl-session"
const basePath = "/"

type AuthDetails struct {
	Username string
	Message  string
}

type IndexDetails struct {
	Authenticated bool
	Processing    bool
	BaseURL       string
	URL           string
	Message       string
	URLs          []string
}

var apiClient *apiclient.Client

func getServerPort() string {
	port := os.Getenv("SERVER_PORT")
	if port != "" {
		return port
	}

	return DefaultPort
}

// LoginHandler authenticates the session
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("frontend/login.html"))

	if r.Method != http.MethodPost {
		t.Execute(w, nil)
		log.Printf("Not a POST - present login form")
		return
	}

	post := &AuthDetails{
		Username: r.FormValue("username"),
	}

	apiClient.SetConnection(strings.Split(r.RemoteAddr, ":")[0])

	session, err := apiClient.Authenticate(r.FormValue("username"), r.FormValue("password"))
	if err != nil {
		post.Message = err.Error()
		t.Execute(w, post)
		log.Printf("Authenticate failed - redo login form: %s", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    cookieSessionName,
		Value:   session,
		Expires: time.Now().Add(time.Minute * 5),
	})

	log.Printf("Validated - redirect /")
	http.Redirect(w, r, basePath, 302)
}

// FrontendHandler provides an html interface to the store service
func FrontendHandler(w http.ResponseWriter, r *http.Request) {

	index := &IndexDetails{
		Processing: false,
	}

	t := template.Must(template.ParseFiles("frontend/index.html"))

	// check authentication
	sessionCookie, _ := r.Cookie(cookieSessionName)
	if sessionCookie == nil {
		log.Printf("No session - redirect /login")
		http.Redirect(w, r, basePath+"login", 302)
		return
	}

	addr := strings.Split(r.RemoteAddr, ":")[0]
	apiClient.SetConnection(addr)

	verr := apiClient.Validate(sessionCookie.Value)
	if verr != nil {
		http.SetCookie(w, &http.Cookie{
			Name:   cookieSessionName,
			Value:  "",
			MaxAge: -1,
		})
		log.Printf("Validate failed - redirect /login")
		http.Redirect(w, r, basePath+"login", 302)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    cookieSessionName,
		Value:   sessionCookie.Value,
		Expires: time.Now().Add(time.Minute * 5),
	})
	index.Authenticated = true

	reqUrl, err := formUrl(r, "url")
	if err != nil {
		index.Message = html.EscapeString(err.Error())
		t.Execute(w, index)
		log.Printf(err.Error())
		return
	}

	if reqUrl == "" {
		t.Execute(w, index)
		return
	}

	baseUrl, err := formUrl(r, "base")
	if err != nil {
		index.Message = html.EscapeString(err.Error())
		t.Execute(w, index)
		log.Printf(err.Error())
		return
	}

	if r.Method == http.MethodPost {
		setQuery(r, "base", reqUrl)
		http.Redirect(w, r, r.URL.String(), 302)
		return
	}

	if baseUrl == "" {
		setQuery(r, "base", reqUrl)
		http.Redirect(w, r, r.URL.String(), 302)
		return
	}

	index.BaseURL = baseUrl

	urls, err := apiClient.Get(context.Background(), sessionCookie.Value, reqUrl)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			if s.Code() == codes.NotFound {
				index.Processing = true
				apiClient.Submit(context.Background(), sessionCookie.Value, reqUrl)
			}
		}
		t.Execute(w, index)
		log.Printf("Store Get error for [%s]- %s", reqUrl, err)
		return
	}
	index.URLs = urls
	t.Execute(w, index)
}

func main() {

	log.Println("starting frontend server, listening on port " + getServerPort())

	apiClient = apiclient.New()

	http.HandleFunc(basePath+"login", LoginHandler)
	http.HandleFunc(basePath, FrontendHandler)
	http.ListenAndServe(":"+getServerPort(), nil)
}

func formUrl(r *http.Request, key string) (string, error) {
	v := r.FormValue(key)
	if v == "" {
		return v, nil
	}

	if _, err := url.ParseRequestURI(v); err != nil {
		return "", fmt.Errorf("Invalid URL - " + v)
	}

	setQuery(r, key, v)

	return v, nil
}

func setQuery(r *http.Request, key string, value string) {
	q := r.URL.Query()
	q.Del(key)
	q.Add(key, value)
	r.URL.RawQuery = q.Encode()
}
