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
	"math/rand"
	"sync"
	"time"
)

var Rand = &randUtil{
	rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	mx:   sync.Mutex{},
}

type randUtil struct {
	rand *rand.Rand
	mx   sync.Mutex
}

// 随机 [0, max-1] 之间的随机数
func (u *randUtil) Rand(max int64) int64 {
	return u.RandStart(0, max)
}

// 随机返回 [start, end-1] 之间的随机数
func (u *randUtil) RandStart(start, end int64) int64 {
	if end <= start {
		return 0
	}

	u.mx.Lock()
	v := u.rand.Int63n(end - start)
	u.mx.Unlock()
	return v + start
}

// 随机 [0, max-1] count 次并把执行索引和随机的数值传给 fn 函数执行
func (u *randUtil) RandF(max, count int64, fn func(i, v int64)) {
	u.RandStartF(0, max, count, fn)
}

// 随机 [start, end-1] count 次并把执行索引和随机的数值传给 fn 函数执行
func (u *randUtil) RandStartF(start, end, count int64, fn func(i, v int64)) {
	if end <= start || count == 0 {
		return
	}

	max := end - start
	for i := int64(0); i < count; i++ {
		u.mx.Lock()
		v := u.rand.Int63n(max)
		u.mx.Unlock()
		fn(i, v+start)
	}
}

// 随机指定长度的文本, 随机字符串候选词来自base
func (u *randUtil) RandText(base string, length int) string {
	if length <= 0 {
		return ""
	}

	tr := []rune(base)
	l := len(tr)

	var buf = &bytes.Buffer{}

	u.mx.Lock()
	for i := 0; i < length; i++ {
		buf.WriteRune(tr[u.rand.Intn(l)])
	}
	u.mx.Unlock()
	return buf.String()
}

// 随机指定长度的文本, 随机字符串候选词配置
func (u *randUtil) RandTextOfConfig(conf *TextConfig, length int) string {
	if length <= 0 {
		return ""
	}

	const text_num = "0123456789"
	const text_lower = "abcdefghijklmnopqrstuvwxyz"
	const text_upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const text_underline = "_"
	const text_space = " "
	const text_special = `~!@#$%^&*()+-=./\|;:<>,?[]{}'"`

	var buff = &bytes.Buffer{}
	buff.Grow(128)
	if conf.Num {
		buff.WriteString(text_num)
	}
	if conf.Lower {
		buff.WriteString(text_lower)
	}
	if conf.Upper {
		buff.WriteString(text_upper)
	}
	if conf.Underline {
		buff.WriteString(text_underline)
	}
	if conf.Space {
		buff.WriteString(text_space)
	}
	if conf.Special {
		buff.WriteString(text_special)
	}
	data := buff.Bytes()

	var out_buff = &bytes.Buffer{}
	if conf.Chinese {
		l := int64(len(data))
		chinese_start := int64(0x4e00)
		chinese_end := int64(0x9fa5)
		u.RandF(l+chinese_end-chinese_start, int64(length), func(i, v int64) {
			if v < l {
				out_buff.WriteByte(data[v])
			} else {
				out_buff.WriteString(string(int(v - l + chinese_start)))
			}
		})
	} else {
		u.RandF(int64(len(data)), int64(length), func(i, v int64) {
			out_buff.WriteByte(data[v])
		})
	}
	return out_buff.String()
}
