package message

import (
	"io"

	"github.com/technomancers/goNTCore/entry"
)

const (
	flagAlreadySeenClientMask byte = 0x01
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
func (sh *ServerHello) MarshalMessage(writer io.Writer) error {
	flags := byte(0x00)
	if !sh.firstTimeClient {
		flags = flags | flagAlreadySeenClientMask
	}
	_, err := writer.Write([]byte{sh.Type()})
	_, err = writer.Write([]byte{flags})
	err = sh.serverName.MarshalEntry(writer)
	return err
}
