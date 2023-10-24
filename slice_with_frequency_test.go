package unary_coding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeSliceWithFrequency(t *testing.T) {
	rawSlice := []uint64{
		10086, 10086, 1024, 10086, 10086, 1024, 10086, 1024, 10086, 10086, 10086, 1024, 10086, 10086,
	}
	bytes, frequency := EncodeSliceWithFrequency(rawSlice)
	assert.Equal(t, 4, len(bytes))
	binaryString := ToBinaryString(bytes)
	assert.Equal(t, "10101101010110101101010101101010", binaryString)
	assert.Equal(t, []uint64{
		uint64(10086),
		uint64(1024),
	}, frequency)

}

func TestDecodeSliceWithFrequency(t *testing.T) {

	rawSlice := []uint64{
		10086, 10086, 1024, 10086, 10086, 1024, 10086, 1024, 10086, 10086, 10086, 1024, 10086, 10086,
	}
	bytes, frequency := EncodeSliceWithFrequency(rawSlice)
	assert.Equal(t, 4, len(bytes))
	binaryString := ToBinaryString(bytes)
	assert.Equal(t, "10101101010110101101010101101010", binaryString)
	assert.Equal(t, []uint64{
		uint64(10086),
		uint64(1024),
	}, frequency)

	unzipSlice := DecodeSliceWithFrequency[uint64](bytes, frequency)
	assert.Equal(t, rawSlice, unzipSlice)

}
