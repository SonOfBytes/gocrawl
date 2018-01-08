package server_test

import (
	"context"
	"testing"

	"github.com/sonofbytes/gocrawl/authentication/pb"
	"github.com/sonofbytes/gocrawl/authentication/server"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := server.New()
	assert.IsType(t, &server.Server{}, s)
}

func TestServer_Authenticate(t *testing.T) {
	s := server.New()

	// test nil request
	_, err := s.Authenticate(context.Background(), nil)
	assert.Error(t, err)

	_, err = s.Authenticate(context.Background(), &authpb.AuthenticateRequest{
		Username: "x",
		Password: "",
	})
	assert.Error(t, err)

	_, err = s.Authenticate(context.Background(), &authpb.AuthenticateRequest{
		Username: "",
		Password: "y",
	})
	assert.Error(t, err)

	_, err = s.Authenticate(context.Background(), &authpb.AuthenticateRequest{
		Username: "x",
		Password: "y",
	})
	assert.Error(t, err)

	r, err := s.Authenticate(context.Background(), &authpb.AuthenticateRequest{
		Username: "someone",
		Password: "hardcoded",
	})
	assert.NoError(t, err)
	assert.Equal(t, 12, len(r.Session))
}

func TestServer_Validate(t *testing.T) {
	s := server.New()

	_, err := s.Validate(context.Background(), nil)
	assert.Error(t, err)

	_, err = s.Validate(context.Background(), &authpb.ValidateRequest{
		Session: "",
	})
	assert.Error(t, err)

	vr, err := s.Validate(context.Background(), &authpb.ValidateRequest{
		Session: "invalid",
	})
	assert.NoError(t, err)
	assert.False(t, vr.Valid)

	r, err := s.Authenticate(context.Background(), &authpb.AuthenticateRequest{
		Username: "someone",
		Password: "hardcoded",
	})

	vr, err = s.Validate(context.Background(), &authpb.ValidateRequest{
		Session: r.Session,
	})
	assert.NoError(t, err)
	assert.IsType(t, &authpb.ValidateReply{}, vr)
	assert.True(t, vr.Valid)

}
