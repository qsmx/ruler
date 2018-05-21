package ruler

import (
	"testing"
	"github.com/astaxie/beego/logs"
	"fmt"
)

func TestParse(t *testing.T) {
	fmt.Printf("%+v\n", data)
	if !Parse(`abc [ a != 10 ]`, data) {
		logs.Error("what")
	}
}