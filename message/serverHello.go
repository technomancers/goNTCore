package message

import (
	"github.com/technomancers/goNTCore/entry"
)

const (
	flagAlreadySeenMask byte = 0x01
)

//ServerHello is a message sent from the server immediatley after it recieves the ClientHello message.
type ServerHello struct {
	message
	firstTimeClient bool
	serverName      *entry.String
}

//NewServerHello creates a new instance of ServerHello.
func NewServerHello(firstTime bool, serverName string) *ServerHello {

	return &ServerHello{
		message: message{
			mType: mTypeServerHello,
		},
		firstTimeClient: firstTime,
		serverName:      entry.NewString(serverName),
	}
}

//MarshalMessage implements Marshaler for Network Table Messages.
func (sh *ServerHello) MarshalMessage() ([]byte, error) {
	snBytes, err := sh.serverName.MarshalEntry()
	if err != nil {
		return nil, err
	}
	var output []byte
	output = append(output, sh.Type())
	if sh.firstTimeClient {
		output = append(output, 0x00)
	} else {
		output = append(output, flagAlreadySeenMask)
	}
	output = append(output, snBytes...)
	return output, nil
}
