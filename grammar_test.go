package ruler

import (
	"testing"
	"fmt"
)

var rulerList [][]interface{}
var data map[string]interface{}

func TestParse(t *testing.T) {
	for _, v := range rulerList {
		if e, ok := Parse(v[0].(string), data); ok == v[2].(bool) {
			t.Log("Parse:", v, "OK")
		} else {
			t.Error("Parse:", v, e)
		}
	}
}

//func TestParseAll(t *testing.T) {
//	size := len(rulerList)
//	rulers := make([]string, size)
//	res := make([]bool, size)
//
//	for idx, v := range rulerList {
//		rulers[idx] = v[0].(string)
//		res[idx] = v[2].(bool)
//	}
//
//	ParseAll(rulers, data)
//}

//func TestValidate(t *testing.T) {
//	for _, v := range rulerList {
//		if e, ok := Validate(v[0].(string)); ok == v[1].(bool) {
//			t.Log(v, "OK")
//		} else {
//			t.Error(v, e)
//		}
//	}
//}

func init() {
	rulerList = [][]interface{} {
		{`int("123") == 123`, true, true},
		{`int("123") == 12`, true, true},
		{`he()`, true, true},
		{`he1(1, 3)`, true, true},
		//{`h1 == 1.2 || h2 == 100`, true, true},
		//{"cde.a.hello == 10", true, true},
		//{"a", true, true},
		//{"abc[|ff>10||kk==10]", true, true},
		//{"mm > 10 && a != 10 || af < 10", true, true},
        //{`abc(hello, "hello-\"world")`, true, true},
        //{`"abcde`, false, false},
        //{"1234345 ==", false, false},
	}

    data = map[string]interface{} {
        "abc": []map[string]interface{} {
            {"a": 10, "b": 20},
            {"a": 101, "b": 201},
            {"a": 111, "b": 221},
            {"a": map[string]interface{} {
				"hello": "world",
				"world": "!",
			}},
        },
		"cde": map[string]interface{} {
			"a": map[string]interface{} {
				"hello": "world",
				"world": "!",
			},
		},
        "h1": 1.2,
        "h2": int(100),
        "h3": "hello world",
        "h4": "",
        "h5": 0,
    }


	Bind("he", Hello)
	Bind("he1", Hello1)
}

func Hello() {
	fmt.Println("Call builtin: Hello")
}

func Hello1(a, b int64) {
	fmt.Println("Call builtin: Hello1", a, b)
}