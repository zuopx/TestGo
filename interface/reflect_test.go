package interface_test

import (
	"fmt"
	"reflect"
	"testing"
	"touchgo/utils"
)

func TestReflect(t *testing.T) {
	reflectDemo2()
	utils.Hello()
}

func reflectDemo1() {
	reflectFunc1(1)
}

// 获得对象的方法，转换参数并调用
func reflectDemo2() {
	s := &Schuduler{}
	v := reflect.ValueOf(s)
	
	t := reflect.TypeOf(s)
	dogIn := []interface{}{"Tom", 3}
	
	for i := 0; i < t.NumMethod(); i++ {
		name := t.Method(i).Name
		method := v.MethodByName(name)
		method.Type().In(0)
		method.Call([]reflect.Value{reflect.ValueOf(dogIn)})
		// fmt.Println(r)
	}
	fmt.Println(s, v, t)
}

func reflectFunc1(a interface{}) {
	b := reflect.ValueOf(a)
	t := reflect.TypeOf(a)
	k := t.Kind()

	fmt.Println(b, t, k)
}

type Schuduler struct{}

type DogIn struct {
	Name string
	Age int
}

func (s *Schuduler) Dog(args []interface{}) {
	fmt.Println(args[0], args[1])
}

// func (s *Schuduler) Cat() {
// 	fmt.Println("Cat.")
// }
