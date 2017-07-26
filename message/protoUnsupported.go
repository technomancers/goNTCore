// Copyright (c) 2017, Technomancers. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package message

import "io"

//ProtoUnsupported is sent by the server to the client if the server does not support the specified protocol version.
type ProtoUnsupported struct {
	message
	protocol [2]byte
}

//NewProtoUnsupported creates a new instance of ProtoUnsupported.
func NewProtoUnsupported(protocol [2]byte) *ProtoUnsupported {
	return &ProtoUnsupported{
		message: message{
			mType: MTypeProtoUnsupported,
		},
		protocol: protocol,
	}
}

//MarshalMessage implements Marshaler for Network Table Messages.
func (pu *ProtoUnsupported) MarshalMessage(writer io.Writer) error {
	_, err := writer.Write([]byte{pu.Type()})
	if err != nil {
		return err
	}
	_, err = writer.Write(pu.protocol[:])
	return err
}

//UnmarshalMessage implements Unmarshaler for Network Table Messages and assumes the message type byte has already been read.
func (pu *ProtoUnsupported) UnmarshalMessage(reader io.Reader) error {
	pu.mType = MTypeProtoUnsupported
	protoBuf := make([]byte, 2)

	_, err := io.ReadFull(reader, protoBuf)
	if err != nil {
		return err
	}

	copy(pu.protocol[:], protoBuf)
	return nil
}
