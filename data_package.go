package ruler

import (
	"reflect"
	"os"
	"fmt"
	"strings"
)

const DEFAULT_SLICE_LABEL = "content"
const DEFAULT_DATA_DEEP = 20

type DataPackage struct {
	data map[string]interface{}

	stack []map[string]interface{}
	pos  int
}

func NewDataPackage(data interface{}) (dp *DataPackage) {
	v := reflect.ValueOf(data)
	kind := v.Kind()

	if v.Len() == 0 {
		throwException(E_DATA_INVALID, "空数据，无需处理")
	}

	dp = &DataPackage{
		stack: make([]map[string]interface{}, DEFAULT_DATA_DEEP),
		pos: 0,
	}

	switch kind {
	case reflect.Map:
		if v.MapKeys()[0].Kind() != reflect.String {
			throwException(E_DATA_INVALID, "map 的 key 必须是 string")
		}

		if val, ok := data.(map[string]interface{}); ok {
			dp.data = val
		} else {
			for _, key := range v.MapKeys() {
				dp.data[key.String()] = v.MapIndex(key).Interface()
			}
		}

	case reflect.Slice, reflect.Array:
		dp.data = map[string]interface{}{
			DEFAULT_SLICE_LABEL: data,
		}

	default:
		throwException(E_DATA_INVALID, "不支持的数据类型 %T", data)
	}

	fmt.Printf("%+v\n", dp.data)
	dp.stack[dp.pos] = dp.data
	return
}

func (dp DataPackage) GetDeepAttr(name string) (v interface{}, ok bool) {
	if name[0] == '|' {
		v = dp.data
		name = name[1:]
	} else {
		v = dp.stack[dp.pos]
	}
	names := strings.Split(name, ".")
	for _, key := range names {
		//fmt.Printf("key: %s, from val = %+v\n", key, v)
		switch v.(type) {
		case map[string]interface{}:
			data := v.(map[string]interface{})
			v, ok = getAttr(key, data)
		default:
			return nil, false
		}
	}

	return
}


func (dp DataPackage) GetAttr(key string) (interface{}, bool) {
	if key[0] == '|' {
		return getAttr(key[1:], dp.data)
	} else {
		return getAttr(key, dp.stack[dp.pos])
	}
}

func getAttr(key string, data map[string]interface{}) (v interface{}, ok bool) {
	fmt.Println("key:", key, "data:", data)
	v, ok = data[key]
	return
}

func (dp *DataPackage) Push(data interface{}) (ok bool) {
	var v map[string]interface{}

	if v, ok = data.(map[string]interface{}); ok {
		if dp.pos + 1 == DEFAULT_DATA_DEEP {
			throwException(E_RULER_DATA_DEEP, "数据深度溢出")
		}

		dp.pos++
		dp.stack[dp.pos] = v
	} else {
		throwException(E_RULER_DATA_NOT_MATCH, "规则要求map, 但数据为 %T", data)
	}

	return
}

func (dp *DataPackage) Pop() {
	if dp.pos == 0 {
		fmt.Println("数据栈异常")
		os.Exit(255)
	}

	dp.pos--
}
