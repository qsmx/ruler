package ruler

import (
	"reflect"
	"fmt"
)

func Validate(src string) bool {
	 _, ok := Hit0(src, nil)
	 return ok
}

func Filter(src string, data interface{}) (res interface{}, ok bool) {
	return Hit0(src, NewDataPackage(data))
}

func Hit(src string, data interface{}) (res interface{}, ok bool) {
	res, ok = Hit0(src, NewDataPackage(data))
	if ok {
		res = ConvertBool(res)
	}

	return
}

func Hit0(src string, dp *DataPackage) (res interface{}, ok bool) {
	var n *node

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case string:
				res = &Exception{1000, "未知错误：" + err.(string)}
			case *Exception, *RulerException:
				res = err
			case *reflect.ValueError:
				res = &Exception{code: E_DATA_INVALID, message: "不支持的数据类型"}
			default:
				fmt.Println(err)
			}

			fmt.Println(res)
			ok = false
		}
	}()

	if ok = parse(src, &n) == 0; ok {
		//n.Debug()

		if dp != nil {
			res = n.Exec(dp).Value
		} else {
			res = nil
		}
	}

	return
}

func parse(src string, node **node) int {
	return yyParse(newLexer(src), node)
}
