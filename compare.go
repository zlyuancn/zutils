package zutils

// Deprecated: 建议使用 lo.Contains  eg: https://github.com/samber/lo
var Compare = new(compareUtil)

type compareUtil struct{}

// 检查 a 是否在 values 中
func (*compareUtil) CheckIn(a interface{}, values ...interface{}) bool {
	for _, v := range values {
		if v == a {
			return true
		}
	}
	return false
}

// 检查 a 是否在 values 中
func (*compareUtil) CheckInInt(a int, values ...int) bool {
	for _, v := range values {
		if v == a {
			return true
		}
	}
	return false
}

// 检查 a 是否在 values 中
func (*compareUtil) CheckInString(a string, values ...string) bool {
	for _, v := range values {
		if v == a {
			return true
		}
	}
	return false
}

// 比较两个字符串数组切片是否相等, 注意: 不会对nil值做比较, 也就是说 []string{} == nil
func (*compareUtil) CompareStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
