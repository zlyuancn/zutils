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

type AutoIncr uint64

// 下一个数, 第一个数是1
func (a *AutoIncr) Next() uint64 {
	return atomic.AddUint64((*uint64)(a), 1)
}

func (a *AutoIncr) Incr(v uint64) uint64 {
	return atomic.AddUint64((*uint64)(a), v)
}

// 创建一个自增计数器
func (u *syncUtil) NewAutoIncr() *AutoIncr {
	return new(AutoIncr)
}

// 获取自增寄存器下一个数
func (u *syncUtil) NextAutoIncrNum() uint64 {
	return u.autoIncr.Next()
}
