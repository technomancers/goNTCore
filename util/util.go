package util

import (
	"io"
)

const (
	sevenBitMask uint32 = 0x7f
	mostSigBit   uint32 = 0x80
)

//EncodeULeb128 encodes the integer using the LEB128 Standard.
//32 bit unsigned integer should be sufficient as that would represent a data length of 4.2 GB.
func EncodeULeb128(value uint32, writer io.Writer) error {
	//figure out the remaining bits
	remaining := value >> 7
	//If there are still bits available
	for remaining != 0 {
		//Grabs the 7 bits from value then prepends a 1
		_, err := writer.Write([]byte{byte(value&sevenBitMask | mostSigBit)})
		if err != nil {
			return err
		}
		//Adjust the new value
		value = remaining
		//Adjust the new remaining
		remaining >>= 7
	}
	//Write the last bit of the int without prepended 0
	_, err := writer.Write([]byte{byte(value & sevenBitMask)})
	return err
}
