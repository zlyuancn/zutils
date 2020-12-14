/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/12/11
   Description :
-------------------------------------------------
*/

package zutils

var Ternary = &ternaryUtil{}

type ternaryUtil struct{}

// 如果boole为真返回v1否则返回v2
func (*ternaryUtil) Ternary(boole bool, v1 interface{}, v2 interface{}) interface{} {
	if boole {
		return v1
	}
	return v2
}

// 如果boole为真则执行fn1并返回它的结果, 否则执行fn2并返回它的结果
func (*ternaryUtil) TernaryDo(boole bool, fn1 func() interface{}, fn2 func() interface{}) interface{} {
	if boole {
		return fn1()
	}
	return fn2()
}

// 顺序判断传入的数据, 如果某个数据不是其数据类型的零值则返回它, 否则返回最后一个数据
func (*ternaryUtil) Or(values ...interface{}) interface{} {
	var v interface{}
	for _, v = range values {
		if !Reflect.IsZero(v) {
			return v
		}
	}
	return v
}

// 顺序执行传入的函数并获取其结果, 如果某个结果不是其数据类型的零值则返回它, 后续函数不会被执行, 否则返回最后一个执行结果
func (*ternaryUtil) OrDo(fns ...func() interface{}) interface{} {
	var v interface{}
	for _, fn := range fns {
		v = fn()
		if !Reflect.IsZero(v) {
			return v
		}
	}
	return v
}
