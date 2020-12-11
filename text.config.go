/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/12/11
   Description :
-------------------------------------------------
*/

package zutils

import (
	"bytes"
	"regexp"
)

type TextConfig struct {
	Num       bool // 数字
	Lower     bool // 小写字母
	Upper     bool // 大写字母
	Underline bool // 下划线
	Space     bool // 空格
	Special   bool // 特殊字符 ~!@#$%^&*()+-=_./\|;:<>,?[]{}'"
	Chinese   bool // 中文
}

// 检查文本符合文本配置
func (u *textUtil) CheckMatchText(conf *TextConfig, text string) bool {
	return u.GetNotMatchText(conf, text) == ""
}

// 获取文本不匹配文本配置的部分字符串
func (*textUtil) GetNotMatchText(conf *TextConfig, text string) string {
	if text == "" {
		return ""
	}

	text_num := "0-9"
	text_lower := "a-z"
	text_upper := "A-Z"
	text_underline := "_"
	text_space := " "
	text_special := `~!@#$%^&*()+\-=./\\|;:<>,?\[\]{}'"`
	text_chinese := "\u4e00-\u9fa5"

	var data_buff = &bytes.Buffer{}
	data_buff.Grow(128)
	data_buff.WriteString("[")
	if conf.Num {
		data_buff.WriteString(text_num)
	}
	if conf.Lower {
		data_buff.WriteString(text_lower)
	}
	if conf.Upper {
		data_buff.WriteString(text_upper)
	}
	if conf.Underline {
		data_buff.WriteString(text_underline)
	}
	if conf.Space {
		data_buff.WriteString(text_space)
	}
	if conf.Special {
		data_buff.WriteString(text_special)
	}
	if conf.Chinese {
		data_buff.WriteString(text_chinese)
	}
	data_buff.WriteString("]+")
	allow_text_compile := regexp.MustCompile(data_buff.String())

	return allow_text_compile.ReplaceAllLiteralString(text, "")
}
