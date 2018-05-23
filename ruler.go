package ruler

import "fmt"

func Validate(src string) (RulerException, bool) {
	return parse(src, nil)
}

func ParseAll(src []string, data interface{}) (e interface{}, ok bool) {
	var node *rulerNode
	dp := NewDataPackage(data)

	for _, ruler := range src {
		if e, ok = parse(ruler, &node); !ok {
			return
		}

		//node.Debug()
		node.Exec(dp)
	}

	return e, true
}

func Parse(src string, data interface{}) (e interface{}, ok bool) {
	return parseone(src, NewDataPackage(data))
}

func parseone(src string, dp *DataPackage) (e interface{}, ok bool) {
	var node *rulerNode

	if e, ok = parse(src, &node); !ok {
		return
	}
	//
	//defer func() {
	//	switch err := recover(); err.(type) {
	//	case nil:
	//		return
	//	case *RulerException:
	//		e = *err.(*RulerException)
	//	case error:
	//		e = RulerException{code: 500, str: err.(error).Error()}
	//	case string:
	//		e = RulerException{code: 500, str: err.(string)}
	//	default:
	//		e = RulerException{code: 500, str: fmt.Sprintf("未知错误 %T, %v", err, err)}
	//	}
	//
	//	ok = false
	//}()

	res := node.Exec(dp)
	fmt.Println(" *** RES = ", res.Value())
	return e, true
}

func parse(src string, node **rulerNode) (e RulerException, ok bool) {
	//defer func() {
	//	switch err := recover(); err.(type) {
	//	case nil:
	//		return
	//	case *RulerException:
	//		e = *err.(*RulerException)
	//	case error:
	//		e = RulerException{code: 500, str: err.(error).Error()}
	//	case string:
	//		e = RulerException{code: 500, str: err.(string)}
	//	default:
	//		e = RulerException{code: 500, str: fmt.Sprintf("未知错误 %T, %v", err, err)}
	//	}
	//
	//	ok = false
	//}()

	ok = yyParse(newLexer(src), node) == 0

	return
}
