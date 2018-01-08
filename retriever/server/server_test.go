package server_test

import (
	"testing"

	"net"

	"time"

	"errors"

	"github.com/golang/mock/gomock"
	"github.com/sonofbytes/gocrawl/retriever/mocks"
	"github.com/sonofbytes/gocrawl/retriever/server"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mas := mocks.NewMockAuthService(mockCtrl)
	mqs := mocks.NewMockQueueService(mockCtrl)
	mss := mocks.NewMockStoreService(mockCtrl)

	za := &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 9999,
	}

	mas.EXPECT().SetConnection(za.IP.String()).Return(nil).AnyTimes()
	mqs.EXPECT().SetConnection(za.IP.String()).Return(nil).AnyTimes()
	mss.EXPECT().SetConnection(za.IP.String()).Return(nil).AnyTimes()
	// authenticate isn't valid to stop go routines processing in tests
	mas.EXPECT().Authenticate("someone", "hardcoded").Return("", errors.New("!")).AnyTimes()

	s := server.New(mas, mqs, mss)
	assert.IsType(t, &server.Server{}, s)

	time.Sleep(time.Millisecond)
}

// the retriever needs reworking to improve injection of mocks for tests (eg http get)
// as well as improving testing of go routines

// an exercise for the future ;)
