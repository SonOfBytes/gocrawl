package sessionmgr

import (
	"math/rand"
	"sync"
	"time"
)

const expireDuration = time.Hour

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Sessions struct {
	expireDuration time.Duration
	sm             sync.Map
}

func init() {
	// ensure our session ids are seeded randomly-ish
	rand.Seed(time.Now().UTC().UnixNano())
}

func New() *Sessions {
	return &Sessions{
		expireDuration: expireDuration,
	}
}

func (s *Sessions) SetExpire(duration time.Duration) {
	s.expireDuration = duration
}

func (s *Sessions) Check(session string) bool {
	if v, ok := s.sm.Load(session); ok {
		expire := v.(time.Time)
		if expire.After(time.Now()) {
			return true
		}
		s.Delete(session)
	}
	return false
}

func (s *Sessions) Delete(session string) {
	s.sm.Delete(session)
}

func (s *Sessions) Add() string {
	session := randSeq(12)
	s.sm.Store(session, time.Now().Add(s.expireDuration))
	return session
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
