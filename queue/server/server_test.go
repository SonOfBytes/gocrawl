package server_test

import (
	"testing"

	"context"

	"fmt"

	"net"

	"github.com/golang/mock/gomock"
	"github.com/sonofbytes/gocrawl/queue/mocks"
	"github.com/sonofbytes/gocrawl/queue/pb"
	"github.com/sonofbytes/gocrawl/queue/server"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/peer"
)

func TestNew(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mas := mocks.NewMockAuthService(mockCtrl)

	s := server.New(mas)
	assert.IsType(t, &server.Server{}, s)
}

func TestServer_Submit(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mas := mocks.NewMockAuthService(mockCtrl)
	s := server.New(mas)

	// nil request
	_, err := s.Submit(context.Background(), nil)
	assert.Error(t, err)

	// no session
	_, err = s.Submit(context.Background(), &queuepb.QueueSubmitRequest{})
	assert.Error(t, err)

	za := &net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8088,
	}

	ctx := peer.NewContext(context.Background(), &peer.Peer{
		Addr: za,
	})

	// test session
	mas.EXPECT().SetConnection(za.IP.String()).Return(nil).AnyTimes()
	mas.EXPECT().Validate("invalid").Return(fmt.Errorf("invalid")).AnyTimes()
	mas.EXPECT().Validate("valid").Return(nil).AnyTimes()

	_, err = s.Submit(ctx, &queuepb.QueueSubmitRequest{
		Session: "invalid",
		Url:     "some url",
		Depth:   0,
		Job:     "somejob",
	})
	assert.Error(t, err)

	_, err = s.Submit(ctx, &queuepb.QueueSubmitRequest{
		Session: "valid",
		Url:     "some url",
		Depth:   0,
		Job:     "somejob",
	})
	assert.NoError(t, err)

	// validate double request
	_, err = s.Submit(ctx, &queuepb.QueueSubmitRequest{
		Session: "valid",
		Url:     "some url",
		Depth:   0,
		Job:     "somejob",
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "AlreadyExists")
}

func TestServer_Get(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mas := mocks.NewMockAuthService(mockCtrl)
	s := server.New(mas)

	// nil request
	_, err := s.Get(context.Background(), nil)
	assert.Error(t, err)

	// no session
	_, err = s.Get(context.Background(), &queuepb.QueueGetRequest{})
	assert.Error(t, err)

	za := &net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8088,
	}

	mas.EXPECT().SetConnection(za.IP.String()).Return(nil).AnyTimes()
	mas.EXPECT().Validate("invalid").Return(fmt.Errorf("invalid")).AnyTimes()
	mas.EXPECT().Validate("valid").Return(nil).AnyTimes()

	ctx := peer.NewContext(context.Background(), &peer.Peer{
		Addr: za,
	})

	// test session
	_, err = s.Get(ctx, &queuepb.QueueGetRequest{
		Session: "valid",
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "NotFound")

	_, err = s.Get(ctx, &queuepb.QueueGetRequest{
		Session: "invalid",
	})
	assert.Error(t, err)

	_, err = s.Submit(ctx, &queuepb.QueueSubmitRequest{
		Session: "valid",
		Url:     "some url",
		Depth:   1,
		Job:     "somejob",
	})
	assert.NoError(t, err)

	r, err := s.Get(ctx, &queuepb.QueueGetRequest{
		Session: "valid",
	})
	assert.NoError(t, err)
	assert.Equal(t, "some url", r.Url)
	assert.Equal(t, int32(1), r.Depth)
}
