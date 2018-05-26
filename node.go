package ruler

import (
	"reflect"
	"fmt"
)

type node struct {
	// name
	Label string

	// Value
	Value interface{}

	// operation code
	op int

	// kind
	kind reflect.Kind

	// chilren
	node []*node
}

func nodeFromAny(label string, val interface{}) (r *node) {
	v := ValueFormat(val)

	switch v.(type) {
	case int:
		r = nodeFromInt(label, v.(int))

	case bool:
		r = nodeFromBool(label, v.(bool))

	case float64:
		r = nodeFromFloat(label, v.(float64))

	case string:
		r = nodeFromString(label, v.(string))

	default:
		r = &node{
			Label: label,
			Value: val,
		}
	}

	return r
}

func nodeFromInt(label string, val int) *node {
	return &node{
		Label: label,
		Value: val,
		kind:  reflect.Int,
	}
}

func nodeFromFloat(label string, val float64) *node {
	return &node{
		Label: label,
		Value: val,
		kind:  reflect.Float64,
	}
}

func nodeFromBool(label string, val bool) *node {
	return &node{
		Label: label,
		Value: val,
		kind:  reflect.Bool,
	}
}

func nodeFromString(label string, val string) *node {
	return &node{
		Label: label,
		Value: val,
		kind:  reflect.String,
	}
}

func nodeFromIdent(op int, label string) *node {
	return &node{
		Label: label,
		op:    op,
	}
}

func nodeFromNode(op int, ns ...*node) *node {
	return nodeFromNode2(op, "", ns...)
}

func nodeFromNode2(op int, label string, ns ...*node) *node {
	c := &node{
		op:    op,
		Label: label,
		kind:  reflect.Slice,
	}

	c.node = append(c.node, ns...)
	return c
}

func nodeFromFunc(op int, name string, ns ...*node) *node {
	c := &node{
		op:    op,
		Label: name,
		kind:  reflect.Func,
	}

	c.node = append(c.node, ns...)
	return c
}

func (n *node) AppendNode(ns ...*node) *node {
	n.node = append(n.node, ns...)
	return n
}

func (r node) Len() int {
	return len(r.node)
}


func calc(op int, l, r interface{}) *node {
	var result interface{}

	switch l.(type) {
	case int:
		m := l.(int)
		n := r.(int)

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
		case BITAND:
			result = m & n
		case BITOR:
			result = m | n
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

	return nodeFromAny("", result)

	// throwException(500, "不支持的运算类型 %T %d[%s] %T, result = %T, %v", l, op, yySymNames[yyXLAT[op]], r, result, result)
	return nil
}

func (r *node) Exec(dp *DataPackage) *node {
	switch r.op {
	case IDENT:
		if v, ok := dp.GetAttr(r.Label); ok {
			return nodeFromAny(r.Label, v)
		} else {
			throwException(E_DATA_INVALID, "找不到 %s", r.Label)
		}
	case ARRAY:
		if v, ok := dp.GetDeepAttr(r.Label); ok {
			return nodeFromAny(r.Label, v)
		} else {
			throwException(E_DATA_INVALID, "找不到 %s", r.Label)
		}

	case NEG:
		var val interface{}
		switch r.kind {
		case reflect.Int:
			val = -r.Value.(int)
		case reflect.Float64:
			val = -r.Value.(float64)
		default:
			throwException(500, "一元减不支持的类型 %T", r.Value)
		}
		return nodeFromAny("", val)
	case EQU:
		r, v := r.node[0].Exec(dp), r.node[1].Exec(dp)
		return nodeFromBool("", r.Value == v.Value)
	case NEQ:
		r, v := r.node[0].Exec(dp), r.node[1].Exec(dp)
		return nodeFromBool("", r.Value != v.Value)

	case '+', '-', '*', '/', '%', BITAND, BITOR, LSS, LEQ, GTR, GEQ:
		left, right := r.node[0].Exec(dp), r.node[1].Exec(dp)

		if left.kind != right.kind {
			switch left.kind {
			case reflect.Int:
				if right.kind == reflect.Float64 {
					left = nodeFromFloat("", float64(left.Value.(int)))
				} else {
					right = nodeFromInt("", ConvertInt(right.Value))
				}
			case reflect.Float64:
				right = nodeFromFloat("", ConvertFloat(right.Value))
			case reflect.String:
				right = nodeFromString("", ConvertString(right.Value))
			}
		}

		return calc(r.op, left.Value, right.Value)

	case AND:
		return nodeFromBool("", ConvertBool(r.node[0].Exec(dp).Value) && ConvertBool(r.node[1].Exec(dp).Value))
	case OR:
		return nodeFromBool("", ConvertBool(r.node[0].Exec(dp).Value) || ConvertBool(r.node[1].Exec(dp).Value))

	case FUNC:
		switch r.Value {
		case "int":
			return nodeFromInt("", ConvertInt(r.node[0].Exec(dp).Value))
		case "float":
			return nodeFromFloat("", ConvertFloat(r.node[0].Exec(dp).Value))
		case "bool":
			return nodeFromBool("", ConvertBool(r.node[0].Exec(dp).Value))
		case "string":
			return nodeFromString("", ConvertString(r.node[0].Exec(dp).Value))
		default:
			var args []interface{}
			for _, node := range r.node {
				args = append(args, node.Exec(dp).Value)
			}

			res := Call(r.Value.(string), args...)
			return nodeFromAny("", res)
		}

	case FILTER:
		fmt.Println(*r.node[0])
		n := r.node[0]
		name := n.Label
		data := reflect.ValueOf(n.Exec(dp).Value)
		fmt.Printf("data: %+v\n", data)
		if data.Kind() != reflect.Slice {
			throwException(500, "%s 不是数组数据，无法过滤", name)
		}

		size := data.Len()
		res := make([]bool, size)
		for i := 0; i < size; i++ {
			dp.Push(data.Index(i).Interface())
			res[i] = ConvertBool(r.node[1].Exec(dp).Value)
			dp.Pop()
		}

		return nodeFromAny("", res)

	case LISTA:
		if ConvertBool(r.node[0].Exec(dp).Value) {
			return r.node[1].Exec(dp)
		} else {
			return nodeFromBool("", false)
		}

	case LISTO:
		if ConvertBool(r.node[0].Exec(dp).Value) {
			return nodeFromBool("", true)
		} else {
			return r.node[1].Exec(dp)
		}

	case IN:
		val := r.node[0].Exec(dp).Value

		for _, v := range r.node[1].node {
			if val == v.Exec(dp).Value {
				return nodeFromBool("", r.Label == "i")
			}
		}

		return nodeFromBool("", r.Label == "n")
	}

	return r
}

func (r node) Debug() {
	r.debug(0)
}

func (r node) debug(lev int) {
	var (
		v  string
		ok bool
	)

	pre := fmt.Sprintf("%*s", 4*lev, "")

	typeName := map[reflect.Kind]string{
		reflect.Invalid: "ident",
		reflect.Int:   "int",
		reflect.Float64: "float64",
		reflect.Bool:    "bool",
		reflect.String:  "string",
		reflect.Func: "func",
	}

	var op string

	if m, ok := yyXLAT[r.op]; ok {
		op = yySymNames[m]
	} else {
		op = "NA"
	}
	if v, ok = typeName[r.kind]; ok {
		fmt.Printf("%s%s: op[%s] %v %v\n", pre, v, op, r.Label, r.Value)
	} else if r.kind == reflect.Slice {
		fmt.Printf("%sSlice: op[%s] %+v\n", pre, yySymNames[yyXLAT[r.op]], r.node)
	}

	for i := 0; i < r.Len(); i++ {
		r.node[i].debug(lev + 1)
	}
}