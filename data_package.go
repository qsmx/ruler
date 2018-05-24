package ruler

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const data_stack_max_size = 20

type DataPackage struct {
	dataMap map[string]interface{}

	dataStack []interface{}
	pos       int8
}

func NewDataPackage(data interface{}) (dp *DataPackage) {
	dp = &DataPackage{
		dataStack: make([]interface{}, data_stack_max_size),
		pos:       0,
	}

	switch data.(type) {
	case string:
		dp.dataMap[data.(string)] = data

	// go里面转换json时除非明确指定类型为 map[int]interface{}
	// 何必做多余的事呢
	//case map[int]interface{}:
	//	   throwException(500, "不支持 map[int]interface{}")

	case map[string]interface{}:
		dp.dataMap = data.(map[string]interface{})

	default:
		throwException(500, "不支持的源数据类型 %T, %+v", data, data)
	}

	dp.dataStack[0] = &dp.dataMap
	return
}

func (dp *DataPackage) Pop() {
	if dp.pos == 0 {
		throwException(ERR_STATUS_DATA_STACK, "数据栈错误")
		os.Exit(-1)
	}

	dp.pos--
}

func (dp *DataPackage) Push(v interface{}) {
	switch v.(type) {
	case map[string]interface{}:
		data := v.(map[string]interface{})
		dp.pos++
		dp.dataStack[dp.pos] = &data
	default:
		throwException(500, "push 只支持 map[string]interface{}：%T, %+v", v, v)
	}
}

func (dp DataPackage) GetDeepAttr(name string) interface{} {
	var v interface{}

	names := strings.Split(name, ".")
	v = dp.dataStack[dp.pos]

	for _, key := range names {
		fmt.Printf("key: %s, from val = %+v\n", key, v)
		switch v.(type) {
		case map[string]interface{}:
			data := v.(map[string]interface{})
			v = valFromAttr(key, &data)
		case *map[string]interface{}:
			data := v.(*map[string]interface{})
			v = valFromAttr(key, data)
		default:
			throwException(500, "错误的类型：%T, %+v", v, v)
		}
	}

	return v
}

func (dp DataPackage) GetBaseAttr(name string) interface{} {
	return getAttr(name, &dp.dataMap)
}

func (dp DataPackage) GetAttr(name string) interface{} {
	v := dp.dataStack[dp.pos]

	switch v.(type) {
	case *map[string]interface{}:
		return getAttr(name, v.(*map[string]interface{}))
	default:
		throwException(500, "错误的类型：%T, %+v", v, v)
	}

	return nil
}

func valFromAttr(name string, data *map[string]interface{}) interface{} {
	v, ok := (*data)[name]
	if !ok {
		throwException(ERR_STATUS_NO_ATTR, "无属性值：%s", name)
	}

	return v
}

func ValueFormat(v interface{}) interface{} {
	switch v.(type) {
	case int8, int16, int32, int, uint8, uint16, uint32, uint64, uint:
		return ConvertInt(v)
	case float32:
		return ConvertFloat(v)
	default:
		return v
	}
}

func getAttr(name string, data *map[string]interface{}) interface{} {
	return ValueFormat(valFromAttr(name, data))
}

func (dp DataPackage) GetBool(name string) bool {
	v := dp.GetAttr(name)

	return ConvertBool(v)
}

func (dp DataPackage) GetInt(name string) int64 {
	v := dp.GetAttr(name)

	return ConvertInt(v)
}

func (dp DataPackage) GetFloat(name string) float64 {
	v := dp.GetAttr(name)

	return ConvertFloat(v)
}

func (dp DataPackage) GetString(name string) string {
	v := dp.GetAttr(name)

	return ConvertString(v)
}

func ConvertBool(v interface{}) bool {
	switch v.(type) {
	case nil:
		return false
	case bool:
		return v.(bool)
	case string:
		return v.(string) != ""
	case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint, float32, float64:
		return ConvertFloat(v) != 0.0
	default:
		w := reflect.ValueOf(v)
		switch w.Kind() {
		case reflect.Array, reflect.Slice, reflect.Map:
			return w.Len() > 0
		}
	}

	throwException(ERR_STATUS_TYPE, "不支持 %v 转换为 bool", v)
	return false
}

func ConvertInt(v interface{}) int64 {
	switch v.(type) {
	case int64:
		return v.(int64)
	case float64:
		return int64(v.(float64))
	case int8, int16, int32, int, uint8, uint16, uint32, uint64, uint, float32:
		m := reflect.ValueOf(v).Convert(reflect.TypeOf(int64(1)))
		return m.Int()
	case string:
		s := v.(string)
		if m, e := strconv.ParseInt(s, 10, 64); e == nil {
			return m
		} else if n, e := strconv.ParseFloat(s, 64); e == nil {
			return int64(n)
		}
	}

	throwException(ERR_STATUS_TYPE, "不支持 %v 转换为 int64", v)
	return 0
}

func ConvertFloat(v interface{}) float64 {
	switch v.(type) {
	case float64:
		return v.(float64)
	case int64:
		return float64(v.(int64))
	case int8, int16, int32, int, uint8, uint16, uint32, uint64, uint, float32:
		m := reflect.ValueOf(v).Convert(reflect.TypeOf(float64(1.0)))
		return m.Float()
	case string:
		s := v.(string)
		if n, e := strconv.ParseFloat(s, 64); e == nil {
			return float64(n)
		}
	}

	throwException(ERR_STATUS_TYPE, "不支持 %v 转换为 float64", v)
	return 0
}

func ConvertString(v interface{}) string {
	switch v.(type) {
	case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint, float32, float64:
		return fmt.Sprintf("%v", v)
	case string:
		return v.(string)
	}

	throwException(ERR_STATUS_TYPE, "不支持 %v 转换为 string", v)
	return ""
}
