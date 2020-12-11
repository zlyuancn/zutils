/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/11/22
   Description :
-------------------------------------------------
*/

package zutils

import (
	"errors"
	"fmt"
)

var Recover = new(recoverCli)

type recoverCli struct{}

func (*recoverCli) WarpCall(fn func() error) (err error) {
	// 包装执行, 拦截panic
	defer func() {
		e := recover()
		switch v := e.(type) {
		case nil:
		case error:
			err = v
		case string:
			err = errors.New(v)
		default:
			err = errors.New(fmt.Sprint(err))
		}
	}()

	err = fn()
	return
}
