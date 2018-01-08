package client

import (
	"context"
	"fmt"

	"github.com/sonofbytes/gocrawl/queue/pb"
)

func (c *Client) Submit(ctx context.Context, session string, url string, depth int, job string) (err error) {
	if c.connection == nil {
		err = fmt.Errorf("SetConnection has not been called")
		return
	}

	qc := queuepb.NewQueueClient(c.connection)
	_, err = qc.Submit(context.Background(), &queuepb.QueueSubmitRequest{
		Session: session,
		Url:     url,
		Depth:   int32(depth),
		Job:     job,
	})
	if err != nil {
		return fmt.Errorf("Submit URL failed: %s", err)
	}

	return nil
}
