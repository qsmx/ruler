package ruler

import (
	"testing"
	"fmt"
)

func TestHit(t *testing.T) {
	data := map[string]interface{} {
		"a": 10,
		"b": 12,
		"c": []map[string]interface{}{
			{"aa": 12, "bb": 23, "cc": "hello"},
			{"aa": 23, "bb": 56, "cc": "hello"},
			{"aa": 12, "bb": 44, "cc": "hello"},
		},
	}

	rulers := [][]interface{} {
		{`a + 1 == 1 && b - 10 == 2`, true, false},
		{`(0x10 & 0x110 == 0x10`, false, true},
		{`a in (1, 111, 12)`, true, false},
		{`a notin (1, 111, 12)`, true, true},
		{`a in (1, 10, 111, 12)`, true, true},
		{`a notin (1, 10, 111, 12)`, true, false},
	}

	for _, ruler := range rulers {
		if res, ok := Hit(ruler[0].(string), data); ok == ruler[1] {
			if ok == false {
				fmt.Println("Err: ", ruler[0], "\nMessage:", res)
				continue
			}
			if res == ruler[2] {
				t.Log(ruler[0], "Ok")
			} else {
				t.Error(ruler[0], "Failed")
			}
		} else {
			t.Error(ruler[0], "err")
		}
	}

	//Hit(`[aa < 20 && bb >=30]`, data["c"])
	//Hit(`c[aa < 20 && aa == |b]`, data)
	//Hit(`c[aa < 20 && bb >= 30] && 1 == 1`, data)
	//Hit(`c[aa < 20 && bb >= 30] || 1 == 1`, data)
}