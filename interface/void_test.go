// interface{}
package interface_test

import (
	"fmt"
	"reflect"
	"testing"
	"touchgo/utils"
)

func TestVoid(t *testing.T) {
	voidDemo1()
	utils.Hello()
}

// 测试空接口
// 用空接口作为类型，对象里是包含了原始类型信息的，可以用reflect.TypeOf得到w
func voidDemo1() {
	var a interface{} 
	a = []int{1, 2}
	r := voidFunc1(a)
	t := reflect.TypeOf(r)
	k := t.Kind()
	fmt.Println(r)
	fmt.Println(t)
	fmt.Println(k)
}

func voidFunc1(a interface{}) interface{} {
	fmt.Println(reflect.TypeOf(a))
	switch a.(type) {
	case int:
		return a.(int) + 10
	case float64:
		return a.(float64) + 20
	}
	return a
}

// 测试空接口数组
// []interface{} 是一种特定的类型，不等同于[]int之流
func voidDemo2() {
	var a []interface{}
	a = append(a, 1)
	a = append(a, 1.0)
	fmt.Println(a)
	b := printSlice(a)
	a[1] = 2.0
	fmt.Println(a)
	fmt.Println(b)
}

func printSlice(a []interface{}) []interface{} {
	for i, v := range a {
		fmt.Println(i, v, reflect.TypeOf(v))
	}
	return a
}

func voidDemo3() {
	m := voidFunc3().(map[int]int)
	fmt.Println(m)
}

func voidFunc3() interface{} {
	return map[int]int{1: 1, 2: 2}
}
