package ruler

import "testing"

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

	Hit("a + 1 == 1 && b - (1 +-2) != 10", data)
	Hit("(0x10 & 0x110) == 0x10", data)

	//Hit(`a in (1, 111, 12)`, data)
	//Hit(`a notin (1, 111, 12)`, data)
	//Hit(`a in (1, 10, 111, 12)`, data)
	//Hit(`a notin (1, 10, 111, 12)`, data)

	//Hit(`[aa < 20 && bb >=30]`, data["c"])
	Hit(`c[aa < 20 && aa == |b]`, data)
	//Hit(`c[aa < 20 && bb >= 30] && 1 == 1`, data)
	//Hit(`c[aa < 20 && bb >= 30] || 1 == 1`, data)
}