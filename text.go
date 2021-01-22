/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/10/26
   Description :
-------------------------------------------------
*/

package zutils

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

var Text = new(textUtil)

type textUtil struct{}

// 将文本切割并转换为int类型的切片
func (*textUtil) SplitToInt(text string, sep string) ([]int, error) {
	if text == "" {
		return []int{}, nil
	}

	vs := strings.Split(text, sep)
	outs := make([]int, len(vs))
	for i, v := range vs {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		outs[i] = n
	}
	return outs, nil
}

// 将文本切割并转换为int32类型的切片
func (*textUtil) SplitToInt32(text string, sep string) ([]int32, error) {
	if text == "" {
		return []int32{}, nil
	}

	vs := strings.Split(text, sep)
	outs := make([]int32, len(vs))
	for i, v := range vs {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		outs[i] = int32(n)
	}
	return outs, nil
}

// 将文本切割并转换为int64类型的切片
func (*textUtil) SplitToInt64(text string, sep string) ([]int64, error) {
	if text == "" {
		return []int64{}, nil
	}

	vs := strings.Split(text, sep)
	outs := make([]int64, len(vs))
	for i, v := range vs {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		outs[i] = int64(n)
	}
	return outs, nil
}

// 将文本切割并转换为float32类型的切片
func (*textUtil) SplitToFloat32(text string, sep string) ([]float32, error) {
	if text == "" {
		return []float32{}, nil
	}

	vs := strings.Split(text, sep)
	outs := make([]float32, len(vs))
	for i, v := range vs {
		n, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return nil, err
		}
		outs[i] = float32(n)
	}
	return outs, nil
}

// 将文本切割并转换为float64类型的切片
func (*textUtil) SplitToFloat64(text string, sep string) ([]float64, error) {
	if text == "" {
		return []float64{}, nil
	}

	vs := strings.Split(text, sep)
	outs := make([]float64, len(vs))
	for i, v := range vs {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		outs[i] = n
	}
	return outs, nil
}

// 将文本切割并转换为map(文本, 组分隔符, kv分隔符)
func (*textUtil) SplitToMap(text string, groupSep string, kvSep string) (map[string]string, error) {
	m := make(map[string]string)
	if text == "" {
		return m, nil
	}

	values := strings.Split(text, groupSep)
	for _, value := range values {
		if value == "" {
			continue
		}

		vs := strings.Split(value, "=")
		if len(vs) != 2 {
			return nil, fmt.Errorf("参数格式为 key%svalue%skey%svalue", kvSep, kvSep, kvSep)
		}

		if vs[0] == "" {
			return nil, fmt.Errorf("无法解析 [%s]", value)
		}

		m[vs[0]] = vs[1]
	}
	return m, nil
}

// 文本水印, 指定开始位(包含)和结束位(不包含), 下标安全, 放心使用
func (*textUtil) Watermark(text string, start, end int, watermark string) string {
	rt := []rune(text)
	if len(rt) <= start {
		return text
	}

	var bs bytes.Buffer
	for i, s := range rt {
		if i >= start && i < end {
			bs.WriteString(watermark)
		} else {
			bs.WriteRune(s)
		}
	}
	return bs.String()
}

// 忽略大小写检查字符相等
func (*textUtil) EqualCharIgnoreCase(c1, c2 int32) bool {
	if c1 == c2 {
		return true
	}
	switch c1 - c2 {
	case 32: // a - A
		return c1 >= 'a' && c1 <= 'z'
	case -32: // A - a
		return c1 >= 'A' && c1 <= 'Z'
	}
	return false
}

// 忽略大小写检查文本相等
func (u *textUtil) EqualIgnoreCase(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}

	r1 := []rune(s1)
	r2 := []rune(s2)
	if len(r1) != len(r2) {
		return false
	}

	for i, r := range r1 {
		if !u.EqualCharIgnoreCase(r, r2[i]) {
			return false
		}
	}
	return true
}

// 忽略大小写替换所有文本
func (u *textUtil) ReplaceAllIgnoreCase(s, old, new string) string {
	return u.ReplaceIgnoreCase(s, old, new, -1)
}

// 替换n次忽略大小写匹配的文本
func (u *textUtil) ReplaceIgnoreCase(s, old, new string, n int) string {
	if n == 0 || old == new || old == "" {
		return s
	}

	ss := []rune(s)
	sub := []rune(old)
	var buff bytes.Buffer
	var num int
	for offset := 0; offset < len(ss); {
		start := u.searchIgnoreCase(ss, sub, offset)
		if start > -1 {
			buff.WriteString(string(ss[offset:start]))
			buff.WriteString(new)
			offset = start + len(sub)
			num++
		}

		if start == -1 || num == n {
			buff.WriteString(string(ss[offset:]))
			break
		}
	}
	return buff.String()
}

// 忽略大小写查找第一个匹配sub的文本所在位置, 如果不存在返回-1
func (u *textUtil) searchIgnoreCase(ss []rune, sub []rune, start int) int {
	if len(ss)-start < len(sub) {
		return -1
	}

	var has bool
	// 查找开头
	for i := start; i < len(ss); i++ {
		if u.EqualCharIgnoreCase(ss[i], sub[0]) {
			start, has = i, true
			break
		}
	}
	if !has || len(ss)-start < len(sub) {
		return -1
	}
	for i := 1; i < len(sub); i++ {
		if !u.EqualCharIgnoreCase(ss[start+i], sub[i]) {
			return u.searchIgnoreCase(ss, sub, start+1)
		}
	}
	return start
}

// 忽略大小写查找第一个匹配sub的文本所在位置, 如果不存在返回-1
func (u *textUtil) IndexIgnoreCase(s, sub string) int {
	return u.searchIgnoreCase([]rune(s), []rune(sub), 0)
}

// 忽略大小写查找s是否包含sub
func (u *textUtil) ContainsIgnoreCase(s, sub string) bool {
	return u.IndexIgnoreCase(s, sub) >= 0
}

// 忽略大小写测试文本s是否以prefix开头
func (u *textUtil) HasPrefixIgnoreCase(s, prefix string) bool {
	return len(s) >= len(prefix) && u.EqualIgnoreCase(s[0:len(prefix)], prefix)
}

// 忽略大小写测试文本s是否以suffix结束
func (u *textUtil) HasSuffixIgnoreCase(s, suffix string) bool {
	return len(s) >= len(suffix) && u.EqualIgnoreCase(s[len(s)-len(suffix):], suffix)
}
