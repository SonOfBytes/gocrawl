package client

import (
	"context"

	"fmt"

	"github.com/sonofbytes/gocrawl/api/pb"
)

func (c *Client) Submit(ctx context.Context, session string, url string) (job string, err error) {
	if c.connection == nil {
		err = fmt.Errorf("SetConnection has not been called")
		return
	}

	sc := apipb.NewAPIClient(c.connection)
	r, err := sc.Submit(context.Background(), &apipb.APISubmitRequest{
		Session: session,
		Url:     url,
	})
	if err != nil {
		return "", err
	}

	return r.Job, nil
}
