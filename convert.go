/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/6/30
   Description :
-------------------------------------------------
*/

package zutils

import (
	"fmt"
	"strconv"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
)

var Convert = &convertUtil{
	defaultGBKEncoder: simplifiedchinese.GBK.NewEncoder(),
	defaultGBKDecoder: simplifiedchinese.GBK.NewDecoder(),
}

type convertUtil struct {
	// 默认gbk编码器
	defaultGBKEncoder *encoding.Encoder
	// 默认gbk解码器
	defaultGBKDecoder *encoding.Decoder
}

func (*convertUtil) IntsToInt64(v []int) []int64 {
	out := make([]int64, len(v))
	for i, a := range v {
		out[i] = int64(a)
	}
	return out
}
func (*convertUtil) IntsToString(v []int) []string {
	out := make([]string, len(v))
	for i, a := range v {
		out[i] = strconv.FormatInt(int64(a), 10)
	}
	return out
}
func (*convertUtil) IntsToInterface(v []int) []interface{} {
	out := make([]interface{}, len(v))
	for i, a := range v {
		out[i] = a
	}
	return out
}

func (*convertUtil) Int64sToInt(v []int64) []int {
	out := make([]int, len(v))
	for i, a := range v {
		out[i] = int(a)
	}
	return out
}
func (*convertUtil) Int64sToString(v []int64) []string {
	out := make([]string, len(v))
	for i, a := range v {
		out[i] = strconv.FormatInt(a, 10)
	}
	return out
}
func (*convertUtil) Int64sToInterface(v []int64) []interface{} {
	out := make([]interface{}, len(v))
	for i, a := range v {
		out[i] = a
	}
	return out
}

func (*convertUtil) StringsToInt(v []string) ([]int, error) {
	out := make([]int, len(v))
	for i, a := range v {
		n, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("[%s]不能转为int", a)
		}
		out[i] = int(n)
	}
	return out, nil
}
func (*convertUtil) StringsToInt64(v []string) ([]int64, error) {
	out := make([]int64, len(v))
	for i, a := range v {
		n, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("[%s]不能转为int64", a)
		}
		out[i] = n
	}
	return out, nil
}
func (*convertUtil) StringsToInterface(v []string) []interface{} {
	out := make([]interface{}, len(v))
	for i, a := range v {
		out[i] = a
	}
	return out
}

func (u *convertUtil) UTF8ToGBK(s string) (string, error) {
	return u.defaultGBKEncoder.String(s)
}
func (u *convertUtil) GBKToUTF8(s string) (string, error) {
	return u.defaultGBKDecoder.String(s)
}

func (u *convertUtil) UTF8ToGBKBytes(s []byte) ([]byte, error) {
	return u.defaultGBKEncoder.Bytes(s)
}
func (u *convertUtil) GBKToUTF8Bytes(s []byte) ([]byte, error) {
	return u.defaultGBKDecoder.Bytes(s)
}

// uint64转为bytes, 从右边开始写入
func (*convertUtil) Uint64ToBytes(v uint64) []byte {
	return []byte{
		byte(v >> 56), byte(v >> 48), byte(v >> 40), byte(v >> 32), byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v),
	}
}

// bytes转为uint64, 从右边开始读取
func (*convertUtil) BytesToUint64(b []byte) uint64 {
	return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
}

// uint32转为bytes, 从右边开始写入
func (*convertUtil) Uint32ToBytes(v uint32) []byte {
	return []byte{
		byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v),
	}
}

// bytes转为uint32, 从右边开始读取
func (*convertUtil) BytesToUint32(b []byte) uint32 {
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

// uint16转为bytes, 从右边开始写入
func (*convertUtil) Uint16ToBytes(v uint16) []byte {
	return []byte{
		byte(v >> 8), byte(v),
	}
}

// bytes转为uint16, 从右边开始读取
func (*convertUtil) BytesToUint16(b []byte) uint16 {
	return uint16(b[1]) | uint16(b[0])<<8
}
