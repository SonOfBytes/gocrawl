package client

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

const linkerdPort = 8088

type Client struct {
	connection *grpc.ClientConn
}

func New() *Client {
	return &Client{}
}

func (c *Client) SetConnection(host string) (err error) {
	if c.connection != nil {
		if state := c.connection.GetState(); state == connectivity.Idle || state == connectivity.Ready || state == connectivity.TransientFailure {
			return nil
		}
	}
	c.connection, err = grpc.Dial(fmt.Sprintf("%s:%d", host, linkerdPort), grpc.WithInsecure())
	return err
}
