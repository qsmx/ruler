package ruler

import (
	"testing"
	"fmt"
)

func TestConvertInt(t *testing.T) {
	a := [][]interface{} {
		{1.1, 1},
		{100.12, 100},
		{1, 1},
	}

	for _, v := range a {
		if ConvertInt(v[0]) == v[1].(int) {
			t.Log("yes")
		} else {
			fmt.Printf("%T %v, %T %v\n", v[0], v[0], v[1], v[1])
			t.Error("v[0]", v[0], "v[1]", v[1])
		}
	}

	if ConvertInt(a) != 3 {
		t.Error("?")
	}
}

func TestConvertFloat(t *testing.T) {

	a := [][]interface{} {
		{1.1, 1.1},
		{1, 1.0},
		{1000, 1000.0},
	}

	for _, v := range a {
		if ConvertFloat(v[0]) == v[1].(float64) {
			t.Log("yes")
		} else {
			fmt.Printf("%T %v, %T %v\n", v[0], v[0], v[1], v[1])
			t.Error("v[0]", v[0], "v[1]", v[1])
		}
	}
}