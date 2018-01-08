package client

import (
	"context"

	"fmt"

	"github.com/sonofbytes/gocrawl/api/pb"
)

func (c *Client) Get(ctx context.Context, session string, url string) (urls []string, err error) {
	if c.connection == nil {
		err = fmt.Errorf("SetConnection has not been called")
		return
	}

	ac := apipb.NewAPIClient(c.connection)
	r, err := ac.Get(ctx, &apipb.APIGetRequest{
		Session: session,
		Url:     url,
	})
	if err != nil {
		return
	}

	return r.Urls, nil
}
