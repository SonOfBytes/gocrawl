package manager

import (
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DB struct {
	sm sync.Map
}

func New() *DB {
	return &DB{}
}

func (db *DB) Get(url string) ([]string, error) {
	if url == "" {
		return []string{}, status.Error(codes.InvalidArgument, "url needs to be set")
	}

	v, ok := db.sm.Load(url)
	if !ok {
		return []string{}, status.Error(codes.NotFound, "submitted URL does not exist")
	}

	return v.([]string), nil
}

func (db *DB) Add(url string, urls []string) error {
	if url == "" {
		return status.Error(codes.InvalidArgument, "url needs to be set")
	}

	db.sm.Store(url, urls)

	return nil
}
