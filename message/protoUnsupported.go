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
			mType: mTypeProtoUnsupported,
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
