package zutils

import (
	"fmt"
)

var Assert = new(assertUtil)

type assertUtil struct{}

func (u *assertUtil) True(a bool, msg ...interface{}) {
	if !a {
		u.raise("assert.True", msg)
	}
}
func (u *assertUtil) False(a bool, msg ...interface{}) {
	if a {
		u.raise("assert.False", msg)
	}
}

func (u *assertUtil) Equal(a, b interface{}, msg ...interface{}) {
	if a != b {
		u.raise("assert.Equal", msg)
	}
}
func (u *assertUtil) NotEqual(a, b interface{}, msg ...interface{}) {
	if a == b {
		u.raise("assert.NotEqual", msg)
	}
}

func (u *assertUtil) Nil(a interface{}, msg ...interface{}) {
	if a != nil {
		u.raise("assert.Nil", msg)
	}
}
func (u *assertUtil) NotNil(a interface{}, msg ...interface{}) {
	if a == nil {
		u.raise("assert.NotNil", msg)
	}
}

func (u *assertUtil) Zero(a interface{}, msg ...interface{}) {
	if !Reflect.IsZero(a) {
		u.raise("assert.Zero", msg)
	}
}
func (u *assertUtil) NotZero(a interface{}, msg ...interface{}) {
	if Reflect.IsZero(a) {
		u.raise("assert.NotZero", msg)
	}
}

func (u *assertUtil) raise(def string, msg []interface{}) {
	if len(msg) == 0 {
		panic(def)
	} else if len(msg) == 1 {
		panic(def + " " + msg[0].(string))
	} else {
		panic(def + " " + fmt.Sprintf(msg[0].(string), msg[1:]...))
	}
}
