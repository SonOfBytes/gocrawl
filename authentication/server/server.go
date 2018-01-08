package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sonofbytes/gocrawl/authentication/pb"
	"github.com/sonofbytes/gocrawl/authentication/sessionmgr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	username = "someone"
	password = "hardcoded"
)

type Server struct {
	server   *grpc.Server
	sessions *sessionmgr.Sessions
}

func New() *Server {
	s := &Server{
		sessions: sessionmgr.New(),
	}

	s.server = grpc.NewServer()
	authpb.RegisterAuthenticationServer(s.server, s)
	// Register reflection service on gRPC server.
	reflection.Register(s.server)

	return s
}

func (s *Server) Serve(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return s.server.Serve(listener)
}

func (s *Server) Authenticate(ctx context.Context, ar *authpb.AuthenticateRequest) (*authpb.AuthenticateReply, error) {
	// This is a simple session server to provide nominal protection in PoC
	if ar == nil || ar.Username == "" || ar.Password == "" {
		return &authpb.AuthenticateReply{}, fmt.Errorf("authentication request must have a valid Username and Password")
	}

	if ar.Username == username && ar.Password == password {
		ns := s.sessions.Add()
		return &authpb.AuthenticateReply{
			Session: ns,
		}, nil
	}
	return &authpb.AuthenticateReply{}, fmt.Errorf("invalid Username or Password")
}

func (s *Server) Validate(ctx context.Context, vr *authpb.ValidateRequest) (*authpb.ValidateReply, error) {
	if vr == nil || vr.Session == "" {
		return &authpb.ValidateReply{}, fmt.Errorf("Validation request requires a Session token to be set")
	}
	return &authpb.ValidateReply{
		Valid: s.sessions.Check(vr.Session),
	}, nil
}
