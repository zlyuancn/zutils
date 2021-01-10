/*
-------------------------------------------------
   Author :       zlyuancn
   date：         2021/1/10
   Description :
-------------------------------------------------
*/

package zutils

var Color = new(colorUtil)

type colorUtil struct{}

type ColorType byte

const (
	ColorDefault = iota + '0' // 默认
	ColorRed                  // 红
	ColorGreen                // 绿
	ColorYellow               // 黄
	ColorBlue                 // 蓝
	ColorMagenta              // 紫
	ColorCyan                 // 深绿
	ColorWrite                // 灰色
)

func (*colorUtil) MakeColorText(color ColorType, text string) string {
	if color == ColorDefault {
		return text
	}
	bs := make([]byte, 9+len(text))
	copy(bs[:5], "\x1b[30m")
	bs[3] = byte(color)
	copy(bs[5:len(bs)-4], text)
	copy(bs[len(bs)-4:], "\x1b[0m")
	return string(bs)
}
