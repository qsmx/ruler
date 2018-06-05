package ruler

import (
	"reflect"
)

type funcsCallback map[string]reflect.Value

var funcs funcsCallback

var keywords map[string]bool

func Bind(name string, fn interface{}) (err error) {
	if _, ok := keywords[name]; ok {
		throwException(500, "不能注册函数 %s", name)
	}

	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		throwException(500, "%s 不是函数", name)
	}

	if n := v.Type().NumOut(); n != 1 {
		throwException(500, "注册函数必须返回一个值 %s: %d", name, n)
	}
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
	v := Call(name, params...)
	return v[0].Interface()
}

func init() {
	funcs = make(funcsCallback)
	keywords = map[string]bool {
		"int": true,
		"float": true,
		"string": true,
		"bool": true,
		"in": true,
		"notin": true,
	}
}
