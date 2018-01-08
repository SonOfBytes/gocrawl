package main

import (
	"log"

	authClient "github.com/sonofbytes/gocrawl/authentication/client"
	queueClient "github.com/sonofbytes/gocrawl/queue/client"
	"github.com/sonofbytes/gocrawl/retriever/server"
	storeClient "github.com/sonofbytes/gocrawl/store/client"
)

func main() {
	log.Println("retriever started")

	s := server.New(authClient.New(), queueClient.New(), storeClient.New())
	s.Wait()
}
