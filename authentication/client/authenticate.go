package client

import (
	"context"
	"fmt"

	"github.com/sonofbytes/gocrawl/authentication/pb"
)

func (c *Client) Authenticate(username, password string) (session string, err error) {
	if c.connection == nil {
		return session, fmt.Errorf("SetConnection has not been called")
	}

	ac := authpb.NewAuthenticationClient(c.connection)
	r, err := ac.Authenticate(context.Background(), &authpb.AuthenticateRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return session, err
	}

	return r.Session, nil
}
