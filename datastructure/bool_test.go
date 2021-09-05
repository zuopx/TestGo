package datastructure_test

import (
	"fmt"
	"testing"
)


type Trainer struct {
	Name string
	Age  int
	City string
	InfoSlice []string
	InfoMap    map[string]string
}

func TestBool(t *testing.T) {
	Tom := &Trainer{}
	if Tom.Age == 0 {
		fmt.Println("Age")
	}
}