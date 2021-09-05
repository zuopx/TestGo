package interface_test

import (
	"fmt"
	"testing"
)

/*
	1.	通过考虑数据类型之间的相同功能来创建抽象，而不是相同字段
	2.	interface{} 的值不是任意类型，而是 interface{} 类型
	3.	接口包含两个字的大小，类似于 (type, value)
	4.	函数可以接受 interface{} 作为参数，但最好不要返回 interface{}
	5.	指针类型可以调用其所指向的值的方法，反过来不可以
	6.	函数中的参数甚至接受者都是通过值传递
	7.	一个接口的值就是就是接口而已，跟指针没什么关系
	8.	如果你想在方法中修改指针所指向的值，使用 * 操作符
*/

func TestInterface(t *testing.T) {
	interfaceDemo2()
	fmt.Println("Hello,World!")
}

func interfaceDemo2() {
	func1(1)
	func1(1.0)
	func1("abc")
	func1([]int{1, 2, 3})
}

// 少用interface{}
func func1(arg interface{}) {
	fmt.Println(arg)
}

func interfaceDemo1() {
	animal1 := Dog{}
	fmt.Println(animal1.Speak())
	animal2 := Cat{}
	fmt.Println(animal2.Speak())
}

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "I'm Dog!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "I'm Cat!"
}
