package interface_test

import (
	// "reflect"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"touchgo/utils"
)

func TestGeneric(t *testing.T) {
	if r, err := genericAdd(1, 2); err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(r.(int))
	}

	utils.Hello()
}

func genericAdd(a interface{}, b interface{}) (interface{}, error) {
	ta, tb := reflect.TypeOf(a).Kind(), reflect.TypeOf(b).Kind()
	if ta != tb {
		return nil, errors.New("Inconsistent types.")
	}
	switch ta {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return a.(int) + b.(int), nil
	case reflect.Float32, reflect.Float64:
		return a.(float64) + b.(float64), nil
	default:
		return nil, errors.New("Unsupported types.")
	}
}

func genericMulti() {

}
