package server

import (
	"context"

	"github.com/sonofbytes/gocrawl/api/pb"

	"log"
	"net"

	"math/rand"

	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// AuthService is an interface to meet the authentication needs
type AuthService interface {
	Authenticate(username, password string) (session string, err error)
	Validate(session string) error
	SetConnection(host string) error
}

// QueueService is an interface to meet the queue needs
type QueueService interface {
	Submit(ctx context.Context, session string, url string, depth int, job string) error
	SetConnection(host string) error
}

// StoreService is an interface to meet the store needs
type StoreService interface {
	Get(ctx context.Context, session string, url string) (urls []string, err error)
	SetConnection(host string) error
}

// Server satisfies gRPC service interface requirements API
type Server struct {
	server       *grpc.Server
	authService  AuthService
	queueService QueueService
	storeService StoreService
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	// ensure our session ids are seeded randomly-ish
	rand.Seed(time.Now().UTC().UnixNano())
}

// New creates a gRPC server instance for API
func New(authService AuthService, queueService QueueService, storeService StoreService) *Server {
	s := &Server{
		authService:  authService,
		queueService: queueService,
		storeService: storeService,
	}

	s.server = grpc.NewServer()
	apipb.RegisterAPIServer(s.server, s)
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

// Submit handles client requests to Submit to the API service
func (s *Server) Submit(ctx context.Context, ar *apipb.APISubmitRequest) (*apipb.APISubmitReply, error) {
	if ar == nil || ar.Session == "" || ar.Url == "" {
		return &apipb.APISubmitReply{}, status.Error(codes.InvalidArgument, "request needs Session token and Url")
	}

	if p, ok := peer.FromContext(ctx); ok {
		if addr, ok := p.Addr.(*net.TCPAddr); ok {
			s.authService.SetConnection(addr.IP.String())
			s.queueService.SetConnection(addr.IP.String())
		}
	}

	err := s.authService.Validate(ar.Session)
	if err != nil {
		return &apipb.APISubmitReply{}, err
	}

	job := randSeq(12)
	err = s.queueService.Submit(ctx, ar.Session, ar.Url, 0, job)
	if err != nil {
		return &apipb.APISubmitReply{}, err
	}

	return &apipb.APISubmitReply{
		Job: job,
	}, nil
}

func (s *Server) Get(ctx context.Context, ar *apipb.APIGetRequest) (*apipb.APIGetReply, error) {
	if ar == nil || ar.Session == "" || ar.Url == "" {
		return &apipb.APIGetReply{}, status.Error(codes.InvalidArgument, "request needs Session token and Url")
	}

	if p, ok := peer.FromContext(ctx); ok {
		if addr, ok := p.Addr.(*net.TCPAddr); ok {
			s.authService.SetConnection(addr.IP.String())
			s.storeService.SetConnection(addr.IP.String())
		}
	}

	err := s.authService.Validate(ar.Session)
	if err != nil {
		return &apipb.APIGetReply{}, err
	}

	urls, err := s.storeService.Get(ctx, ar.Session, ar.Url)

	return &apipb.APIGetReply{
		Urls: urls,
	}, err
}

func (s *Server) Authenticate(ctx context.Context, ar *apipb.APIAuthenticateRequest) (*apipb.APIAuthenticateReply, error) {
	if ar == nil || ar.Username == "" || ar.Password == "" {
		return &apipb.APIAuthenticateReply{}, status.Error(codes.InvalidArgument, "request needs Username and Password")
	}

	if p, ok := peer.FromContext(ctx); ok {
		if addr, ok := p.Addr.(*net.TCPAddr); ok {
			s.authService.SetConnection(addr.IP.String())
		}
	}

	session, err := s.authService.Authenticate(ar.Username, ar.Password)

	return &apipb.APIAuthenticateReply{
		Session: session,
	}, err
}

func (s *Server) Validate(ctx context.Context, ar *apipb.APIValidateRequest) (*apipb.APIValidateReply, error) {
	if ar == nil || ar.Session == "" {
		return &apipb.APIValidateReply{}, status.Error(codes.InvalidArgument, "request needs Username and Password")
	}

	if p, ok := peer.FromContext(ctx); ok {
		if addr, ok := p.Addr.(*net.TCPAddr); ok {
			s.authService.SetConnection(addr.IP.String())
		}
	}

	err := s.authService.Validate(ar.Session)

	return &apipb.APIValidateReply{
		Valid: err == nil,
	}, err
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
