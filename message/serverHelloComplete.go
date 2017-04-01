package message

//ServerHelloComplete is sent from the server after it has finished telling the new client what entries the server has.
type ServerHelloComplete struct {
	message
}

//NewServerHelloComplete creates a new instance of ServerHelloComplete.
func NewServerHelloComplete() *ServerHelloComplete {
	return &ServerHelloComplete{
		message{
			mType: mTypeServerHelloComplete,
		},
	}
}

//MarshalMessage implements Marshaler for Network Table Messages.
func (shc *ServerHelloComplete) MarshalMessage() ([]byte, error) {
	var output []byte
	return append(output, shc.Type()), nil
}
