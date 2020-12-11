/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/12/11
   Description :
-------------------------------------------------
*/

package zutils

import (
	"regexp"
	"strings"
)

var Regex = &regexUtil{
	escape:       regexp.MustCompile(`([\^$.\[\]*\\?+{}|()])`),
	escapeSimple: regexp.MustCompile(`([\^$.\[\]\\+{}|()])`),
}

type regexUtil struct {
	escape       *regexp.Regexp // 关闭所有regex转义字符
	escapeSimple *regexp.Regexp // 不管星号和问号
}

// 关闭文本的所有正则表达式转义字符
func (u *regexUtil) MakeRegexEscapeString(text string) string {
	return u.escape.ReplaceAllString(text, `\$1`)
}

// 将正则表达式视为基础匹配语法, ?表示0或1个值, *表示任何数量的值
func (u *regexUtil) MakeRegexSimpleEscapeString(text string) string {
	text = u.escapeSimple.ReplaceAllString(text, `\$1`)
	text = strings.ReplaceAll(text, "?", ".?")
	text = strings.ReplaceAll(text, "*", ".*?")
	return text
}
