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
)

var Sync = &syncUtil{
	autoIncr: new(AutoIncr),
}

type syncUtil struct {
	autoIncr *AutoIncr
}

// 执行fn, 如果ctx结束返回err, 注意: ctx结束不会打断已经开始执行的fn
func (*syncUtil) DoWithContext(ctx context.Context, fn func() interface{}) (out interface{}) {
	if ctx == nil || ctx == context.Background() || ctx == context.TODO() {
		return fn()
	}

	done := make(chan struct{}, 1)
	go func() {
		out = fn()
		done <- struct{}{}
	}()

	select {
	case <-done:
		return out
	case <-ctx.Done():
		return ctx.Err()
	}
}
