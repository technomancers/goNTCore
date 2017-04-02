package goNTCore

import (
	"fmt"
	"net"
)

//Client is a Network Table client.
type Client struct {
	conn net.Conn
}

//NewClient creates a new client to communicate to a Network Table server.
func NewClient(serverHost string) (*Client, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverHost, PORT))
	if err != nil {
		return nil, err
	}
	return &Client{
		conn: conn,
	}, nil
}

//Close closes the connection to the Network Table server.
func (c *Client) Close() error {
	return c.conn.Close()
}
