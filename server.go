package goNTCore

import (
	"fmt"
	"net"
)

//Server is an instance of a Network Table server.
type Server struct {
	l     net.Listener
	conns []net.Conn
}

//NewServer creates a new Network Table server.
func NewServer() (*Server, error) {
	l, err := net.Listen("tct", fmt.Sprintf(":%d", PORT))
	if err != nil {
		return nil, err
	}
	return &Server{
		l: l,
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
	err := s.l.Close()
	return err
}
