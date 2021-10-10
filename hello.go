package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var a int = 100
	if a < 20 {
		fmt.Println("a 小于20")
	} else {
		fmt.Println("a 不小于20")
	}

	var s string = "一"
	switch s {
	case "一":
		fmt.Println(1)
	case "二":
		fmt.Println(2)
	default:
		fmt.Println(3)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var f bool = true
	for f {
		fmt.Println(f)
		f = false
	}

	fmt.Println(add(1, 2))

}

func add(a int, b int) (int, int) {
	return a + b, a - b
}
