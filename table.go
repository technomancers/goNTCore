package goNTCore

import (
	"github.com/technomancers/goNTCore/util"
)

//NetworkTabler is the interface implemented by types that can read and write Network Table at a high level.
type NetworkTabler interface {
	ContainsKey(key string) bool
	ContainsTable(key string) bool
	Delete(key string)
	DeleteAll()
	IsPersisted(key string) bool
	GetKeys() []string
	GetTable(key string) NetworkTabler
	GetBoolean(key string, def bool) bool
	PutBoolean(key string, val bool) bool
	GetNumber(key string, def float64) float64
	PutNumber(key string, val float64) bool
	GetString(key string, def string) string
	PutString(key string, val string) bool
	GetRaw(key string, def []byte) []byte
	PutRaw(key string, val []byte) bool
	GetBooleanArray(key string, def []bool) []bool
	PutBooleanArray(key string, val []bool) bool
	GetNumberArray(key string, def []float64) []float64
	PutNumberArray(key string, val []float64) bool
	GetStringArray(key string, def []string) []string
	PutStringArray(key string, val []string) bool
}

//Table implements NetworkTabler based on a backend that needs to exist on create.
type Table struct {
	data Data
	root string
}

//NewTable creates a new Network Table that is based off of the data and root passed in.
//Pass in an empty string or "/" for root table.
func NewTable(d Data, root string) *Table {
	if root == "" {
		root = "/"
	}
	root = util.SanatizeKey(root)
	return &Table{
		data: d,
		root: root,
	}
}

//ContainsKey returns true if the key exist in the table.
func (t *Table) ContainsKey(key string) bool {
	key = t.getKey(key)
	return t.data.IsKey(key)
}

//ContainsTable return true if the table exist in the table.
func (t *Table) ContainsTable(key string) bool {
	key = t.getKey(key)
	return t.data.IsTable(key)
}

//Delete deletes the given key from the table.
func (t *Table) Delete(key string) {
	key = t.getKey(key)
	t.data.DeleteEntry(key)
}

//DeleteAll deletes all keys from the table.
func (t *Table) DeleteAll() {
	t.data.DeleteAll()
}

//IsPersisted returns true if the key is to be persisted.
//Returns false if the key does not exist.
func (t *Table) IsPersisted(key string) bool {
	key = t.getKey(key)
	entry, err := t.data.GetEntry(key)
	if err != nil {
		return false
	}
	return entry.Persitant
}

//GetKeys returns all the keys in the table.
func (t *Table) GetKeys() []string {
	return nil
}

//GetTable gets a table with the specified key.
//If table does not exist it creates a new one.
func (t *Table) GetTable(key string) NetworkTabler {
	key = t.getKey(key)
	return NewTable(t.data, key)
}

//GetBoolean gets the value of key as a boolean.
//If the value is not of type boolean it returns the default value passed in.
func (t *Table) GetBoolean(key string, def bool) bool {
	key = t.getKey(key)
	entry, err := t.data.GetEntry(key)
	if err != nil {
		return def
	}
	b, ok := entry.Value.(bool)
	if !ok {
		return def
	}
	return b
}

//PutBoolean puts the boolean value in the table.
//If value exist it updates it.
//If value doesn't exist it will add it.
//Returns false if key exist of a different type.
func (t *Table) PutBoolean(key string, val bool) bool {
	return false
}

//GetNumber gets the value of key as a float64.
//If the value is not of type float64 it returns the default value passed in.
func (t *Table) GetNumber(key string, def float64) float64 {
	key = t.getKey(key)
	entry, err := t.data.GetEntry(key)
	if err != nil {
		return def
	}
	d, ok := entry.Value.(float64)
	if !ok {
		return def
	}
	return d
}

//PutNumber puts the float64 value in the table.
//If value exist it updates it.
//If value doesn't exist it will add it.
//Returns false if key exist of a different type.
func (t *Table) PutNumber(key string, val float64) bool {
	return false
}

//GetString gets the value of key as a string.
//If the value is not of type string it returns the default value passed in.
func (t *Table) GetString(key string, def string) string {
	key = t.getKey(key)
	entry, err := t.data.GetEntry(key)
	if err != nil {
		return def
	}
	s, ok := entry.Value.(string)
	if !ok {
		return def
	}
	return s
}

//PutString puts the string value in the table.
//If value exist it updates it.
//If value doesn't exist it will add it.
//Returns false if key exist of a different type.
func (t *Table) PutString(key string, val string) bool {
	return false
}

//GetRaw gets the value of key as a slice of bytes.
//If the value is not of type byte slice it returns the default value passed in.
func (t *Table) GetRaw(key string, def []byte) []byte {
	key = t.getKey(key)
	entry, err := t.data.GetEntry(key)
	if err != nil {
		return def
	}
	r, ok := entry.Value.([]byte)
	if !ok {
		return def
	}
	return r
}

//PutRaw puts the taw value in the table.
//If value exist it updates it.
//If value doesn't exist it will add it.
//Returns false if key exist of a different type.
func (t *Table) PutRaw(key string, val []byte) bool {
	return false
}

//GetBooleanArray gets the value of key as a slice of booleans.
//If the value is not of type boolean slice it returns the default value passed in.
func (t *Table) GetBooleanArray(key string, def []bool) []bool {
	key = t.getKey(key)
	entry, err := t.data.GetEntry(key)
	if err != nil {
		return def
	}
	ba, ok := entry.Value.([]bool)
	if !ok {
		return def
	}
	return ba
}

//PutBooleanArray puts the slice of boolean in the table.
//If value exist it updates it.
//If value doesn't exist it will add it.
//Returns false if key exist of a different type.
func (t *Table) PutBooleanArray(key string, val []bool) bool {
	return false
}

//GetNumberArray gets the value of key as a slice of float64s.
//If the value is not of type float64 slice it returns the default value passed in.
func (t *Table) GetNumberArray(key string, def []float64) []float64 {
	key = t.getKey(key)
	entry, err := t.data.GetEntry(key)
	if err != nil {
		return def
	}
	da, ok := entry.Value.([]float64)
	if !ok {
		return def
	}
	return da
}

//PutNumberArray puts the slice of float64 in the table.
//If value exist it updates it.
//If value doesn't exist it will add it.
//Returns false if key exist of a different type.
func (t *Table) PutNumberArray(key string, val []float64) bool {
	return false
}

//GetStringArray gets the value of key as a slice of strings.
//If the value is not of type string slice it returns the default value passed in.
func (t *Table) GetStringArray(key string, def []string) []string {
	key = t.getKey(key)
	entry, err := t.data.GetEntry(key)
	if err != nil {
		return def
	}
	sa, ok := entry.Value.([]string)
	if !ok {
		return def
	}
	return sa
}

//PutStringArray puts the slice of string in the table.
//If value exist it updates it.
//If value doesn't exist it will add it.
//Returns false if key exist of a different type.
func (t *Table) PutStringArray(key string, val []string) bool {
	return false
}

func (t *Table) getKey(key string) string {
	return util.KeyJoin(t.root, key)
}
