// 反射是指在程序运行期对程序本身进行访问和修改的能力。
// 对类进行反射，拿到属性和方法
// 拿到函数的输入、输出

package reflect_test

import (
	"fmt"
	"reflect"
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

// func (d *Dog) Eat(food string) int {
// 	d.Food = append(d.Food, food)
// 	return len(d.Food)
// }

type Arg []interface{}

func (d *Dog) Speak(arg []interface{}) {
	fmt.Println(arg)
}

func TestDog(t *testing.T) {
	suite.Run(t, new(St))
}
func (st *St) SetUpSuite() {
	st.d = &Dog{}
}

func (st *St) TestMethod() {
	vDog, tDog := reflect.ValueOf(st.d), reflect.TypeOf(st.d)
	tMethod := tDog.Method(0)
	fmt.Println(tMethod.Name)
	vMethod := vDog.MethodByName(tMethod.Name)
	in := []reflect.Value{reflect.ValueOf([]interface{}{1, 2, 3})}
	vMethod.Call(in)
}

func (st *St) TestField() {

}


func (st *St) TestFuncIn() {
	tDog := reflect.TypeOf(st.d)
	method := tDog.Method(0)
	fmt.Println(method.Name)
}

func (st *St) TestFuncOut() {

}

// func (t reflect.Type) Elem() reflect.Type
// 如果t不等于指针类型（包括Slice，Map），则报错。
func (st *St) TestElement() {
	pType := reflect.TypeOf(st.d)
	eType := pType.Elem()
	fmt.Println(eType.Name())
}

func (st *St) TestReflectValue() {
	rValue := reflect.ValueOf(st.d)
	rValue.Field(1)
}