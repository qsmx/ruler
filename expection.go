package ruler

import "fmt"

func throw_exception(vfmt string, args ...interface{}) {
	v := vfmt

	if len(args) != 0 {
		v = fmt.Sprintf(vfmt, args...)
	}

	panic(v)
}