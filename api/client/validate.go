package client

import (
	"context"
	"fmt"

	"time"

	"github.com/sonofbytes/gocrawl/api/pb"
)

const sessionCacheExpire = time.Second * 5

func (c *Client) Validate(session string) error {
	if expire, ok := c.sessionCache.Load(session); ok {
		if expire.(time.Time).After(time.Now()) {
			return nil
		}
	}

	if c.connection == nil {
		return fmt.Errorf("SetConnection has not been called")
	}

	ac := apipb.NewAPIClient(c.connection)
	r, err := ac.Validate(context.Background(), &apipb.APIValidateRequest{Session: session})
	if err != nil {
		return fmt.Errorf("Session validation failed: %s", err)
	}

	if !r.Valid {
		return fmt.Errorf("Invalid session token: %s", session)
	}

	c.sessionCache.Store(session, time.Now().Add(sessionCacheExpire))

	return nil
}
