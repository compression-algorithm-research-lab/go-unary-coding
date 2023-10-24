package unary_coding

import (
	bit_buffer "github.com/compression-algorithm-research-lab/go-bit-buffer"
	"github.com/golang-infrastructure/go-gtypes"
)

// 以切片的形式一次编码或者解码多个数字

// EncodeSlice 对多个无符号整数进行一元编码，返回编码后的字节数组
func EncodeSlice[T gtypes.Unsigned](slice []T) []byte {
	buffer := bit_buffer.New()
	for _, value := range slice {

		// 全为1的字节有多少个
		fillByteCount := value / 8
		for ; fillByteCount > 0; fillByteCount-- {
			buffer.WriteByte(0xFF)
		}

		// 生成最后一个字节的值，先填充字节的右边的位数，然后再往左位移
		for c := value % 8; c > 0; c-- {
			buffer.WriteBit(1)
		}
		buffer.WriteBit(0)
	}
	return buffer.Bytes()
}

// DecodeSlice 从一元编码的字节数组中解码出无符号整数数组
func DecodeSlice[T gtypes.Unsigned](bytes []byte) []T {
	buffer := bit_buffer.New().SetBytes(bytes).SeekHead()
	slice := make([]T, 0)
	for !buffer.IsTail() {

		// 先读取一位看看
		bit := buffer.ReadBit()
		if bit == 0 {
			// 说明出现了两个连续的0，那意味着已经读取结束了
			break
		}

		// 然后开始读取连续的1，直到结束
		c := 1
		for {
			bit := buffer.ReadBit()
			if bit == 1 {
				c++
			} else {
				break
			}
		}
		slice = append(slice, T(c))
	}
	return slice
}
