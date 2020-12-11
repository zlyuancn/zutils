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
