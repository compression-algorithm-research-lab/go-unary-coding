package unary_coding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncode(t *testing.T) {

	assert.Equal(t, "00000000", ToBinaryString(Encode(uint64(0))))
	assert.Equal(t, "10000000", ToBinaryString(Encode(uint64(1))))
	assert.Equal(t, "1111111111111110", ToBinaryString(Encode(uint64(15))))
	assert.Equal(t, "111111111111111100000000", ToBinaryString(Encode(uint64(16))))

}

func TestToBinaryString(t *testing.T) {
	binaryString := ToBinaryString([]byte{
		0x1,
		0xF0,
	})
	assert.Equal(t, "0000000111110000", binaryString)
}

func TestDecode(t *testing.T) {
	bytes := Encode(uint64(15))
	value := Decode(bytes)
	assert.Equal(t, uint64(15), value)
}
