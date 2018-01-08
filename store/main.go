package main

import (
	"log"

	authClient "github.com/sonofbytes/gocrawl/authentication/client"
	"github.com/sonofbytes/gocrawl/store/server"
)

const (
	address = ":8888"
)

func main() {
	log.Println("store started")

	s := server.New(authClient.New())
	if err := s.Serve(address); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
