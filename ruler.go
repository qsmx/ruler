package ruler

func Parse(src string, data map[string]interface{}) bool {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
			return
		}
	}()
    var res bool
    rulerParse(newLexer(src), NewDataPackage(data), &res)
    fmt.Println("res:", res)
    return res
}