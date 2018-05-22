package ruler

import "fmt"

const (
	ERR_STATUS_NO_ATTR = 400 + iota // 无属性
	ERR_STATUS_TYPE                 // 不支持的类型转换
	ERR_STATUS_ABORT 				// 解析token出错

	ERR_STATUS_DATA_STACK = 500 // 数据栈出现问题
)

type RulerException struct {
	code int
	str  string
}

func throwException(code int, vfmt string, args ...interface{}) {
	v := vfmt

	if len(args) != 0 {
		v = fmt.Sprintf(vfmt, args...)
	}

	panic(&RulerException{code: code, str: v})
}
