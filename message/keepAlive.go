package message

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
func (ka *KeepAlive) MarshalMessage() ([]byte, error) {
	var output []byte
	return append(output, ka.Type()), nil
}
