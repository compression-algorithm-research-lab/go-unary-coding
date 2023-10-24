package unary_coding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeSlice(t *testing.T) {
	rawSlice := []uint64{
		1, 2, 3, 4, 3, 2, 1,
	}
	bytes := EncodeSlice(rawSlice)
	binaryString := ToBinaryString(bytes)
	assert.Equal(t, "101101110111101110110100", binaryString)
}

func TestDecodeSlice(t *testing.T) {

	rawSlice := []uint64{
		1, 2, 3, 4, 3, 2, 1,
	}
	bytes := EncodeSlice(rawSlice)

	slice := DecodeSlice[uint64](bytes)
	assert.Equal(t, rawSlice, slice)

}
