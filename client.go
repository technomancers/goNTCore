package goNTCore

import (
	"errors"
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
	Log       chan LogMessage
	sendOnly  bool
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
		Log:       make(chan LogMessage),
		sendOnly:  false,
	}, nil
}

//NewSendOnlyClient creates a client that does not keep up with keys or values.
//All this client has is the ability to add and remove variables from the server.
func NewSendOnlyClient(serverHost string, name string) (*Client, error) {
	c, err := NewClient(serverHost, name)
	if err != nil {
		return nil, err
	}
	c.sendOnly = true
	return c, nil
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
			c.Log <- NewErrorMessage(err)
			c.Close()
			continue
		}
		msg, err := message.Unmarshal(possibleMsgType[0], c)
		if err != nil {
			c.Log <- NewErrorMessage(err)
			c.Close()
			continue
		}
		c.handler(msg)
	}
}

//SendMsg to the connected server
func (c *Client) SendMsg(msg message.Messager) error {
	return SendMsg(msg, c)
}

//StartHandshake starts the handshake with the server.
func (c *Client) StartHandshake() error {
	return c.SendMsg(message.NewClientHello(ProtocolVersion, c.name))
}

func (c *Client) handler(msg message.Messager) {
	switch msg.Type() {
	case message.MTypeKeepAlive:
		return
	case message.MTypeServerHello:
		c.status = LISTENING
	case message.MTypeServerHelloComplete:
		c.notifyOfDifference()
	case message.MTypeProtoUnsupported:
		c.Log <- NewErrorMessage(errors.New("Unsupported Protocol Version"))
		c.Close()
	case message.MTypeEntryAssign:
		if c.sendOnly {
			return
		}
		c.Log <- NewLogMessage("Entry Assign Not Implemented")
	case message.MTypeEntryUpdate:
		if c.sendOnly {
			return
		}
		c.Log <- NewLogMessage("Entry Update Not Implemented")
	case message.MTypeEntryFlagUpdate:
		if c.sendOnly {
			return
		}
		c.Log <- NewLogMessage("Entry Flag Update Not Implemented")
	case message.MTypeEntryDelete:
		if c.sendOnly {
			return
		}
		c.Log <- NewLogMessage("Entry Delete Not Implemented")
	case message.MTypeClearAllEntries:
		if c.sendOnly {
			return
		}
		c.Log <- NewLogMessage("Clear All Entries Not Implemented")
	default:
		c.Log <- NewErrorMessage(errors.New("Could not handle the message"))
		c.Close()
	}
}

func (c *Client) notifyOfDifference() {
	//find the differences and create new entries for each.
	err := c.SendMsg(message.NewClientHelloComplete())
	if err != nil {
		c.Log <- NewErrorMessage(err)
	}
	c.status = READY
}
