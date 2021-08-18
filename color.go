/*
-------------------------------------------------
   Author :       zlyuancn
   date：         2021/1/10
   Description :
-------------------------------------------------
*/

package zutils

var Color = &colorUtil{
	Default: 48,
	Red:     49,
	Green:   50,
	Yellow:  51,
	Blue:    52,
	Magenta: 53,
	Cyan:    54,
	Write:   55,
}

type ColorType byte

type colorUtil struct {
	Default ColorType // 默认
	Red     ColorType // 红
	Green   ColorType // 绿
	Yellow  ColorType // 黄
	Blue    ColorType // 蓝
	Magenta ColorType // 紫
	Cyan    ColorType // 深绿
	Write   ColorType // 灰色
}

func (*colorUtil) MakeColorText(color ColorType, text string) string {
	if color == Color.Default {
		return text
	}
	bs := make([]byte, 9+len(text))
	copy(bs[:5], "\x1b[30m")
	bs[3] = byte(color)
	copy(bs[5:len(bs)-4], text)
	copy(bs[len(bs)-4:], "\x1b[0m")
	return *Convert.BytesToString(bs)
}
