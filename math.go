package zutils

var Math = new(mathUtil)

type mathUtil struct{}

// 获取最大值
func (*mathUtil) Max(a int, b ...int) int {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MaxInt(a int, b ...int) int {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MaxInt8(a int8, b ...int8) int8 {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MaxInt16(a int16, b ...int16) int16 {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MaxInt32(a int32, b ...int32) int32 {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MaxInt64(a int64, b ...int64) int64 {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MaxUint(a uint, b ...uint) uint {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MaxUint8(a uint8, b ...uint8) uint8 {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MaxUint16(a uint16, b ...uint16) uint16 {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MaxUint32(a uint32, b ...uint32) uint32 {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MaxUint64(a uint64, b ...uint64) uint64 {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MaxFloat32(a float32, b ...float32) float32 {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MaxFloat64(a float64, b ...float64) float64 {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}

// 获取最小值
func (*mathUtil) Min(a int, b ...int) int {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MinInt(a int, b ...int) int {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MinInt8(a int8, b ...int8) int8 {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MinInt16(a int16, b ...int16) int16 {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MinInt32(a int32, b ...int32) int32 {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MinInt64(a int64, b ...int64) int64 {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MinUint(a uint, b ...uint) uint {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MinUint8(a uint8, b ...uint8) uint8 {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MinUint16(a uint16, b ...uint16) uint16 {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MinUint32(a uint32, b ...uint32) uint32 {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MinUint64(a uint64, b ...uint64) uint64 {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MinFloat32(a float32, b ...float32) float32 {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}
func (*mathUtil) MinFloat64(a float64, b ...float64) float64 {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}
