/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/12/11
   Description :
-------------------------------------------------
*/

package zutils

// Deprecated: 建议使用 lo.Ternary   eg: https://github.com/samber/lo
var Ternary = &ternaryUtil{}

type ternaryUtil struct{}

// 如果boole为真返回v1否则返回v2
func (*ternaryUtil) Ternary(boole bool, v1 interface{}, v2 interface{}) interface{} {
	if boole {
		return v1
	}
	return v2
}
func (*ternaryUtil) TernaryString(boole bool, v1, v2 string) string {
	if boole {
		return v1
	}
	return v2
}
func (*ternaryUtil) TernaryBytes(boole bool, v1, v2 []byte) []byte {
	if boole {
		return v1
	}
	return v2
}
func (*ternaryUtil) TernaryBool(boole bool, v1, v2 bool) bool {
	if boole {
		return v1
	}
	return v2
}

func (*ternaryUtil) TernaryInt(boole bool, v1, v2 int) int {
	if boole {
		return v1
	}
	return v2
}
func (*ternaryUtil) TernaryInt8(boole bool, v1, v2 int8) int8 {
	if boole {
		return v1
	}
	return v2
}
func (*ternaryUtil) TernaryInt16(boole bool, v1, v2 int16) int16 {
	if boole {
		return v1
	}
	return v2
}
func (*ternaryUtil) TernaryInt32(boole bool, v1, v2 int32) int32 {
	if boole {
		return v1
	}
	return v2
}
func (*ternaryUtil) TernaryInt64(boole bool, v1, v2 int64) int64 {
	if boole {
		return v1
	}
	return v2
}

func (*ternaryUtil) TernaryUint(boole bool, v1, v2 uint) uint {
	if boole {
		return v1
	}
	return v2
}
func (*ternaryUtil) TernaryUint8(boole bool, v1, v2 uint8) uint8 {
	if boole {
		return v1
	}
	return v2
}
func (*ternaryUtil) TernaryUint16(boole bool, v1, v2 uint16) uint16 {
	if boole {
		return v1
	}
	return v2
}
func (*ternaryUtil) TernaryUint32(boole bool, v1, v2 uint32) uint32 {
	if boole {
		return v1
	}
	return v2
}
func (*ternaryUtil) TernaryUint64(boole bool, v1, v2 uint64) uint64 {
	if boole {
		return v1
	}
	return v2
}

func (*ternaryUtil) TernaryFloat32(boole bool, v1, v2 float32) float32 {
	if boole {
		return v1
	}
	return v2
}
func (*ternaryUtil) TernaryFloat64(boole bool, v1, v2 float64) float64 {
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
func (*ternaryUtil) OrString(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

// 顺序判断传入的数据, 如果某个数据不是nil则返回它, 否则返回最后一个数据
func (*ternaryUtil) OrBytes(values ...[]byte) []byte {
	for _, v := range values {
		if v != nil {
			return v
		}
	}
	return nil
}
func (*ternaryUtil) OrBool(values ...bool) bool {
	for _, v := range values {
		if v {
			return true
		}
	}
	return false
}

func (*ternaryUtil) OrInt(values ...int) int {
	for _, v := range values {
		if v != 0 {
			return v
		}
	}
	return 0
}
func (*ternaryUtil) OrInt8(values ...int8) int8 {
	for _, v := range values {
		if v != 0 {
			return v
		}
	}
	return 0
}
func (*ternaryUtil) OrInt16(values ...int16) int16 {
	for _, v := range values {
		if v != 0 {
			return v
		}
	}
	return 0
}
func (*ternaryUtil) OrInt32(values ...int32) int32 {
	for _, v := range values {
		if v != 0 {
			return v
		}
	}
	return 0
}
func (*ternaryUtil) OrInt64(values ...int64) int64 {
	for _, v := range values {
		if v != 0 {
			return v
		}
	}
	return 0
}

func (*ternaryUtil) OrUint(values ...uint) uint {
	for _, v := range values {
		if v != 0 {
			return v
		}
	}
	return 0
}
func (*ternaryUtil) OrUint8(values ...uint8) uint8 {
	for _, v := range values {
		if v != 0 {
			return v
		}
	}
	return 0
}
func (*ternaryUtil) OrUint16(values ...uint16) uint16 {
	for _, v := range values {
		if v != 0 {
			return v
		}
	}
	return 0
}
func (*ternaryUtil) OrUint32(values ...uint32) uint32 {
	for _, v := range values {
		if v != 0 {
			return v
		}
	}
	return 0
}
func (*ternaryUtil) OrUint64(values ...uint64) uint64 {
	for _, v := range values {
		if v != 0 {
			return v
		}
	}
	return 0
}

func (*ternaryUtil) OrFloat32(values ...float32) float32 {
	for _, v := range values {
		if v != 0 {
			return v
		}
	}
	return 0
}
func (*ternaryUtil) OrFloat64(values ...float64) float64 {
	for _, v := range values {
		if v != 0 {
			return v
		}
	}
	return 0
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
