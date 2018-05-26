package ruler

func Validate(src string) bool {
	return parse(src, nil) == 0
}

func Filter(src string, data interface{}) interface{} {
	return Hit0(src, NewDataPackage(data))
}

func Hit(src string, data interface{}) bool {
	v := Hit0(src, NewDataPackage(data))
	return ConvertBool(v)
}

func Hit0(src string, dp *DataPackage) interface{} {
	var n *node
	parse(src, &n)
	//n.Debug()
	return n.Exec(dp).Value
}

func parse(src string, node **node) int {
	return yyParse(newLexer(src), node)
}
