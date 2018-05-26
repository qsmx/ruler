package ruler

import (
	"fmt"
	"reflect"
)

const (
	E_RULER_SYNTAX = 600 + iota
	E_RULER_DATA_DEEP
	E_RULER_DATA_NOT_MATCH
	E_DATA_INVALID
)

type Exception struct {
	code    int
	message string
}

type RulerException struct {
	Exception

	ruler  string
	buffer string
}

type DataException struct {
	Exception

	ruler string
	name  string
	value interface{}
}

func CatchException() {
	if err := recover(); err != nil {
		switch err.(type) {
		case *reflect.ValueError:
			throwException(E_DATA_INVALID, "不支持的数据类型")
		default:
			panic(err)
		}
	}
}

func ThrowException(code int, message string, rulers ...string) {
	e := Exception{code, message}
	var ex interface{} = e

	switch len(rulers) {
	case 1:
		ex = RulerException{Exception: e, ruler: rulers[0]}
	case 2:
		ex = RulerException{Exception: e, ruler: rulers[0], buffer: rulers[0]}
	case 3:
		ex = DataException{Exception: e, ruler: rulers[0], name: rulers[1], value: rulers[2]}
	}

	fmt.Printf("panic: %+v\n", ex)
	panic(ex)
}

func throwException(code int, message string, args ...interface{}) {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	fmt.Println("panic:", message)
	e := Exception{code, message}
	panic(e)
}
