package utils

import (
	"strconv"
)

// StrToInt string to int
func StrToInt(s string) (n int, err error) {
	n, err = strconv.Atoi(s)
	return
}

// StrToInt64 string to int64
func StrToInt64(s string) (n int64, err error) {
	n, err = strconv.ParseInt(s, 10, 64)
	return
}

// StrToInt32 string to int32
func StrToInt32(s string) (n int32, err error) {
    in, err := strconv.ParseInt(s, 10, 32)
    return int32(in), err
}

// StrToInt32 string to int64
func StrToUInt32(s string) (uint32, error) {
	n, err := strconv.ParseUint(s, 10, 32)
	return uint32(n), err
}

// StrToFloat64 string to float64
func StrToFloat64(s string) (f float64,err error) {
	f, err = strconv.ParseFloat(s, 64)
	return
}

//IsNumeric is_numeric
func IsNumeric(s string) bool {
	if "" == s {
		return false
	}
	if "0" == s {
		return true
	}
	_, err := StrToInt(s)
	if err != nil {
		return false
	}
	return true
}

// Substr 字符串截取
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		end = rl + start + 1
		start = end - length
	} else {
		end = start + length
	}

	return string(rs[start:end])
}
