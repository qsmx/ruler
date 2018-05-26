package ruler

import "fmt"

func Validate(src string) bool {
	return parse(src, nil) == 0
}

func Hit(src string, data interface{}) bool {
	v := Hit0(src, NewDataPackage(data))
	fmt.Println(v)
	return ConvertBool(v)
}

func Hit0(src string, dp *DataPackage) interface{} {
	var n *node
	parse(src, &n)
	n.Debug()
	k := n.Exec(dp)

	return k.Value
}

func parse(src string, node **node) int {
	return yyParse(newLexer(src), node)
}
