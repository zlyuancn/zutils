/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/5/7
   Description :
-------------------------------------------------
*/

package zutils

import (
	"fmt"
)

// 如果err不为空则panic
func PanicOnError(err error, msg string) {
	if err != nil {
		panic(fmt.Errorf("%s: %s", msg, err))
	}
}
