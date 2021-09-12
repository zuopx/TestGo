package datastructure_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

)

func TestArray(t *testing.T) {
	// arrayDemo1()
	sliceDemo4()
}

// n := [10]int{} 完成声明和初始化，等价于 var n [10]int
func arrayDemo1() {
	var n [10]int
	for i := 0; i < 10; i++ {
		n[i] = i + 100
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("Element[%d] = %d\n", i, n[i])
	}
}

// 切片 即 为指定大小的数组
// n 的类型为 [10]int, m 的类型为 []int
// 空切片 等于 nil
func sliceDemo1() {
	n := [10]int{}
	m := n[:]
	for i := 0; i < 10; i++ {
		fmt.Printf("Element[%d] = %d\n", i, m[i])
	}
}

// 切片函数：len(), cap(), append(), copy()
func sliceDemo2() {
	m := make([]int, 3, 10)

	m = append(m, 1)
	PrintSlice(m)

	m = append(m, []int{2, 3}...)
	PrintSlice(m)
}

// 随机删除一个元素
func sliceDemo3() {
	m := []int{1, 2, 3, 4, 5}
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(m))
	e := m[idx]
	m = append(m[:idx], m[idx+1:]...)
	fmt.Println(e)
	PrintSlice(m)
}

func sliceDemo4 () {
	var m []int
	// m := []int{}
	for _, i := range m {
		fmt.Println(i)
	}
	fmt.Println(len(m))
	fmt.Println(m == nil)
}

func PrintSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
