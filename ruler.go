package ruler

import "fmt"

func Validate(src string) bool {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
			return
		}
	}()
	return yyParse(newLexer(src), nil) == 0
}

func Parse(src string, data map[string]interface{}) bool {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
			return
		}
	}()

	var node *rulerNode
	yyParse(newLexer(src), &node)
	if node == nil {
		return false
	}
	node.Debug()
	return true
}
