package tools

import (
	"bytes"
	"math/rand"
	"time"
)

// 判断数组或切片中是否存在某值，数据类型应一致
func ArrayContainsInteger(items []int, item int) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

// 判断字符串是否包含在byte切片中
func ByteContains(s string, subSlice []byte) bool {
	b := bytes.ToUpper([]byte(s))
	return bytes.Index(subSlice, b) != -1
}

// 忽略大小写地查找byte流中的某一段所在的位置
func BytesIndexIgnoreCase(s, sep []byte) int {
	return bytes.Index(bytes.ToUpper(s), bytes.ToUpper(sep))
}

// 数字转化为byte类型的值
func Number2Bytes(number, length int, reverse bool) []byte {
	var b []byte
	for i := 0; i < length; i++ {
		b = append(b, uint8(number>>(i*8)&0xff))
	}
	if reverse {
		return b
	}
	return ReverseBytes(b)
}

// 反转byte切片
func ReverseBytes(b []byte) []byte {
	for beg, end := 0, len(b)-1; beg < end; beg, end = beg+1, end-1 {
		b[beg], b[end] = b[end], b[beg]
	}
	return b
}

// 随机获取一个单词组合
func RandomWords(length, mode int32) string {
	/*
		mode: 0 lower
		mode: 1 upper
		mode: 2 lower + upper
		mode: 3 any ascii

	*/
	rand.Seed(time.Now().Unix())
	if length <= 0 {
		length = 1 + rand.Int31n(10)
	}

	var array []byte
	if mode == 0 {
		for len(array) < int(length) {
			array = append(array, byte(97+rand.Int31n(26)))
		}
	} else if mode == 1 {
		for len(array) < int(length) {
			array = append(array, byte(65+rand.Int31n(26)))
		}
	} else if mode == 2 {
		for len(array) < int(length) {
			random := 65 + rand.Int31n(58)
			if random >= 91 && random <= 96 {
				continue
			}
			array = append(array, byte(random))
		}
	} else {
		for len(array) < int(length) {
			array = append(array, byte(33+rand.Int31n(95)))
		}
	}
	return string(array)
}
