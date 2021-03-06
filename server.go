// Copyright (c) 2017, Technomancers. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goNTCore

import (
	"errors"
	"fmt"
	"io"
	"net"

	"time"

	"github.com/technomancers/goNTCore/message"
	"github.com/technomancers/goNTCore/util"
)

//Server is an instance of a Network Table server.
type Server struct {
	l         net.Listener
	conns     []*Client
	name      string
	periodic  *time.Ticker
	dataTable NetworkTabler
	quit      chan bool
	Log       chan LogMessage
}

//NewServer creates a new Network Table server.
func NewServer(name string, data Data) (*Server, error) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		return nil, err
	}
	return &Server{
		l:         l,
		name:      name,
		Log:       make(chan LogMessage),
		quit:      make(chan bool),
		dataTable: NewTable(data, ""),
	}, nil
}

//Close closes all connections to the server and the listener.
func (s *Server) Close() error {
	for _, c := range s.conns {
		err := c.Close()
		if err != nil {
			return err
		}
	}
	s.quit <- true
	err := s.l.Close()
	return err
}

//Listen starts listening on the network for messages.
//Spin off a new goroutine to connect to the client.
//Keep the connection to the client open to allow for communication in both directions.
func (s *Server) Listen() {
	for {
		conn, err := s.l.Accept()
		if err != nil {
			s.Log <- NewErrorMessage(err)
			continue
		}
		cl := new(Client)
		cl.Conn = conn
		cl.connected = true
		cl.status = PENDING
		go s.handleConn(cl)
	}
}

//SendMsg sends a message to each connected client that is ready.
//Never returns an error and does not wait for execution to finish.
func (s *Server) SendMsg(msg message.Messager) error {
	for _, c := range s.conns {
		if c.connected && c.status == READY {
			go func(cl *Client) {
				if err := cl.SendMsg(msg); err != nil {
					cl.Close()
				}
			}(c)
		}
	}
	return nil
}

//StartPeriodicClean cleans up instances of connections that have been closed.
//It cleans every d (durtaion).
func (s *Server) StartPeriodicClean(d time.Duration) {
	s.periodic = time.NewTicker(d)
	go func() {
		for {
			select {
			case <-s.periodic.C:
				s.cleanClients()
			case <-s.quit:
				s.periodic.Stop()
				return
			}
		}
	}()
}

func (s *Server) clientExist(cl *Client) bool {
	for _, c := range s.conns {
		if c.name == cl.name {
			return true
		}
	}
	return false
}

func (s *Server) addClient(cl *Client) {
	s.conns = append(s.conns, cl)
}

func (s *Server) cleanClients() {
	//filtering without allocating
	temp := s.conns[:0]
	for _, c := range s.conns {
		if c.connected {
			if err := c.SendMsg(message.NewKeepAlive()); err != nil {
				c.Close()
				continue
			}
			temp = append(temp, c)
		}
	}
	s.conns = temp
}

//handleConn takes the connection and starts reading.
func (s *Server) handleConn(cl *Client) {
	for cl.connected {
		possibleMsgType := make([]byte, 1)
		_, err := io.ReadFull(cl, possibleMsgType)
		if err != nil {
			if err != io.EOF {
				s.Log <- NewErrorMessage(err)
			}
			cl.Close()
			continue
		}
		msg, err := message.Unmarshal(possibleMsgType[0], cl)
		if err != nil {
			s.Log <- NewErrorMessage(err)
			cl.Close()
			continue
		}
		s.handler(msg, cl)
	}
}

func (s *Server) handler(msg message.Messager, cl *Client) {
	switch msg.Type() {
	case message.MTypeKeepAlive:
		return
	case message.MTypeClientHello:
		s.startingHandshake(msg.(*message.ClientHello), cl)
	case message.MTypeClientHelloComplete:
		s.finishHandshake(cl)
	case message.MTypeEntryAssign:
		s.Log <- NewLogMessage("Entry Assign Not Implemented")
	case message.MTypeEntryUpdate:
		s.Log <- NewLogMessage("Entry Update Not Implemented")
	case message.MTypeEntryFlagUpdate:
		s.Log <- NewLogMessage("Entry Flag Update Not Implemented")
	case message.MTypeEntryDelete:
		s.Log <- NewLogMessage("Entry Delete Not Implemented")
	case message.MTypeClearAllEntries:
		s.Log <- NewLogMessage("Clear All Entries Not Implemented")
	default:
		s.Log <- NewErrorMessage(errors.New("Could not handle the message"))
		cl.Close()
	}
}

func (s *Server) startingHandshake(msg *message.ClientHello, cl *Client) {
	cl.name = msg.GetClientName()
	msgProto := msg.GetProtocol()
	if !util.Match(msgProto[:], ProtocolVersion[:]) {
		err := cl.SendMsg(message.NewProtoUnsupported(ProtocolVersion))
		if err != nil {
			s.Log <- NewErrorMessage(err)
		}
		cl.Close()
		return
	}
	exist := s.clientExist(cl)
	if !exist {
		s.addClient(cl)
	}
	err := cl.SendMsg(message.NewServerHello(!exist, s.name))
	if err != nil {
		s.Log <- NewErrorMessage(err)
		cl.Close()
	}
	//Send data about Table
	err = cl.SendMsg(message.NewServerHelloComplete())
	if err != nil {
		s.Log <- NewErrorMessage(err)
		cl.Close()
	}
}

func (s *Server) finishHandshake(cl *Client) {
	cl.status = READY
}
