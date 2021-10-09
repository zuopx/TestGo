package datastructure_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StructTestSuite struct {
	suite.Suite
}

func TestStruct(t *testing.T) {
	suite.Run(t, new(StructTestSuite))
}

func (st *StructTestSuite) SetUpSuite() {
	
}

type StructA struct {
	Field1 int
	Field2 string
}

// 匿名字段
type StructB struct {
	StructA
	Field2 string
}

// 递归取字段
func (st *StructTestSuite) TestInherit() {
	b := StructB{
		StructA: StructA{
			Field1: 0,
			Field2: "Im in A",
		},
		Field2: "Im in B",
	}
	fmt.Println(b.Field1, b.Field2, b.StructA.Field2)
}
