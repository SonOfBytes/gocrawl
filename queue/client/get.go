package client

import (
	"context"

	"fmt"

	"github.com/sonofbytes/gocrawl/queue/pb"
)

func (c *Client) Get(ctx context.Context, session string) (url string, depth int, job string, err error) {
	if c.connection == nil {
		err = fmt.Errorf("SetConnection has not been called")
		return
	}

	qc := queuepb.NewQueueClient(c.connection)
	r, err := qc.Get(ctx, &queuepb.QueueGetRequest{
		Session: session,
	})
	if err != nil {
		return
	}

	return r.Url, int(r.Depth), string(r.Job), nil
}
