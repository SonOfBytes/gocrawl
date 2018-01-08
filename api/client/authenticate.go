package client

import (
	"context"
	"fmt"

	"github.com/sonofbytes/gocrawl/api/pb"
)

func (c *Client) Authenticate(username, password string) (session string, err error) {
	if c.connection == nil {
		return session, fmt.Errorf("SetConnection has not been called")
	}

	ac := apipb.NewAPIClient(c.connection)
	r, err := ac.Authenticate(context.Background(), &apipb.APIAuthenticateRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return session, err
	}

	return r.Session, nil
}
