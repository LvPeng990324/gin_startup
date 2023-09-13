package util

import "strconv"

// 字符串转整型
func String2Int(str string) (int, error) {
	return strconv.Atoi(str)
}

// 整型转无符号整型
func Int2Uint(i int) uint {
	return uint(i)
}

// 字符串转无符号整型
func String2Uint(str string) (uint, error) {
	i, err := String2Int(str)
	return Int2Uint(i), err
}
