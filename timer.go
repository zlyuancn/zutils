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

// 创建一个Ticker
func NewTicker(d time.Duration) (<-chan time.Time, context.CancelFunc) {
	done := make(chan struct{})
	cc := make(chan time.Time, 1)
	go func() {
		timer := time.NewTicker(d)
		defer func() {
			timer.Stop()
			close(cc)
		}()
		for {
			select {
			case t := <-timer.C:
				select {
				case cc <- t:
				default:
				}
			case <-done:
				return
			}
		}
	}()
	return cc, func() {
		close(done)
	}
}

// 创建一个Ticker, 每隔 d 会执行一次 fn
func NewDoTicker(d time.Duration, fn func(i int, t time.Time)) context.CancelFunc {
	done := make(chan struct{})
	go func() {
		timer := time.NewTicker(d)
		defer timer.Stop()
		var i int
		for {
			select {
			case t := <-timer.C:
				fn(i, t)
				i++
			case <-done:
				return
			}
		}
	}()
	return func() {
		close(done)
	}
}
