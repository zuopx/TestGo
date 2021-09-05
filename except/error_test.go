package except_test

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
)

func TestError(t *testing.T) {
	// errorDemo1()
	fmt.Println("Hello, World!")
}

func errorDemo1() {
	// err := getError()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println("Valid")
	// }

	if err := getError(); err != nil {
		fmt.Println(err.Error())
		err := getError()
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Valid")
	}

	err := getError()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Valid")
	}
}

func getError() error {
	if rand.Intn(10) < 5 {
		return errors.New("Invalid")
	} else {
		return nil
	}
}
