/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/12/11
   Description :
-------------------------------------------------
*/

package zutils

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"unsafe"

	"github.com/vmihailenco/msgpack"
)

var Reflect = new(reflectUtil)

type reflectUtil struct{}

// 深拷贝, dst必须传入指针
func (*reflectUtil) DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	err := msgpack.NewEncoder(&buf).Encode(src)
	if err != nil {
		return err
	}

	return msgpack.NewDecoder(bytes.NewReader(buf.Bytes())).Decode(dst)
}

// 从获取结构体的字段名, 如果tag存在则优先取tag的值
func (*reflectUtil) GetStructFields(a interface{}, tag string) ([]string, error) {
	rt := reflect.TypeOf(a)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	} else if rt.Kind() == reflect.Interface {
		rt = rt.Elem()
	}

	if rt.Kind() != reflect.Struct {
		return nil, errors.New("input value must a struct type.")
	}

	fieldCount := rt.NumField()
	keys := make([]string, 0, fieldCount)
	for i := 0; i < fieldCount; i++ {
		field := rt.Field(i)
		if field.PkgPath != "" { // 未导出的
			continue
		}

		k := field.Tag.Get(tag)
		if k == "" {
			k = field.Name
		}
		keys = append(keys, k)
	}
	return keys, nil
}

// 判断传入参数是否为该类型的零值
func (*reflectUtil) IsZero(a interface{}) bool {
	switch v := a.(type) {

	case nil:
		return true

	case string:
		return v == ""
	case []byte:
		return len(v) == 0
	case bool:
		return v

	case int:
		return v == 0
	case int8:
		return v == 0
	case int16:
		return v == 0
	case int32:
		return v == 0
	case int64:
		return v == 0

	case uint:
		return v == 0
	case uint8:
		return v == 0
	case uint16:
		return v == 0
	case uint32:
		return v == 0
	case uint64:
		return v == 0

	case float32:
		return v == 0
	case float64:
		return v == 0
	}

	r_v := reflect.Indirect(reflect.ValueOf(a))

	switch r_v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return r_v.Len() == 0
	case reflect.Invalid:
		return true
	case reflect.Chan, reflect.UnsafePointer:
		return r_v.IsNil()
	}

	nv := reflect.New(r_v.Type()).Elem().Interface()
	return r_v.Interface() == nv
}

// 对未导出字段解锁, 执行完设置函数后重新上锁
func (*reflectUtil) UnlockUnexported(v *reflect.Value, fn func()) {
	const flagRO uintptr = 96
	if v.CanSet() {
		fn()
	} else {
		o_vf := reflect.ValueOf(v).Elem().Field(2).UnsafeAddr() // 获取Value.flag变量的地址
		raw_f := *(*uintptr)(unsafe.Pointer(o_vf))              // 获取标记的值
		f := raw_f
		f ^= f & flagRO                       // 设置为导出字段
		*(*uintptr)(unsafe.Pointer(o_vf)) = f // 重新设置标记
		defer func() {
			*(*uintptr)(unsafe.Pointer(o_vf)) = raw_f // 还原标记
		}()
		fn()
	}
}

func (*reflectUtil) mustTo(v interface{}, kind reflect.Kind) *reflect.Value {
	v_value := reflect.ValueOf(v)
	for v_value.Kind() == reflect.Ptr || v_value.Kind() == reflect.Interface {
		v_value = v_value.Elem()
	}
	if v_value.Kind() != kind {
		panic(fmt.Sprintf("input value must a %s type, but got %s type.", kind.String(), v_value.Kind().String()))
	}

	return &v_value
}

// 将struct转为map, 即使struct的字段未导出也会转换, 注意: 它的值不是深拷贝
func (u *reflectUtil) StructToMap(v interface{}) map[string]interface{} {
	v_value := u.mustTo(v, reflect.Struct)
	v_type := v_value.Type()

	num := v_value.NumField()
	out := make(map[string]interface{}, num)
	for i := 0; i < num; i++ {
		field := v_type.Field(i)
		value := reflect.NewAt(field.Type, unsafe.Pointer(v_value.Field(i).UnsafeAddr()))
		out[field.Name] = value.Elem().Interface()
	}

	return out
}

// 将map转为struct, 即使struct的字段未导出也会转换, 注意: 它的值不是深拷贝
func (u *reflectUtil) MapToStruct(m interface{}, out interface{}) {
	m_value := u.mustTo(m, reflect.Map)
	o_value := u.mustTo(out, reflect.Struct)
	o_type := o_value.Type()

	num := o_value.NumField()
	for i := 0; i < num; i++ {
		o_v := o_value.Field(i)
		o_t := o_type.Field(i)

		m_v := m_value.MapIndex(reflect.ValueOf(o_t.Name))
		if m_v.Kind() == reflect.Invalid { // map中无此字段
			continue
		}
		if m_v.Kind() == reflect.Interface {
			m_v = m_v.Elem()
		}

		u.UnlockUnexported(&o_v, func() {
			o_v.Set(m_v)
		})
	}

}
