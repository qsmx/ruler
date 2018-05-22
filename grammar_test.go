package ruler

import (
	"testing"
)

var rulerList []string
var data map[string]interface{}

func TestParse(t *testing.T) {
	for _, v := range rulerList {
		if Parse(v, data) {
			t.Log("Parse:", v, "OK")
		} else {
			t.Error("Parse:", v, "Error")
		}
	}
}

func TestValidate(t *testing.T) {
	for _, v := range rulerList {
		if Validate(v) {
			t.Log(v, "OK")
		} else {
			t.Error(v, "Error")
		}
	}
}

func init() {
	rulerList = []string {
		"abc == 10",
		"a",
		"abc[ff>10||kk==10]",
		"mm > 10 && a != 10 || af < 10",
        `abc(hello, "hello-\"world")`,
        `1"abc == 19`,
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