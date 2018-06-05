package ruler

import (
	"fmt"
	"reflect"
)

var typeName map[reflect.Kind]string

type rulerNode struct {
	/**
	 * 当前操作符
	 */
	op int

	/**
	 * 数据类型
	 */
	typ reflect.Kind

	/**
	 * 当前数据
	 */
	value interface{}

	/**
	 * 子节点
	 */
	node []*rulerNode
}

func nodeFromInterface(v interface{}) (r *rulerNode) {
	val := ValueFormat(v)

	switch val.(type) {
	case int64:
		r = nodeFromInt(val.(int64))
	case string:
		r = nodeFromString(STRING, val.(string))
	case float64:
		r = nodeFromFloat(val.(float64))
	case bool:
		r = nodeFromBool(val.(bool))
	default:
		kind := reflect.ValueOf(val).Kind()
		switch kind {
		case reflect.Slice, reflect.Array, reflect.Map:
			r = &rulerNode{
				typ:   kind,
				value: val,
			}
		default:
			throwException(500, "不支持的自动转换 %T, %+v", val, val)
		}
	}

	return r
}

func nodeFromNode(op int, n ...*rulerNode) *rulerNode {
	r := &rulerNode{
		op:   op,
		node: n,
		typ:  reflect.Slice,
	}
	return r
}

func nodeFromIdent(op int, ident string) *rulerNode {
	return &rulerNode{
		op:    op,
		value: ident,
	}
}

func nodeFromInt(value int64) *rulerNode {
	return &rulerNode{
		value: value,
		typ:   reflect.Int64,
	}
}

func nodeFromFloat(value float64) *rulerNode {
	return &rulerNode{
		value: value,
		typ:   reflect.Float64,
	}
}

func nodeFromString(op int, value string) *rulerNode {
	return &rulerNode{
		op:    op,
		value: value,
		typ:   reflect.String,
	}
}

func nodeFromBool(value bool) *rulerNode {
	return &rulerNode{
		value: value,
		typ:   reflect.Bool,
	}
}

func (r *rulerNode) AppendNode(n ...*rulerNode) *rulerNode {
	r.node = append(r.node, n...)
	return r
}

func (r *rulerNode) ReplaceNode(n *rulerNode) *rulerNode {
	r.node = n.node
	return r
}

func (r rulerNode) Len() int {
	return len(r.node)
}

func (r rulerNode) Value() interface{} {
	return r.value
}

func calc(op int, l, r interface{}) *rulerNode {
	var result interface{}

	switch l.(type) {
	case int64:
		m := l.(int64)
		n := r.(int64)

		switch op {
		case '+':
			result = m + n
		case '-':
			result = m - n
		case '*':
			result = m * n
		case '/':
			result = m / n
		case '%':
			result = m % n
		case LSS:
			result = m < n
		case LEQ:
			result = m <= n
		case GTR:
			result = m > n
		case GEQ:
			result = m >= n
		}
	case float64:
		m := l.(float64)
		n := r.(float64)

		switch op {
		case '+':
			result = m + n
		case '-':
			result = m - n
		case '*':
			result = m * n
		case '/':
			result = m / n
		case LSS:
			result = m < n
		case LEQ:
			result = m <= n
		case GTR:
			result = m > n
		case GEQ:
			result = m >= n
		}
	case string:
		m := l.(string)
		n := r.(string)

		switch op {
		case '+':
			result = m + n
		case LSS:
			result = m < n
		case LEQ:
			result = m <= n
		case GTR:
			result = m > n
		case GEQ:
			result = m >= n
		}
	}

	switch result.(type) {
	case bool:
		return nodeFromBool(result.(bool))
	case int64:
		return nodeFromInt(result.(int64))
	case float64:
		return nodeFromFloat(result.(float64))
	case string:
		return nodeFromString(STRING, result.(string))
	}

	throwException(500, "不支持的运算类型 %T %d[%s] %T, result = %T, %v", l, op, yySymNames[yyXLAT[op]], r, result, result)
	return nil
}

func (r *rulerNode) Exec(dp *DataPackage) *rulerNode {
	switch r.op {
	case IDENT:
		var val interface{}
		s := r.value.(string)
		if s[0] == '|' {
			val = dp.GetBaseAttr(s[1:])
		} else {
			val = dp.GetAttr(s)
		}

		//fmt.Printf("Ident: %T, %v\n", val, val)
		return nodeFromInterface(val)
	case ARRAYS:
		val := dp.GetDeepAttr(r.value.(string))
		//fmt.Printf("Arrays: %T, %v\n", val, val)
		return nodeFromInterface(val)

	case NEG:
		var val interface{}
		switch r.typ {
		case reflect.Int64:
			val = -r.value.(int64)
		case reflect.Float64:
			val = -r.value.(float64)
		default:
			throwException(500, "一元减不支持的类型 %T", r.value)
		}
		return nodeFromInterface(val)
	case EQU:
		r, v := r.node[0].Exec(dp), r.node[1].Exec(dp)
		return nodeFromBool(r.value == v.value)
	case NEQ:
		r, v := r.node[0].Exec(dp), r.node[1].Exec(dp)
		return nodeFromBool(r.value != v.value)

	case '+', '-', '*', '/', '%', LSS, LEQ, GTR, GEQ:
		left, right := r.node[0].Exec(dp), r.node[1].Exec(dp)

		if left.typ != right.typ {
			switch left.typ {
			case reflect.Int64:
				if right.typ == reflect.Float64 {
					left = nodeFromFloat(float64(left.value.(int64)))
				} else {
					right = nodeFromInt(ConvertInt(right.value))
				}
			case reflect.Float64:
				right = nodeFromFloat(ConvertFloat(right.value))
			case reflect.String:
				right = nodeFromString(STRING, ConvertString(right.value))
			}
		}

		return calc(r.op, left.value, right.value)

	case AND:
		return nodeFromBool(ConvertBool(r.node[0].Exec(dp).value) && ConvertBool(r.node[1].Exec(dp).value))
	case OR:
		return nodeFromBool(ConvertBool(r.node[0].Exec(dp).value) || ConvertBool(r.node[1].Exec(dp).value))

	case FUNC:
		switch r.value {
		case "int":
			return nodeFromInt(ConvertInt(r.node[0].Exec(dp).value))
		case "float":
			return nodeFromFloat(ConvertFloat(r.node[0].Exec(dp).value))
		case "bool":
			return nodeFromBool(ConvertBool(r.node[0].Exec(dp).value))
		case "string":
			return nodeFromString(STRING, ConvertString(r.node[0].Exec(dp).value))
		default:
			var args []interface{}
			for _, node := range r.node {
				args = append(args, node.Exec(dp).value)
			}

			res := Call(r.value.(string), args...)
			return nodeFromInterface(res)
		}

	case FILTER:
		name := r.node[0].value
		data := reflect.ValueOf(r.node[0].Exec(dp).value)
		if data.Kind() != reflect.Slice {
			throwException(500, "%s 不是数组数据，无法过滤", name)
		}

		size := data.Len()
		res := make([]bool, size)
		for i := 0; i < size; i++ {
			dp.Push(data.Index(i).Interface())
			res[i] = ConvertBool(r.node[1].Exec(dp).value)
			dp.Pop()
		}

		return nodeFromInterface(res)

	case LISTA:
		if ConvertBool(r.node[0].Exec(dp).value) {
			return r.node[1].Exec(dp)
		} else {
			return nodeFromBool(false)
		}

	case LISTO:
		if ConvertBool(r.node[0].Exec(dp).value) {
			return nodeFromBool(true)
		} else {
			return r.node[1].Exec(dp)
		}

	case IN:
		val := r.node[0].Exec(dp).value

		for _, v := range r.node[1].node {
			if val == v.Exec(dp).value {
				return nodeFromBool(r.value.(string) == "i")
			}
		}

		return nodeFromBool(r.value.(string) == "n")
	}

	return r
}

func inArray(v interface{}, c []interface{}) bool {
	if len(c) == 0 {
		return false
	}
	for _, val := range c {
		if v == val {
			return true
		}
	}

	return false
}

func (r rulerNode) Debug(level ...int) {
	var (
		v  string
		ok bool
	)

	lev := 0
	if len(level) != 0 {
		lev = level[0]
	}
	pre := fmt.Sprintf("%*s", 4*lev, "")

	op := yySymNames[yyXLAT[r.op]]
	if v, ok = typeName[r.typ]; ok {
		fmt.Printf("%s%s: op[%s] %v\n", pre, v, op, r.value)
	} else if r.typ == reflect.Slice {
		fmt.Printf("%sSlice: op[%s]\n", pre, yySymNames[yyXLAT[r.op]])
	}

	for i := 0; i < r.Len(); i++ {
		r.node[i].Debug(lev + 1)
	}
}

func init() {
	typeName = map[reflect.Kind]string{
		reflect.Invalid: "ident",
		reflect.Int64:   "int64",
		reflect.Float64: "float64",
		reflect.Bool:    "bool",
		reflect.String:  "string",
	}
}
