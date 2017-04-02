package message

import "github.com/technomancers/goNTCore/entry"
import "io"

//ClientHello is sent when a client is first communicating to a server.
type ClientHello struct {
	message
	protocol   [2]byte
	clientName *entry.String
}

//NewClientHello creates a new instance of ClientHello with the specified Protocol and Client Name.
func NewClientHello(protocol [2]byte, clientName string) *ClientHello {
	return &ClientHello{
		message: message{
			mType: mTypeClientHello,
		},
		protocol:   protocol,
		clientName: entry.NewString(clientName),
	}
}

//MarshalMessage implements Marshaler for Network Table Messages.
func (ch *ClientHello) MarshalMessage(writer io.Writer) error {
	_, err := writer.Write([]byte{ch.Type()})
	if err != nil {
		return err
	}
	_, err = writer.Write(ch.protocol[:])
	if err != nil {
		return err
	}
	err = ch.clientName.MarshalEntry(writer)
	return err
}
