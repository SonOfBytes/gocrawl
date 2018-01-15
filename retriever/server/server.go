package server

import (
	"sync"
	"time"

	"log"

	"context"

	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthService is an interface to meet the authentication needs
type AuthService interface {
	Authenticate(Username string, Password string) (session string, err error)
	SetConnection(host string) error
}

// QueueService is an interface to meet the queue needs
type QueueService interface {
	Submit(ctx context.Context, session string, url string, depth int, job string) error
	Get(ctx context.Context, session string) (url string, depth int, job string, err error)
	SetConnection(host string) error
}

// StoreService is an interface to meet the store needs
type StoreService interface {
	Submit(ctx context.Context, session string, url string, urls []string) error
	SetConnection(host string) error
}

// Server performs the retrieval from Queue and gets the URL contents
type Server struct {
	authService   AuthService
	queueService  QueueService
	storeService  StoreService
	session       string
	sessionExpire time.Time
	lock          sync.Mutex
	wg            sync.WaitGroup
}

// DoOnfLoop will loop forever while true - false when testing or cancelled
var DoInfLoop = true

// New creates a gRPC server instance for Retriever
func New(authService AuthService, queueService QueueService, storeService StoreService) *Server {
	s := &Server{
		authService:  authService,
		queueService: queueService,
		storeService: storeService,
	}

	addr := "127.0.0.1"
	addrs, err := net.LookupHost("linkerd")
	if err == nil && len(addrs) > 0 {
		addr = addrs[0]
	}

	s.authService.SetConnection(addr)
	s.queueService.SetConnection(addr)
	s.storeService.SetConnection(addr)

	return s
}

func (s *Server) setSession(session string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.session = session
}

func (s *Server) getSession() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	for s.session == "" || s.sessionExpire.Before(time.Now()) {
		session, err := s.authService.Authenticate("someone", "hardcoded")
		if err != nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		s.session = session
		s.sessionExpire = time.Now().Add(time.Minute)
	}
	return s.session
}

// Run loops and processes the queue
func (s *Server) Run(process ...func(url string, depth int, job string)) {
	// set a default process if not defined in the Run call
	if process == nil {
		process = append(process, s.process)
	}

	var loop = true
	// process queue while loop is true
	for loop {
		url, depth, job, err := s.queueService.Get(context.Background(), s.getSession())
		if err != nil {
			if s, ok := status.FromError(err); ok {
				if s.Code() == codes.NotFound || s.Code() == codes.Unavailable {
					time.Sleep(100 * time.Millisecond)
					continue
				}
			}
			log.Printf("failed Get: %s", err)
			time.Sleep(time.Second)
		}
		for _, p := range process {
			p(url, depth, job)
		}
		loop = DoInfLoop
	}
}
