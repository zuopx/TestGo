package datastructure_test

import (
	"fmt"
	"testing"
	"touchgo/utils"
)

func TestMap(t *testing.T) {
	mapDemo3()
	utils.Hello()
}

// nil map and non-nil map
func mapDemo1() {
	var m1 map[int]int
	fmt.Println(m1 == nil)

	m2 := map[int]int{}
	fmt.Println(m2 == nil)

	m3 := make(map[int]int)
	fmt.Println(m3 == nil)
}

// complex map
// 要分层初始化
func mapDemo2() {
	m1 := map[int]map[int]int{}
	m1[1] = map[int]int{}
	m1[1][1] = 1
	m1[1][2] = 2
	PrintMap(m1[1])
	fmt.Println(m1)
}

// 未初始化的slice和map都是nil
// nil-slice可以使用append方法
// nil-map不可以赋值，但可以取值，会返回默认初始化值
// append返回的是一个新的引用
// 所有map一定要初始化，slice可以通过append等方式达到初始化
func mapDemo3() {
	m1 := map[int]map[int][]int{}
	m1[1] = map[int][]int{}
	m1[1][2] = append(m1[1][2], 0)
	a1, ok:= m1[1]
	a1[1] = []int{}
	b1, ok:= a1[1]
	b1 = append(b1, 1)
	a1[1] = append(a1[1], 2)
	// m1[1][1] = []int{1, 2, 3}
	// fmt.Println(m1)

	// m2 := map[int]map[int]map[int]bool{}
	// m2[1][1][1] = true
	fmt.Println(m1, a1, b1, ok)
}

func PrintMap(m map[int]int) {
	fmt.Println("{")
	for k, v := range m {
		fmt.Printf("\t%d:%d\n", k, v)
	}
	fmt.Println("}")
}
