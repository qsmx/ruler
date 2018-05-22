package ruler

import "fmt"

func Validate(src string) (RulerException, bool) {
	return parse(src, nil)
}

func Parse(src string) (e RulerException, ok bool) {
	var node *rulerNode

	if e, ok = parse(src, &node); !ok {
		return
	}

	node.Debug()
	return e, true
}

func parse(src string, node **rulerNode) (e RulerException, ok bool) {
	//e = &RulerException{}
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

		fmt.Println(e, &e)
		ok = false
	}()

	ok = yyParse(newLexer(src), node) == 0

	return
}
