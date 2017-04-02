package message

import "io"

//KeepAlive is a message sent from clients to server to check on the network status.
type KeepAlive struct {
	message
}

//NewKeepAlive creates a new KeepAlive message.
func NewKeepAlive() *KeepAlive {
	return &KeepAlive{
		message{
			mType: mTypeKeepAlive,
		},
	}
}

//MarshalMessage implements Marshaler for Network Table Messages.
func (ka *KeepAlive) MarshalMessage(writer io.Writer) error {
	_, err := writer.Write([]byte{ka.Type()})
	return err
}

//UnmarshalMessage implements Unmarshaler for Network Table Messages and assumes the message type byte has already been read.
func (ka *KeepAlive) UnmarshalMessage(reader io.Reader) error {
	ka.mType = mTypeKeepAlive
	return nil
}
