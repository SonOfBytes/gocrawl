package client

import (
	"context"
	"fmt"

	"github.com/sonofbytes/gocrawl/store/pb"
)

func (c *Client) Submit(ctx context.Context, session string, url string, urls []string) (err error) {
	if c.connection == nil {
		err = fmt.Errorf("SetConnection has not been called")
		return
	}

	sc := storepb.NewStoreClient(c.connection)
	_, err = sc.Submit(context.Background(), &storepb.StoreSubmitRequest{
		Session: session,
		Url:     url,
		Urls:    urls,
	})
	if err != nil {
		return fmt.Errorf("Submit URL failed: %s", err)
	}

	return nil
}
