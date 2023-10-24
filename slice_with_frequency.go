package unary_coding

import (
	bit_buffer "github.com/compression-algorithm-research-lab/go-bit-buffer"
	"github.com/golang-infrastructure/go-gtypes"
	"github.com/golang-infrastructure/go-tuple"
	"sort"
)

// EncodeSliceWithFrequency 基于频率对切片进行一元数组编码
func EncodeSliceWithFrequency[T gtypes.Unsigned](slice []T) ([]byte, []T) {

	// 对无符号整数切片中的数字统计词频
	countMap := make(map[T]int, 0)
	for _, v := range slice {
		countMap[v] += 1
	}
	// 根据词频倒序排序
	countSlice := make([]*tuple.Tuple2[T, int], 0)
	for v, c := range countMap {
		countSlice = append(countSlice, tuple.New2(v, c))
	}
	sort.Slice(countSlice, func(i, j int) bool {
		// 逆序
		return countSlice[i].V2 > countSlice[j].V2
	})
	weightSlice := make([]T, 0)
	weightMap := make(map[T]int, 0)
	for i, t := range countSlice {
		weightSlice = append(weightSlice, t.V1)
		weightMap[t.V1] = i + 1
	}

	// 好了，现在，终于可以开始一元编码进行压缩了
	buffer := bit_buffer.New()
	for _, v := range slice {
		weight := weightMap[v]
		for weight > 0 {
			buffer.WriteBit(1)
			weight--
		}
		buffer.WriteBit(0)
	}
	return buffer.Bytes(), weightSlice
}

// DecodeSliceWithFrequency 基于频率对切片进行一元数组编码
func DecodeSliceWithFrequency[T gtypes.Unsigned](bytes []byte, weightSlice []T) []T {

	// 生成权重到无符号整数的映射
	weightMap := make(map[int]T, 0)
	for weight, v := range weightSlice {
		weightMap[weight+1] = v
	}

	// 然后开始解压
	slice := make([]T, 0)
	buffer := bit_buffer.New().SetBytes(bytes)
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
		slice = append(slice, weightMap[c])
	}
	return slice
}
