package message

import "github.com/technomancers/goNTCore/entry"

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
func (ch *ClientHello) MarshalMessage() ([]byte, error) {
	cnBytes, err := ch.clientName.MarshalEntry()
	if err != nil {
		return nil, err
	}
	var output []byte
	output = append(output, ch.Type())
	output = append(output, ch.protocol[:]...)
	output = append(output, cnBytes...)
	return output, nil
}
