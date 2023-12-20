/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/12/11
   Description :
-------------------------------------------------
*/

package zutils

import (
	"context"
	"time"
)

var Timer = new(timerUtils)

type timerUtils struct{}

// 创建一个Ticker, 每隔 d 会执行一次 fn
func (timerUtils) NewDoTicker(d time.Duration, fn func(i int, t time.Time)) context.CancelFunc {
	timer := time.NewTicker(d)
	go func() {
		i := 0
		for t := range timer.C {
			fn(i, t)
			i++
		}
	}()
	return timer.Stop
}
