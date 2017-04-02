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
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte{flags})
	if err != nil {
		return err
	}
	err = sh.serverName.MarshalEntry(writer)
	return err
}

//UnmarshalMessage implements Unmarshaler for Network Table Messages and assumes the message type byte has already been read.
func (sh *ServerHello) UnmarshalMessage(reader io.Reader) error {
	sh.mType = mTypeServerHello
	flagBuf := make([]byte, 1)
	st := new(entry.String)

	_, err := io.ReadFull(reader, flagBuf)
	if err != nil {
		return err
	}
	err = st.UnmarshalEntry(reader)
	if err != nil {
		return err
	}

	sh.firstTimeClient = flagBuf[0]&flagAlreadySeenClientMask != flagAlreadySeenClientMask
	sh.serverName = st
	return nil
}
