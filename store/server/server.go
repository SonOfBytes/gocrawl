package server

import (
	"context"
	"fmt"

	"github.com/sonofbytes/gocrawl/store/pb"

	"log"
	"net"

	"github.com/sonofbytes/gocrawl/store/manager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
)

// AuthService is an interface to meet the authentication needs
type AuthService interface {
	SetConnection(host string) error
	Validate(session string) error
}

// Server satisfies gRPC service interface requirements Store
type Server struct {
	authService AuthService
	server      *grpc.Server
	store       *manager.DB
}

// New creates a gRPC server instance for Store
func New(authSerice AuthService) *Server {
	s := &Server{
		authService: authSerice,
		store:       manager.New(),
	}

	s.server = grpc.NewServer()
	storepb.RegisterStoreServer(s.server, s)
	// Register reflection service on gRPC server.
	reflection.Register(s.server)

	return s
}

// Serve handles gRPC requests for service Store
func (s *Server) Serve(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return s.server.Serve(listener)
}

// SubmitURLs handles client requests to Submit to the Store service
func (s *Server) Submit(ctx context.Context, sr *storepb.StoreSubmitRequest) (*storepb.StoreSubmitReply, error) {
	// This is a simple session server to provide nominal protection in PoC
	if sr == nil || sr.Session == "" || sr.Url == "" {
		return &storepb.StoreSubmitReply{}, fmt.Errorf("request needs Session token and Url")
	}

	if p, ok := peer.FromContext(ctx); ok {
		if addr, ok := p.Addr.(*net.TCPAddr); ok {
			s.authService.SetConnection(addr.IP.String())
		}
	}

	err := s.authService.Validate(sr.Session)
	if err != nil {
		return &storepb.StoreSubmitReply{}, err
	}

	err = s.store.Add(sr.Url, sr.Urls)
	if err != nil {
		return &storepb.StoreSubmitReply{}, err
	}

	return &storepb.StoreSubmitReply{}, nil
}

// GetURLs retrieves an available URL's associated link urls
func (s *Server) Get(ctx context.Context, gr *storepb.StoreGetRequest) (*storepb.StoreGetReply, error) {
	if gr == nil || gr.Session == "" || gr.Url == "" {
		return &storepb.StoreGetReply{}, fmt.Errorf("request needs Session token and Url")
	}

	if p, ok := peer.FromContext(ctx); ok {
		if addr, ok := p.Addr.(*net.TCPAddr); ok {
			s.authService.SetConnection(addr.IP.String())
		}
	}

	err := s.authService.Validate(gr.Session)
	if err != nil {
		return &storepb.StoreGetReply{}, err
	}

	urls, err := s.store.Get(gr.Url)
	return &storepb.StoreGetReply{
		Urls: urls,
	}, err
}
