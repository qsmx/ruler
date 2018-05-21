package ruler

import (
	"reflect"
	"strconv"
	"fmt"
)

const data_stack_max_size = 20

type DataPackage struct {
	dataMap		map[string]interface{}

	dataStack	[]interface{}
	pos 		int8
}

func NewDataPackage(data interface{}) *DataPackage {
	if v, ok := data.(map[string]interface{}); ok {
		dp := &DataPackage{
			dataMap: v,
			dataStack: make([]interface{}, data_stack_max_size),
			pos: 0,
		}

        dp.dataStack[0] = &dp.dataMap
        return dp
	}

	return nil
}

func (dp *DataPackage) Pre() {
	if dp.pos == 0 {
		throwException(ERR_STATUS_DATA_STACK, "数据栈错误")
	}

	dp.pos--
}

func (dp *DataPackage) Next(name string) {
	d := dp.GetAttr(name)

	fmt.Printf("%T, %+v\n", d, d)

	v, _ := d.([]map[string]interface{})
	dp.pos++
	dp.dataStack[dp.pos] = &(v[0])

	//switch reflect.ValueOf(d).Kind() {
	//case reflect.Slice, reflect.Array:
	//	dp.pos++
	//	dp.dataStack[dp.pos] = &d
	//}
}

func (dp DataPackage) GetAttr(name string) interface{} {
	v := dp.dataStack[dp.pos]

	switch v.(type) {
	case *map[string]interface{}:
		return getAttr(name, v.(*map[string]interface{}))
	}

	return nil
}

func getAttr(name string, data *map[string]interface{}) interface{} {
	if v, ok := (*data)[name]; ok {
		return v
	}

	throwException(ERR_STATUS_NO_ATTR, "无属性值：%s", name)
	return nil
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
	}

	w := reflect.ValueOf(v)
	switch w.Kind() {
	case reflect.Array, reflect.Slice:
		return w.Len() > 0
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

func ConvertType(a, b interface{}) (m, n interface{}) {
	m, ais := a.(float64)
	n, bis := b.(float64)
	if ais || bis {
		if !ais {
			m = ConvertFloat(a)
		}
		if !bis {
			n = ConvertFloat(b)
		}
	}

	return
}
