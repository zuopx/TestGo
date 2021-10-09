package datastructure_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type NumberTestSuite struct {
	suite.Suite
}

func TestNumber(t *testing.T) {
	suite.Run(t, new(NumberTestSuite))
}

// int, int32, int64, interface{} 互相转换
func (st *NumberTestSuite) TestInterface2Int() {
	a := []interface{}{int(1), int32(2), int64(3)}
	var err error

	// _, err = fmt.Println(a[0].(int))
	// st.Nil(err)

	// _, err = fmt.Println(a[0].(int32))
	// st.Nil(err)

	// _, err = fmt.Println(a[0].(int64))
	// st.Nil(err)

	// _, err = fmt.Println(a[1].(int))
	// st.Nil(err)

	// _, err = fmt.Println(a[1].(int32))
	// st.Nil(err)

	_, err = fmt.Println(a[1].(int64))
	st.Nil(err)

	// _, err = fmt.Println(a[2].(int))
	// st.Nil(err)

	// _, err = fmt.Println(a[2].(int32))
	// st.Nil(err)

	// _, err = fmt.Println(a[2].(int64))
	// st.Nil(err)
}

func (st *NumberTestSuite) TestInt() {
	i := 1
	i32 := 1
	iu32 := 1
	// i64 := int64(1)
	// iu64 := uint64(1)
	fmt.Println(int64(i), int64(i32), int64(iu32))
}
