package main

import (
	"fmt"
)

func main() {
	var num int64 = 65535
	decodeResult := zigzagDecode(int64(num))
	encodeResult := zigzagEncode(decodeResult)
	fmt.Println(decodeResult, "-->", encodeResult)
}

func zigzagEecode(_num int64) []byte {
	var result []byte
	for {
		if _num>>7&0x7f == 0 {
			result = append(result, byte(_num))
			return result
		}

		v := _num&0x7f | 0x80
		result = append(result, byte(v))
		_num = _num >> 7
	}
}

func zigzagDncode(_decode []byte) int64 {
	var result int64
	for i, v := range _decode {
		var tmp int64 = int64(v)
		if tmp > 0x80 {
			tmp -= 0x80
		}
		tmp = tmp << (7 * uint64(i))
		result += tmp
	}
	return result
}
