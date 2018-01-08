package sessionmgr_test

import (
	"github.com/sonofbytes/gocrawl/authentication/sessionmgr"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	s := sessionmgr.New()
	assert.IsType(t, &sessionmgr.Sessions{}, s)
}

func TestSessions_Add(t *testing.T) {
	s := sessionmgr.New()
	ns := s.Add()
	assert.Len(t, ns, 12)
	assert.True(t, s.Check(ns))
}

func TestSessions_Delete(t *testing.T) {
	s := sessionmgr.New()
	ns := s.Add()
	s.Delete(ns)
	assert.False(t, s.Check(ns))
}

func TestExpire(t *testing.T) {
	s := sessionmgr.New()
	s.SetExpire(time.Millisecond)
	ns := s.Add()
	assert.True(t, s.Check(ns))
	time.Sleep(time.Millisecond)
	assert.False(t, s.Check(ns))
}
