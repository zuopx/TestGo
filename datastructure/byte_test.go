package datastructure_test

import (
	"fmt"
	"testing"
	"touchgo/utils"

	"github.com/stretchr/testify/suite"
)

type ByteTestSuite struct {
	suite.Suite
}

func TestByte(t *testing.T) {
	suite.Run(t, new(ByteTestSuite))
}

func (st *ByteTestSuite) TestHex() {
	a := 0x16
	fmt.Println(a & 0x0F)

	utils.Hello()
}
