// 反射是指在程序运行期对程序本身进行访问和修改的能力。 
// 对类进行反射，拿到属性和方法
package reflect_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type Dog struct {
	Name   string
	Age    int
	Kind   int
	Food   []string
	Parent map[string]string
}

type DogTestSuite struct {
	suite.Suite
	d *Dog
}

type St = DogTestSuite

func (d *Dog) String() string {
	return fmt.Sprintf("Name:%s, Age:%d, Kind:%d", d.Name, d.Age, d.Kind)
}

func TestDog(t *testing.T) {
	suite.Run(t, new(St))
}
func (st *St) SetUpSuite() {
	st.d = &Dog{}
}

func (st *St) TestMethod() {

}

func (st *St) TestField() {

}
