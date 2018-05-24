package ruler

import (
	"fmt"
)

func Validate(src string) (e RulerException, ok bool) {
	defer func() {
		switch err := recover(); err.(type) {
		case nil:
			return
		case *RulerException:
			e = *err.(*RulerException)
		case error:
			e = RulerException{code: 500, str: err.(error).Error()}
		case string:
			e = RulerException{code: 500, str: err.(string)}
		default:
			e = RulerException{code: 500, str: fmt.Sprintf("未知错误 %T, %v", err, err)}
		}

		ok = false
	}()

	ok = parse(src, nil) == 0

	return
}

func Filter(src string, data interface{}) interface{} {
	_, res := hit(src, NewDataPackage(data))
	return res
}

func HilAll(src []string, data interface{}) (e RulerException, ok bool) {
	dp := NewDataPackage(data)

	for _, ruler := range src {
		if e, ok = Hit0(ruler, dp); ok {
			return
		}
	}

	return e, false
}

func Hit(src string, data interface{}) (e RulerException, ok bool) {
	return Hit0(src, NewDataPackage(data))
}

func Hit0(src string, dp *DataPackage) (e RulerException, ok bool) {
	var res interface{}
	e, res = hit(src, dp)
	return e, ConvertBool(res)
}

func hit(src string, dp *DataPackage) (e RulerException, ok interface{}) {
	defer func() {
		switch err := recover(); err.(type) {
		case nil:
			return
		case *RulerException:
			e = *err.(*RulerException)
		case error:
			e = RulerException{code: 500, str: err.(error).Error()}
		case string:
			e = RulerException{code: 500, str: err.(string)}
		default:
			e = RulerException{code: 500, str: fmt.Sprintf("未知错误 %T, %v", err, err)}
		}

		ok = false
	}()

	ok = Parse0(src, dp)

	return
}

func Parse0(src string, dp *DataPackage) interface{} {
	var node *rulerNode

	if parse(src, &node) != 0 {
		throwException(500, "规则与数据不匹配或规则格式错误", src)
	}

	return node.Exec(dp).Value()
}

func parse(src string, node **rulerNode) int {
	return yyParse(newLexer(src), node)
}
