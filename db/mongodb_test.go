package db_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"touchgo/utils"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
	// InfoSlice []string
	// InfoMap    map[string]string
}

type TrainerTestSuite struct {
	suite.Suite
	coll *mongo.Collection
}

type ST = TrainerTestSuite

func (st *ST) SetupSuite() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	st.coll = client.Database("test").Collection("trainers")
}

func TestMongoDB(t *testing.T) {
	suite.Run(t, new(ST))
}

// 写入struct时，键转换成小写
// 传入结构、残缺的结构
// 传入字典
// 传入bson.M
func (st *ST) TestInsertOne() {
	// ash := &Trainer{"Ash", 10, "Pallet Town"}
	brock := bson.M{
		"Name": "Brock",
	}
	// st.coll.InsertOne(context.TODO(), ash)
	st.coll.InsertOne(context.TODO(), brock)
	utils.Hello()
}

func (st *ST) TestFindOne() {
	filter := bson.D{{"name", "Brock"}}
	var rtn Trainer
	st.coll.FindOne(context.TODO(), filter).Decode(&rtn)
	fmt.Println(rtn)
}

// cur.Decode会自动转换大小写
func (st *ST) TestFindMany() {
	cur, _ := st.coll.Find(context.TODO(), bson.D{})
	for cur.Next(context.TODO()) {
		var rtn Trainer
		err := cur.Decode(&rtn)
		if err != nil {
			continue
		}
		fmt.Println(rtn)
	}
}

func (st *ST) TestUpdate() {
	filter := bson.D{{"name", "Tom"}}
	u := bson.D{{"name", "Misty"}, {"Age", 10}, {"city", "Cerulean City"}}
	update := bson.D{
		{"$set", bson.D{u[0]}},
		{"$set", bson.D{u[1]}},
		{"$set", bson.D{u[2]}},
	}
	// SetUpsert(true)，如果没找到，先按照filter进行插入，再按照update进行更新
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	var rtn Trainer
	st.coll.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&rtn)
	fmt.Println(rtn)
}
