package unary_coding

import (
	"github.com/golang-infrastructure/go-gtypes"
	"strings"
)

// Encode 对无符号数字编码
func Encode[T gtypes.Unsigned](v T) []byte {

	// 全为1的字节有多少个
	fillByteCount := v / 8
	// 生成最后一个字节的值，先填充字节的右边的位数，然后再往左位移
	lastByte := ((2 << (v % 8)) - 1) << (8 - (v % 8))

	// 填充1
	buff := make([]byte, 0)
	for ; fillByteCount > 0; fillByteCount-- {
		buff = append(buff, 0xFF)
	}

	buff = append(buff, byte(lastByte&0xFF))

	return buff
}

// Decode 对之前使用一元编码的字节数组进行解码
func Decode(bytes []byte) uint64 {
	var r uint64
	for _, b := range bytes {
		// 只要最高位为1，就读取统计一次
		for (b & 0x80) == 0x80 {
			r++
			b <<= 1
		}
	}
	return r
}

// ToBinaryString 把字节数组转为二进制字符串的形式
func ToBinaryString(bytes []byte) string {
	buff := strings.Builder{}
	for _, b := range bytes {
		for offset := 7; offset >= 0; offset-- {
			if ((b >> offset) & 0x1) == 1 {
				buff.WriteRune('1')
			} else {
				buff.WriteRune('0')
			}
		}
	}
	return buff.String()
}
