// Copyright (c) 2017, Technomancers. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goNTCore

import (
	"bytes"
	"io"

	"github.com/technomancers/goNTCore/message"
)

const (
	//PORT is the port on which all clients and servers communicate on.
	PORT = 1735
	//PENDING is the client status used to make sure that the handshake has been completed.
	PENDING = "pending"
	//LISTENING is used during the handshake to specify it if looking to see what the server has.
	LISTENING = "listening"
	//READY is used to state that the client has finished the handshake.
	READY = "ready"
)

var (
	//ProtocolVersion is what protocol this package supports
	ProtocolVersion = [2]byte{0x03, 0x00}
)

//SendMsg adds a buffer to the Marshaling before sending so the whole message is sent at once.
func SendMsg(msg message.Messager, writer io.Writer) error {
	sendBuf := new(bytes.Buffer)
	err := msg.MarshalMessage(sendBuf)
	if err != nil {
		return err
	}
	_, err = sendBuf.WriteTo(writer)
	if err != nil {
		return err
	}
	return nil
}
