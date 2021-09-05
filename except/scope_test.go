package except_test

import (
	"fmt"
	"testing"
	"touchgo/utils"
)

func TestScope(t *testing.T) {
	scopeDemo1()
	utils.Hello()
}


func scopeDemo1() {
	var a int
	b := 1
	if b > 10 {
		fmt.Println(a)
	}else {
		fmt.Println(b)
	}
}
