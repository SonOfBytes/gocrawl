package main

import (
	"log"

	"github.com/sonofbytes/gocrawl/authentication/server"
)

const (
	address = ":8888"
)

func main() {
	log.Println("authentication started")

	s := server.New()

	if err := s.Serve(address); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
