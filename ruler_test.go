package ruler

import (
	"testing"
)

var data map[string]interface{}
var dp *DataPackage

func catch(t *testing.T) {
	if err := recover(); err != nil {
		t.Error("Recover:", err)
	}
}

//func TestNewDataPackage(t *testing.T) {
//	defer catch(t)
//
//	NewDataPackage(map[string]interface{}{"abc":1})
//	NewDataPackage(1)
//}
//
//func TestDataPackage_GetAttr(t *testing.T) {
//	defer catch(t)
//
//	if dp.GetAttr("abc") == 1 {
//		t.Log("Pass")
//	}
//}
//
//func TestDataPackage_GetBool(t *testing.T) {
//	if dp.GetBool("h4") {
//		t.Error("h4 = ''")
//	} else {
//		t.Log("yes")
//	}
//
//	if dp.GetBool("h5") {
//		t.Error("is 0")
//	} else {
//		t.Log("yes")
//	}
//}
//
//func TestDataPackage_GetInt(t *testing.T) {
//	defer catch(t)
//
//	if dp.GetInt("h1") == 1 {
//		t.Log("h1 == 1, OK")
//	} else {
//		t.Error("h1 == 1, Failed")
//	}
//
//	if dp.GetInt("h3") == 0 {
//		t.Error("Failed")
//	}
//
//}
//
//func TestDataPackage_GetString(t *testing.T) {
//	defer catch(t)
//
//	if dp.GetString("abc") == "1" {
//		t.Log("abc = '1'")
//	} else {
//		t.Error("Error")
//	}
//}

func init() {
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

	dp = NewDataPackage(data)
}