package datastructure_test

import (
	"fmt"
	"testing"
)

func TestRanges(t *testing.T) {
	rangeDemo2()
	fmt.Println("Hello, World!")	
}

func rangeDemo1() {
	n := []int{1, 2, 3,}

	for i := range n {
		fmt.Println(i)
	}

	for idx, i := range n {
		fmt.Println(idx, i)
	}
}

func rangeDemo2() {
	m := map[int]int{1:1, 2:2, 3:3}
	
	for k := range m {
		fmt.Println(k)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
