package datastructure_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
)

type StringTestSuite struct {
	suite.Suite
}

func TestString(t *testing.T) {
	suite.Run(t, new(StringTestSuite))
}

// TestFormat 测试fmt.Sprintf中不同类型数据转换成字符串。
// ref: https://www.runoob.com/go/go-fmt-sprintf.html
func (st *StringTestSuite) TestFormat() {
	b := true
	i := 10
	f := 3.14
	s := "abc"
	sc := bson.E{"name", "percy"}               // struct
	l := []interface{}{b, i, f, s, sc}                 // slice
	m := map[int]string{1: "A", 2: "B", 3: "C"} // map

	st.Equal("true", fmt.Sprintf("%v", b))
	st.Equal("10", fmt.Sprintf("%d", i))
	st.Equal("3.140000", fmt.Sprintf("%f", f))
	st.Equal("3.140", fmt.Sprintf("%.3f", f))
	fmt.Printf("%+v\n%v\n%+v\n", l, m, sc)
}
