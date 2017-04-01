package message

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
func (pu *ProtoUnsupported) MarshalMessage() ([]byte, error) {
	var output []byte
	output = append(output, pu.Type())
	output = append(output, pu.protocol[:]...)
	return output, nil
}
