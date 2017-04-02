package goNTCore

import (
	"fmt"
	"io"
	"net"

	"github.com/technomancers/goNTCore/message"
)

//Client is a Network Table client.
type Client struct {
	net.Conn
	connected bool
	name      string
	status    string
}

//NewClient creates a new client to communicate to a Network Table server.
func NewClient(serverHost string, name string) (*Client, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverHost, PORT))
	if err != nil {
		return nil, err
	}
	return &Client{
		Conn:      conn,
		connected: true,
		name:      name,
		status:    "pending",
	}, nil
}

//Close closes the connection to the Network Table server.
func (c *Client) Close() error {
	c.connected = false
	return c.Conn.Close()
}

//Listen for messages sent from the server.
//Best to start this in a go routine.
func (c *Client) Listen() {
	for c.connected {
		possibleMsgType := make([]byte, 1)
		_, err := io.ReadFull(c, possibleMsgType)
		if err != nil {
			if err == io.EOF {
				continue
			}
			c.Close()
			continue
		}
		_, err = message.Unmarshal(possibleMsgType[0], c)
		if err != nil {
			c.Close()
			continue
		}
		//Handle the message that was recieved
	}
}
