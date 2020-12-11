/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/9/4
   Description :
-------------------------------------------------
*/

package zutils

import (
	"reflect"
	"testing"
)

func Test_sortUtil_Sort(t *testing.T) {
	tests := []struct {
		name    string
		values  []int
		reverse bool
		want    []int
	}{
		{"int", []int{2, 6, 1, 7, 2, 8, 4, 7, 9}, false, []int{1, 2, 2, 4, 6, 7, 7, 8, 9}},
		{"int_reverse", []int{2, 6, 1, 7, 2, 8, 4, 7, 9}, true, []int{9, 8, 7, 7, 6, 4, 2, 2, 1}},
	}

	so := &sortUtil{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			so.Sort(len(tt.values), func(i, j int) bool {
				return tt.values[i] < tt.values[j]
			}, func(i, j int) {
				tt.values[i], tt.values[j] = tt.values[j], tt.values[i]
			}, tt.reverse)
			if !reflect.DeepEqual(tt.values, tt.want) {
				t.Errorf("Sort() = %v, want %v", tt.values, tt.want)
			}
		})
	}
}
