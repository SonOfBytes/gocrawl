package client

import (
	"context"

	"fmt"

	"github.com/sonofbytes/gocrawl/store/pb"
)

func (c *Client) Get(ctx context.Context, session string, url string) (urls []string, err error) {
	if c.connection == nil {
		err = fmt.Errorf("SetConnection has not been called")
		return
	}

	sc := storepb.NewStoreClient(c.connection)
	r, err := sc.Get(ctx, &storepb.StoreGetRequest{
		Session: session,
		Url:     url,
	})
	if err != nil {
		return
	}

	return r.Urls, nil
}
