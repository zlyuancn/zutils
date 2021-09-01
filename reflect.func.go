package zutils

import (
	"fmt"
	"reflect"
)

type ReflectMethod struct {
	Name        string          // 方法名
	MethodValue *reflect.Value  // 反射的方法实体
	Method      *reflect.Method // 反射的方法
}

func (r *ReflectMethod) call(callMethod func(in []reflect.Value) []reflect.Value, in []interface{}) []interface{} {
	input := make([]reflect.Value, len(in))
	for i, v := range in {
		input[i] = reflect.ValueOf(v)
	}

	result := callMethod(input)
	output := make([]interface{}, len(result))
	for i, r := range result {
		output[i] = r.Interface()
	}
	return output
}
func (r *ReflectMethod) callValueN(callMethdo func(in []reflect.Value) []reflect.Value, n int, in ...interface{}) []interface{} {
	result := r.call(callMethdo, in)
	if len(result) != n {
		panic(fmt.Errorf("%s: expect len(result) = %d, but got %d", r.Name, n, len(result)))
	}
	return result
}

func (r *ReflectMethod) Call(in ...interface{}) []interface{} {
	return r.call(r.MethodValue.Call, in)
}
func (r *ReflectMethod) Call1(in ...interface{}) interface{} {
	result := r.call(r.MethodValue.Call, in)
	return result[0]
}
func (r *ReflectMethod) Call2(in ...interface{}) (interface{}, interface{}) {
	result := r.call(r.MethodValue.Call, in)
	return result[0], result[1]
}
func (r *ReflectMethod) Call3(in ...interface{}) (interface{}, interface{}, interface{}) {
	result := r.call(r.MethodValue.Call, in)
	return result[0], result[1], result[2]
}

func (r *ReflectMethod) CallSlice(in ...interface{}) []interface{} {
	return r.call(r.MethodValue.CallSlice, in)
}
func (r *ReflectMethod) CallSlice1(in ...interface{}) interface{} {
	result := r.call(r.MethodValue.CallSlice, in)
	return result[0]
}
func (r *ReflectMethod) CallSlice2(in ...interface{}) (interface{}, interface{}) {
	result := r.call(r.MethodValue.CallSlice, in)
	return result[0], result[1]
}
func (r *ReflectMethod) CallSlice3(in ...interface{}) (interface{}, interface{}, interface{}) {
	result := r.call(r.MethodValue.CallSlice, in)
	return result[0], result[1], result[2]
}

// GetMethods 获取a的方法列表, a必须是一个带指针的结构体
//
//  example:
//  type A struct{}
//  func (*A) A(a int) int { return a + 1 }
//  func (A) B(b int) int  { return b + b }
//  func main() {
//      methods := zutils.Reflect.GetMethods(new(A))
//      fmt.Println(methods["A"].Call1(123)) // 124
//      fmt.Println(methods["B"].Call1(123)) // 246
//  }
func (*reflectUtil) GetMethods(a interface{}) map[string]*ReflectMethod {
	aType := reflect.TypeOf(a)
	if aType.Kind() != reflect.Ptr {
		panic("a must be a structure with a pointer")
	}
	aTypeElem := aType.Elem()
	if aTypeElem.Kind() == reflect.Invalid {
		panic("a is invalid, a may be not initialized")
	}
	if aTypeElem.Kind() != reflect.Struct {
		panic("a must a struct")
	}

	aValue := reflect.ValueOf(a)
	out := make(map[string]*ReflectMethod)
	for i := 0; i < aType.NumMethod(); i++ {
		method := aType.Method(i)
		if method.PkgPath != "" {
			continue
		}

		methodValue := aValue.Method(i)
		out[method.Name] = &ReflectMethod{
			Name:        method.Name,
			MethodValue: &methodValue,
			Method:      &method,
		}
	}
	return out
}
