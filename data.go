package goNTCore

//Data is the interface implemented by types that can read and write Network Table at a low level.
type Data interface {
	PutEntry(ent *Entry) error //Creates if new, updates otherwise
	GetEntries(root string) ([]string, error)
	GetEntry(key string) (*Entry, error)
	DeleteEntry(key string) error
	DeleteAll(root string) error
	IsTable(key string) bool
	IsKey(key string) bool
}

type NoopData struct {
}

func (n *NoopData) PutEntry(ent *Entry) error {
	return nil
}
func (n *NoopData) GetEntries(root string) ([]string, error) {
	return nil, nil
}
func (n *NoopData) GetEntry(key string) (*Entry, error) {
	return nil, nil
}
func (n *NoopData) DeleteEntry(key string) error {
	return nil
}
func (n *NoopData) DeleteAll(root string) error {
	return nil
}
func (n *NoopData) IsTable(key string) bool {
	return false
}
func (n *NoopData) IsKey(key string) bool {
	return false
}
