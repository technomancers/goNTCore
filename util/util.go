package util

import (
	"bytes"
)

const (
	sevenBitMask uint32 = 0x7f
	mostSigBit   uint32 = 0x80
)

//EncodeULeb128 encodes the integer using the LEB128 Standard.
//32 bit unsigned integer should be sufficient as that would represent a data length of 4.2 GB.
func EncodeULeb128(value uint32) []byte {
	//figure out the remaining bits
	remaining := value >> 7
	var buf = new(bytes.Buffer)
	//If there are still bits available
	for remaining != 0 {
		//Grabs the 7 bits from value then prepends a 1
		buf.WriteByte(byte(value&sevenBitMask | mostSigBit))
		//Adjust the new value
		value = remaining
		//Adjust the new remaining
		remaining >>= 7
	}
	//Write the last bit of the int without prepended 0
	buf.WriteByte(byte(value & sevenBitMask))
	return buf.Bytes()
}
