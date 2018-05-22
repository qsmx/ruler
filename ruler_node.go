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

func newRulerNode() *rulerNode {
	r := &rulerNode{
		typ: reflect.Invalid,
	}

	return r
}

func (r *rulerNode) AppendNode(n *rulerNode) *rulerNode {
    r.node = append(r.node, n)
    return r
}

func nodeFromNode(op int, n ...*rulerNode) *rulerNode {
	r := &rulerNode{
		op:  op,
		typ: reflect.Slice,
	}
	r.node = append(r.node, n...)
	return r
}

func nodeFromIdent(op int, ident string) *rulerNode {
	return &rulerNode{
		op:    op,
		value: ident,
	}
}

func nodeFromInt(op int, value string) *rulerNode {
	return &rulerNode{
		op:    op,
		value: ConvertInt(value),
		typ:   reflect.Int64,
	}
}

func nodeFromFloat(op int, value string) *rulerNode {
	return &rulerNode{
		op:    op,
		value: ConvertFloat(value),
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

func nodeFromBool(op int, value bool) *rulerNode {
	return &rulerNode{
		op:    op,
		value: value,
		typ:   reflect.Bool,
	}
}

func nodeFromInterface(op int, value interface{}) *rulerNode {
	return &rulerNode{
		op:    op,
		value: value,
		typ:   reflect.ValueOf(value).Kind(),
	}
}

func (r rulerNode) Len() int {
	return len(r.node)
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
