package reflect_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
}

func TestPackage(t *testing.T) {
	suite.Run(t, new(PackageTestSuite))
}

// 动态获得变量名和值
// 例如 constant 包中定义了 const(StateNone = 0)，
// 通过 "constant"和"StateNone"拿到对应的值0
func (st *PackageTestSuite) TestGetVarible() {
}