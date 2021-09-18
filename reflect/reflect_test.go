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

type NewInt int

type NewDogIn struct {
	Name   string `validate:"required"`
	Age    NewInt
	Food   []string
	Parent map[string]string
}

func (d *Dog) NewDog(in NewDogIn) {
}

func (d *Dog) Eat(food string) int {
	d.Food = append(d.Food, food)
	return len(d.Food)
}

func (d *Dog) SetParent(parent map[string]string) {
	d.Parent = parent
}

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
	v, t := reflect.ValueOf(st.d), reflect.TypeOf(st.d)
	for i := 0; i < t.NumMethod(); i++ {
		name := t.Method(i).Name
		if name != "NewDog" {
			continue
		}
		method := v.MethodByName(name)
		numIn := method.Type().NumIn()
		fmt.Println(name)
		for j := 0; j < numIn; j++ {
			t1 := method.Type().In(j)
			fmt.Println(t1)
			for j := 0; j < t1.NumField(); j++ {
				field := t1.Field(j)
				k := field.Type.Kind()
				sk := k.String()
				if k == reflect.Slice {
					sk = fmt.Sprintf("[%s]", field.Type.Elem().String())
				} else if k == reflect.Map {
					sk = fmt.Sprintf("{%s:%s}", field.Type.Key().Kind().String(), field.Type.Elem().Kind().String())
				}

				require := field.Tag.Get("validate")
				if require == "required" {
					sk += ", required"
				}

				// sk := ""
				// switch k {
				// case reflect.Int, reflect.Int64:
				// 	sk = "int"
				// case reflect.String:
				// 	sk = "str"
				// case reflect.Slice:
				// 	if field.Type.Elem().Kind() == reflect.Int || field.Type.Elem().Kind() == reflect.Int64 {
				// 		sk = "[int,]"
				// 	} else {
				// 		sk = "[str,]"
				// 	}
				// case reflect.Map:
				// 	kk, vk := "", ""
				// 	if field.Type.Key().Kind() == reflect.Int || field.Type.Key().Kind() == reflect.Int64 {
				// 		kk = "int"
				// 	} else {
				// 		kk = "str"
				// 	}
				// 	if field.Type.Elem().Kind() == reflect.Int || field.Type.Elem().Kind() == reflect.Int64 {
				// 		vk = "int"
				// 	} else {
				// 		vk = "str"
				// 	}
				// 	sk = fmt.Sprintf("{%s:%s,}", kk, vk)
				// }
				fmt.Println("\t", field.Name, sk)
			}
		}
		fmt.Println()
	}
}

// slice和element类型
func (st *St) TestSliceElem() {
	intSlice := []int{}
	strSlice := []string{}

	t1, t2 := reflect.TypeOf(intSlice), reflect.TypeOf(strSlice)
	fmt.Println(t1, t2)
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
