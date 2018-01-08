package main

import (
	"context"
	"fmt"
	"log"
	"os"

	apiclient "github.com/sonofbytes/gocrawl/api/client"
	authclient "github.com/sonofbytes/gocrawl/authentication/client"
	storeclient "github.com/sonofbytes/gocrawl/store/client"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app      = kingpin.New("chat", "A command-line chat application.")
	serverIP = app.Flag("server", "Server address.").Default("0.0.0.0").IP()

	authenticate         = app.Command("authenticate", "Authenticate session")
	authenticateUsername = authenticate.Arg("username", "User name.").Required().String()
	authenticatePassword = authenticate.Arg("password", "Password of user.").Required().String()

	submit        = app.Command("submit", "Submit a url crawl request.")
	submitSession = submit.Arg("session", "Authentication session.").Required().String()
	submitURL     = submit.Arg("url", "Url to crawl.").Required().String()

	get        = app.Command("get", "Get an url's crawl results.")
	getSession = get.Arg("session", "Authentication session.").Required().String()
	getURL     = get.Arg("url", "Url to get.").Required().String()
)

func main() {

	command := kingpin.MustParse(app.Parse(os.Args[1:]))

	switch command {
	// Authenticate user
	case authenticate.FullCommand():
		ac := authclient.New()
		ac.SetConnection((*serverIP).String())
		r, err := ac.Authenticate(*authenticateUsername, *authenticatePassword)
		if err != nil {
			log.Fatalf("could not authentication: %s", err)
		}
		fmt.Printf("%s\n", r)

		// Submit URL
	case submit.FullCommand():
		sc := apiclient.New()
		sc.SetConnection((*serverIP).String())
		r, err := sc.Submit(context.Background(), *submitSession, *submitURL)
		if err != nil {
			log.Fatalf("submit failure: %v", err)
		}
		fmt.Printf("Job: %s\n", r)

		// Get URL
	case get.FullCommand():
		stc := storeclient.New()
		stc.SetConnection((*serverIP).String())
		r, err := stc.Get(context.Background(), *getSession, *getURL)
		if err != nil {
			log.Fatalf("get failure: %v", err)
		}
		for _, v := range r {
			fmt.Printf("%s\n", v)
		}
	}
}
