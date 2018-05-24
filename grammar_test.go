package ruler

import (
	"fmt"
	"testing"
)

var rulerList [][]interface{}
var data map[string]interface{}

func TestFilter(t *testing.T) {
	fmt.Println(Filter(`abc[a == 10 || a == 101]`, data))
	fmt.Println(Filter(`abc[a in (10, 101)]`, data))
	fmt.Println(Filter(`abc[a notin (111)]`, data))
	fmt.Println(Filter(`abc[a != 111]`, data))
}

func TestHit(t *testing.T) {
	for _, v := range rulerList {
		if e, ok := Hit(v[0].(string), data); ok == v[2].(bool) {
			t.Log("Parse:", v, "OK")
		} else {
			t.Error("Parse:", v, e)
		}
	}
}

func init() {
	rulerList = [][]interface{}{
		//{`int("123") == 12`, true, true},
		//{`he()`, true, true},
		//{`he1(1, 3)`, true, true},
		//{`h1 == 1.2 || h2 == 100`, true, true},
		//{"cde.a.hello == 10", true, true},
		//{"a", true, true},
		{"in in (1, 2, 3, ha)", true, true},
		{"in notin (1, 2, 3)", true, true},
		// {"h1 > 1 && abc[a!=10]", true, true},
		//{"mm > 10 && a != 10 || af < 10", true, true},
		//{`abc(hello, "hello-\"world")`, true, true},
		//{`"abcde`, false, false},
		//{"1234345 ==", false, false},
	}

	data = map[string]interface{}{
		"abc": []map[string]interface{}{
			{"a": 10, "b": 20},
			{"a": 101, "b": 201},
			{"a": 111, "b": 221},
		},
		"cde": map[string]interface{}{
			"a": map[string]interface{}{
				"hello": "world",
				"world": "!",
				"kk":    12,
			},
		},
		"h1": 1.2,
		"h2": int(100),
		"h3": "hello world",
		"h4": "",
		"h5": 0,
		"in": 123,
		"ha": 123,
	}

	Bind("he", Hello)
	Bind("he1", Hello1)
}

func Hello() int {
	fmt.Println("Call builtin: Hello")
	return 1
}

func Hello1(a, b int64) int64 {
	fmt.Println("Call builtin: Hello1", a, b)
	return a + b
}
