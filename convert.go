package ruler

import (
	"reflect"
	"strconv"
	"fmt"
)

func ConvertBool(v interface{}) (val bool) {
	switch v.(type) {
	case bool:
		val = v.(bool)

	case string:
		val = len(v.(string)) != 0

	case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint, float32, float64:
		val = v != 0

	default:
		k := reflect.ValueOf(v)
		switch k.Kind() {
		case reflect.Array, reflect.Map, reflect.Slice:
			val = k.Len() > 0
		}
	}

	return val
}

func ConvertInt(v interface{}) (val int) {
	switch v.(type) {
	case int:
		val = v.(int)

	case int8, int16, int32, int64, uint8, uint16, uint32, uint64, uint, float32, float64:
		r := reflect.ValueOf(v)
		val = int(r.Convert(reflect.TypeOf(int(1))).Int())

	case string:
		str := v.(string)
		if m, err := strconv.ParseInt(str, 0, 64); err == nil {
			val = int(m)
		} else if n, err := strconv.ParseFloat(str, 64); err == nil {
			val = int(n)
		}

	default:
		k := reflect.ValueOf(v)
		switch k.Kind() {
		case reflect.Array, reflect.Map, reflect.Slice:
			val = k.Len()
		}
	}

	return
}

func ConvertFloat(v interface{}) (val float64) {
	switch v.(type) {
	case float64:
		val = v.(float64)

	case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint, float32:
		r := reflect.ValueOf(v)
		val = float64(r.Convert(reflect.TypeOf(float64(1))).Float())

	case string:
		str := v.(string)
		if n, err := strconv.ParseFloat(str, 64); err == nil {
			val = n
		}
	}

	return
}

func ConvertString(v interface{}) (val string) {
	switch v.(type) {
	case string:
		val = v.(string)

	case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint, float32, float64:
		val = fmt.Sprintf("%v", v)

	default:
		throwException(E_DATA_INVALID, "%T 不允许转换为字符中", v)
	}

	return
}

func ValueFormat(v interface{}) interface{} {
	switch v.(type) {
	case bool:
		return v.(bool)

	case string:
		return v.(string)

	case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint:
		return ConvertInt(v)

	case float32, float64:
		return ConvertFloat(v)
	}

	return v
}