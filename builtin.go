package ruler

import (
	"reflect"
)

type funcsCallback map[string]reflect.Value

var funcs funcsCallback

func Bind(name string, fn interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			throwException(500, "%s 不是函数", name)
		}
	}()

	v := reflect.ValueOf(fn)
	v.Type().NumIn()
	funcs[name] = v

	return
}

func Call(name string, params ...interface{}) (result []reflect.Value) {
	if _, ok := funcs[name]; !ok {
		throwException(500, "%s 函数不存在", name)
		return
	}
	if len(params) != funcs[name].Type().NumIn() {
		throwException(500, "%s 函数参数个数与调用不一致", name)
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = funcs[name].Call(in)

	return
}

func Caller(name string, params ...interface{}) interface{} {
	v := Call(name, params ...)
	return v[0].Interface()
}

func init() {
	funcs = make(funcsCallback)
}
