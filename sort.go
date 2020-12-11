/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/9/4
   Description :
-------------------------------------------------
*/

package zutils

import (
	"sort"
)

var Sort = new(sortUtil)

type sortUtil struct{}

var _ sort.Interface = (*sortAny)(nil)

// 比较索引i的数据比j的数据小
type LessFunc func(i, j int) bool

// 互换索引i和j的数据
type SwapFunc func(i, j int)

type sortAny struct {
	count   int
	less    LessFunc
	swap    SwapFunc
	reverse bool
}

func (s *sortAny) Len() int {
	return s.count
}
func (s *sortAny) Less(i, j int) bool {
	return s.reverse != s.less(i, j)
}
func (s *sortAny) Swap(i, j int) {
	s.swap(i, j)
}

// 排序(数量, 比较函数, 交换函数, 是否倒序)
func (*sortUtil) Sort(count int, less LessFunc, swap SwapFunc, reverse ...bool) {
	s := &sortAny{
		count:   count,
		less:    less,
		swap:    swap,
		reverse: len(reverse) > 0 && reverse[0],
	}
	sort.Sort(s)
}
