package main

import (
	"log"

	"github.com/sonofbytes/gocrawl/api/server"
	authClient "github.com/sonofbytes/gocrawl/authentication/client"
	queueClient "github.com/sonofbytes/gocrawl/queue/client"
	storeClient "github.com/sonofbytes/gocrawl/store/client"
)

const (
	address = ":8888"
)

func main() {
	log.Println("api started")

	s := server.New(authClient.New(), queueClient.New(), storeClient.New())

	if err := s.Serve(address); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
