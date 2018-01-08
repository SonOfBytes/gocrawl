package server

import (
	"context"
	"fmt"

	"github.com/sonofbytes/gocrawl/queue/pb"

	"log"
	"net"

	"github.com/sonofbytes/gocrawl/queue/manager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
)

// AuthService is an interface to meet the authentication needs
type AuthService interface {
	Validate(session string) error
	SetConnection(host string) error
}

// Server satisfies gRPC service interface requirements Queue
type Server struct {
	server      *grpc.Server
	authService AuthService
	queue       *manager.DB
}

// New creates a gRPC server instance for API
func New(authService AuthService) *Server {
	s := &Server{
		authService: authService,
		queue:       manager.New(),
	}

	s.server = grpc.NewServer()
	queuepb.RegisterQueueServer(s.server, s)
	// Register reflection service on gRPC server.
	reflection.Register(s.server)

	return s
}

// Serve handles gRPC requests for service API
func (s *Server) Serve(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return s.server.Serve(listener)
}

// Submit handles client requests to Submit to the Queue service
func (s *Server) Submit(ctx context.Context, sr *queuepb.QueueSubmitRequest) (*queuepb.QueueSubmitReply, error) {
	// This is a simple session server to provide nominal protection in PoC
	if sr == nil || sr.Session == "" || sr.Url == "" {
		return &queuepb.QueueSubmitReply{}, fmt.Errorf("request needs Session token and Url")
	}

	if p, ok := peer.FromContext(ctx); ok {
		if addr, ok := p.Addr.(*net.TCPAddr); ok {
			s.authService.SetConnection(addr.IP.String())
		}
	}

	err := s.authService.Validate(sr.Session)
	if err != nil {
		return &queuepb.QueueSubmitReply{}, err
	}

	err = s.queue.Add(sr.Url, sr.Depth, sr.Job)
	if err != nil {
		return &queuepb.QueueSubmitReply{}, err
	}

	return &queuepb.QueueSubmitReply{}, nil
}

// Get retrieves an available URL and recursion depth for processing
func (s *Server) Get(ctx context.Context, gr *queuepb.QueueGetRequest) (*queuepb.QueueGetReply, error) {
	if gr == nil || gr.Session == "" {
		return &queuepb.QueueGetReply{}, fmt.Errorf("request needs Session token")
	}

	if p, ok := peer.FromContext(ctx); ok {
		if addr, ok := p.Addr.(*net.TCPAddr); ok {
			s.authService.SetConnection(addr.IP.String())
		}
	}

	err := s.authService.Validate(gr.Session)
	if err != nil {
		return &queuepb.QueueGetReply{}, err
	}

	url, depth, job, err := s.queue.Get()
	return &queuepb.QueueGetReply{
		Url:   url,
		Depth: depth,
		Job:   job,
	}, err
}
