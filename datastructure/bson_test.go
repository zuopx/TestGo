package datastructure_test

import (
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBson(t *testing.T) {
	bsonDemo1()
}

// type bson.D []primitive.E
// type primitive.E struct {Key string, Value interface{}}
// type bson.M map[string]interface{}
// 本质上bson.D是有序字典，用slice和struct的方法构造有序字典
// 把bson.D当作字典，bson.E当作键值对
func bsonDemo1() {
	d := bson.D{{Key: "foo", Value: "bar"}, {"hello", "world"}}
	e := primitive.E{"hello", "world"}
	fmt.Println(d)
	fmt.Println(e)
}

func bsonDemo2() {

}

