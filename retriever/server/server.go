package server

import (
	"sync"
	"time"

	"log"

	"context"

	"net"

	"github.com/remeh/sizedwaitgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const maxConcurrentGet = 4

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
	authService  AuthService
	queueService QueueService
	storeService StoreService
	session      string
	lock         sync.Mutex
	wg           sync.WaitGroup
}

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

	// keeps a valid session active (assumes sessions longer than a minute)
	go func() {
		for {
			var err error
			session, err := s.authService.Authenticate("someone", "hardcoded")
			s.setSession(session)
			if err == nil {
				time.Sleep(time.Minute)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	s.wg.Add(1)
	go s.Run()
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
	return s.session
}

// Wait waits until the waitgroup is complete
func (s *Server) Wait() {
	s.wg.Wait()
}

// Run loops and processes the queue
func (s *Server) Run() {
	defer s.wg.Done()
	// wait until authentication session is valid
	for s.getSession() == "" {
		time.Sleep(time.Millisecond)
	}
	// use sized wait group to throttle process concurrency
	swg := sizedwaitgroup.New(maxConcurrentGet)

	// process queue
	for {
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
		swg.Add()
		go s.process(&swg, url, depth, job)
	}
}
