package ruler

import (
	"testing"
)

var dpData map[string]interface{}

func TestNewDataPackage(t *testing.T) {
	dp := NewDataPackage(dpData)

	v, _ := dp.GetAttr("b")
	if v != 100 {
		t.Error("dpData.b")
	}

	a, _ := dp.GetAttr("a")
	dp.Push(a)

	b, _ := dp.GetAttr("b")
	if b != "he" {
		t.Error("dpData.a.b", b)
	}
}

func TestDataPackage_GetDeepAttr(t *testing.T) {
	dp := NewDataPackage(dpData)

	if v, ok := dp.GetDeepAttr("a.c"); ok && v != 1.20 {
		t.Error("...")
	}
}

func init() {
	dpData = map[string]interface{} {
		"a": map[string]interface{} {
			"a": 1,
			"b": "he",
			"c": 1.20,
		},
		"b": 100,
	}
}