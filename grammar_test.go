package ruler

import (
	"testing"
	"fmt"
)

var rulerList [][]interface{}
var data map[string]interface{}

// func TestParse(t *testing.T) {
// 	for _, v := range rulerList {
// 		if Parse(v, data) {
// 			t.Log("Parse:", v, "OK")
// 		} else {
// 			t.Error("Parse:", v, "Error")
// 		}
// 	}
// }

func TestValidate(t *testing.T) {
	for _, v := range rulerList {
		if e, ok := Validate(v[0].(string)); ok == v[1].(bool) {
			t.Log(v, "OK")
		} else {
			fmt.Println(e, &e)
			t.Error(v, e)
		}
	}
}

func init() {
	rulerList = [][]interface{} {
		{"abc == 10", true},
		{"a", true},
		{"abc[ff>10||kk==10]", true},
		{"mm > 10 && a != 10 || af < 10", true},
        {`abc(hello, "hello-\"world")`, true},
        {`"abcde`, false},
        {"1234345 ==", true},
	}

    data = map[string]interface{} {
        "abc": []map[string]interface{} {
            {"a": 10, "b": 20},
            {"a": 101, "b": 201},
            {"a": 111, "b": 221},
        },
        "h1": 1.2,
        "h2": int(100),
        "h3": "hello world",
        "h4": "",
        "h5": 0,
    }
}