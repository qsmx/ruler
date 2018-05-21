package ruler

import (
	"testing"
	"fmt"
)

func TestParse(t *testing.T) {
	fmt.Printf("%+v\n", data)
	if !Parse(`abc [ a != 10 ]`, data) {
		t.Error("what")
	}
}