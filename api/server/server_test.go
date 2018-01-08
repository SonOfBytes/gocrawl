package server_test

import (
	"context"
	"errors"
	"testing"

	"fmt"
	"net"

	"math/rand"

	"github.com/golang/mock/gomock"
	"github.com/sonofbytes/gocrawl/api/mocks"
	"github.com/sonofbytes/gocrawl/api/pb"
	"github.com/sonofbytes/gocrawl/api/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func TestNew(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mas := mocks.NewMockAuthService(mockCtrl)
	mqs := mocks.NewMockQueueService(mockCtrl)
	mss := mocks.NewMockStoreService(mockCtrl)

	s := server.New(mas, mqs, mss)
	assert.IsType(t, &server.Server{}, s)
}

func TestServer_Submit(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mas := mocks.NewMockAuthService(mockCtrl)
	mqs := mocks.NewMockQueueService(mockCtrl)
	mss := mocks.NewMockStoreService(mockCtrl)

	s := server.New(mas, mqs, mss)

	za := &net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8088,
	}

	ctx := peer.NewContext(context.Background(), &peer.Peer{
		Addr: za,
	})

	mas.EXPECT().SetConnection(za.IP.String()).Return(nil).AnyTimes()
	mas.EXPECT().Validate("invalid").Return(fmt.Errorf("invalid")).AnyTimes()
	mas.EXPECT().Validate("valid").Return(nil).AnyTimes()

	// test nil request
	_, err := s.Submit(ctx, nil)
	assert.Error(t, err)

	_, err = s.Submit(ctx, &apipb.APISubmitRequest{
		Session: "valid",
		Url:     "",
	})
	assert.Error(t, err)

	_, err = s.Submit(ctx, &apipb.APISubmitRequest{
		Session: "",
		Url:     "someurl",
	})
	assert.Error(t, err)

	// set random seed to something known
	rand.Seed(1234)

	mqs.EXPECT().SetConnection(za.IP.String()).Return(nil).AnyTimes()
	mqs.EXPECT().
		Submit(ctx, "valid", "invalidurl", 0, "QFfCmSqpMUJf").
		Return(status.Error(codes.InvalidArgument, "invalid url"))
	mqs.EXPECT().
		Submit(ctx, "valid", "someurl", 0, "COdmXhBvKfaW").
		Return(nil)

	r, err := s.Submit(ctx, &apipb.APISubmitRequest{
		Session: "invalid",
		Url:     "someurl",
	})
	assert.Error(t, err)
	assert.Equal(t, &apipb.APISubmitReply{}, r)

	r, err = s.Submit(ctx, &apipb.APISubmitRequest{
		Session: "valid",
		Url:     "invalidurl",
	})
	assert.Error(t, err)
	e, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, e.Code())
	assert.Equal(t, &apipb.APISubmitReply{Job: ""}, r)

	r, err = s.Submit(ctx, &apipb.APISubmitRequest{
		Session: "valid",
		Url:     "someurl",
	})
	assert.NoError(t, err)
}

func TestServer_Get(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mas := mocks.NewMockAuthService(mockCtrl)
	mqs := mocks.NewMockQueueService(mockCtrl)
	mss := mocks.NewMockStoreService(mockCtrl)

	s := server.New(mas, mqs, mss)

	za := &net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8088,
	}

	ctx := peer.NewContext(context.Background(), &peer.Peer{
		Addr: za,
	})

	mas.EXPECT().SetConnection(za.IP.String()).Return(nil).AnyTimes()
	mss.EXPECT().SetConnection(za.IP.String()).Return(nil).AnyTimes()
	mas.EXPECT().Validate("invalid").Return(fmt.Errorf("invalid")).AnyTimes()
	mas.EXPECT().Validate("valid").Return(nil).AnyTimes()

	// test nil request
	_, err := s.Get(ctx, nil)
	assert.Error(t, err)
	e, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, e.Code())

	_, err = s.Get(ctx, &apipb.APIGetRequest{
		Session: "valid",
		Url:     "",
	})
	assert.Error(t, err)
	e, ok = status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, e.Code())

	_, err = s.Get(ctx, &apipb.APIGetRequest{
		Session: "",
		Url:     "someurl",
	})
	assert.Error(t, err)
	e, ok = status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, e.Code())

	r, err := s.Get(ctx, &apipb.APIGetRequest{
		Session: "invalid",
		Url:     "someurl",
	})
	assert.Error(t, err)
	assert.Equal(t, &apipb.APIGetReply{}, r)

	mss.EXPECT().Get(ctx, "valid", "someurl")
	r, err = s.Get(ctx, &apipb.APIGetRequest{
		Session: "valid",
		Url:     "someurl",
	})
	assert.NoError(t, err)
}

func TestServer_Authenticate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mas := mocks.NewMockAuthService(mockCtrl)
	mqs := mocks.NewMockQueueService(mockCtrl)
	mss := mocks.NewMockStoreService(mockCtrl)

	s := server.New(mas, mqs, mss)

	za := &net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8088,
	}

	ctx := peer.NewContext(context.Background(), &peer.Peer{
		Addr: za,
	})

	mas.EXPECT().SetConnection(za.IP.String()).Return(nil).AnyTimes()

	// test nil
	_, err := s.Authenticate(ctx, nil)
	assert.Error(t, err)
	e, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, e.Code())

	_, err = s.Authenticate(ctx, &apipb.APIAuthenticateRequest{
		Username: "someone",
		Password: "",
	})
	assert.Error(t, err)
	e, ok = status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, e.Code())

	_, err = s.Authenticate(ctx, &apipb.APIAuthenticateRequest{
		Username: "",
		Password: "hardcoded",
	})
	assert.Error(t, err)
	e, ok = status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, e.Code())

	mas.EXPECT().Authenticate("someone", "badpassword").
		Return("", errors.New("invalid"))
	r, err := s.Authenticate(ctx, &apipb.APIAuthenticateRequest{
		Username: "someone",
		Password: "badpassword",
	})
	assert.Error(t, err)
	assert.Equal(t, &apipb.APIAuthenticateReply{}, r)

	mas.EXPECT().Authenticate("someone", "goodpassword").
		Return("asession", nil)
	r, err = s.Authenticate(ctx, &apipb.APIAuthenticateRequest{
		Username: "someone",
		Password: "goodpassword",
	})
	assert.NoError(t, err)
	assert.Equal(t, &apipb.APIAuthenticateReply{Session: "asession"}, r)
}

func TestServer_Validate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mas := mocks.NewMockAuthService(mockCtrl)
	mqs := mocks.NewMockQueueService(mockCtrl)
	mss := mocks.NewMockStoreService(mockCtrl)

	s := server.New(mas, mqs, mss)

	za := &net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8088,
	}

	ctx := peer.NewContext(context.Background(), &peer.Peer{
		Addr: za,
	})

	mas.EXPECT().SetConnection(za.IP.String()).Return(nil).AnyTimes()

	// test nil
	_, err := s.Validate(ctx, nil)
	assert.Error(t, err)
	e, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, e.Code())

	_, err = s.Validate(ctx, &apipb.APIValidateRequest{
		Session: "",
	})
	assert.Error(t, err)
	e, ok = status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, e.Code())

	mas.EXPECT().Validate("invalid").
		Return(errors.New("invalid"))
	r, err := s.Validate(ctx, &apipb.APIValidateRequest{
		Session: "invalid",
	})
	assert.Error(t, err)
	assert.Equal(t, &apipb.APIValidateReply{Valid: false}, r)

	mas.EXPECT().Validate("asessiom").
		Return(nil)
	r, err = s.Validate(ctx, &apipb.APIValidateRequest{
		Session: "asessiom",
	})
	assert.NoError(t, err)
	assert.Equal(t, &apipb.APIValidateReply{Valid: true}, r)
}
