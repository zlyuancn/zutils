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
func (timerUtils) NewDoTicker(d time.Duration, fn func(count int, t time.Time)) context.CancelFunc {
	timer := time.NewTicker(d)
	stopCh := make(chan struct{})
	go func() {
		count := 0
		start := true
		for start {
			select {
			case t := <-timer.C:
				count++
				fn(count, t)
			case <-stopCh:
				start = false
			}
		}

		timer.Stop()
		stopCh <- struct{}{}
	}()
	return func() {
		stopCh <- struct{}{}
		<-stopCh
	}
}
