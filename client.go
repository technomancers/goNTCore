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
			c.Log <- NewErrorMessage(err)
			continue
		}
		msg, err := message.Unmarshal(possibleMsgType[0], c)
		if err != nil {
			c.Log <- NewErrorMessage(err)
			continue
		}
		c.handler(msg)
	}
}

//StartHandshake starts the handshake with the server.
func (c *Client) StartHandshake() error {
	return SendMsg(message.NewClientHello(ProtocolVersion, c.name), c)
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
		c.Log <- NewLogMessage("Entry Assign Not Implemented")
	case message.MTypeEntryUpdate:
		c.Log <- NewLogMessage("Entry Update Not Implemented")
	case message.MTypeEntryFlagUpdate:
		c.Log <- NewLogMessage("Entry Flag Update Not Implemented")
	case message.MTypeEntryDelete:
		c.Log <- NewLogMessage("Entry Delete Not Implemented")
	case message.MTypeClearAllEntries:
		c.Log <- NewLogMessage("Clear All Entries Not Implemented")
	default:
		c.Log <- NewErrorMessage(errors.New("Could not handle the message"))
		c.Close()
	}
}

func (c *Client) notifyOfDifference() {
	//find the differences and create new entries for each.
	err := SendMsg(message.NewClientHelloComplete(), c)
	if err != nil {
		c.Log <- NewErrorMessage(err)
	}
	c.status = READY
}
