/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/12/11
   Description :
-------------------------------------------------
*/

package zutils

import (
	"sync/atomic"
)

type autoIncr struct {
	v uint64
}

// 下一个数, 第一个数是1
func (m *autoIncr) Next() uint64 {
	return atomic.AddUint64(&m.v, 1)
}

// 获取自增寄存器下一个数
func (u *syncUtil) NextAutoIncrNum() uint64 {
	return u.autoIncr.Next()
}
