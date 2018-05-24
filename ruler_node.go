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
		throwException(500, "不支持的自动转换 %T, %+v", val, val)
	}

	return r
}

func nodeFromNode(op int, n ...*rulerNode) *rulerNode {
	r := &rulerNode{
		op: op,
		node: n,
		typ: reflect.Slice,
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
		op: op,
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
	case int64:
		return nodeFromInt(result.(int64))
	case float64:
		return nodeFromFloat(result.(float64))
	case string:
		return nodeFromString(STRING, result.(string))
	}

	throwException(500, "不支持的运算类型 %T %c[%d] %T, result = %T, %v", l, op, op, r, result, result)
	return nil
}

func (r *rulerNode) Exec(dp *DataPackage) (* rulerNode) {
    switch r.op {
	case IDENT:
		r.value = ValueFormat(dp.GetAttr(r.value.(string)))
		fmt.Println("Ident:", r.value)
	case ARRAYS:
		r.value = ValueFormat(dp.GetDeepAttr(r.value.(string)))
		fmt.Println("Arrays:", r.value)
	case NEG:
		switch r.typ {
		case reflect.Int64:
			r.value = -r.value.(int64)
		case reflect.Float64:
			r.value = -r.value.(float64)
		default:
			throwException(500, "一元减不支持的类型 %T", r.value)
		}
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
					left = nodeFromFloat(float64(r.value.(int64)))
				} else {
					right = nodeFromInt(ConvertInt(right.value))
				}
			case reflect.Float64:
				right = nodeFromFloat(ConvertFloat(left.value))
			case reflect.String:
				right = nodeFromString(STRING, ConvertString(left.value))
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
		fmt.Println("Filter: ", r.node[0], "Args:", r.node[1:])
    }

    return r
}

func syncType(a, b interface{}) (m, n interface{}, ok bool) {
	m, ais := a.(float64)
	n, bis := b.(float64)
	if ok = ais || bis; ok {
		if !ais {
			m = ConvertFloat(a)
		}
		if !bis {
			n = ConvertFloat(b)
		}
	}

	return
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
    pre := fmt.Sprintf("%*s", 4 * lev, "")

    op := yySymNames[yyXLAT[r.op]]
	if v, ok = typeName[r.typ]; ok {
		fmt.Printf("%s%s: op[%s] %v\n", pre, v, op, r.value)
	} else if r.typ == reflect.Slice {
		fmt.Printf("%sSlice: op[%s]\n", pre, op)
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
