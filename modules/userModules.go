package modules

import (
	"context"
	"fmt"
	con "goApiPractice/configs"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User - data T
type User struct {
	Username string `json:"username"`
	Acc      string `json:"acc"`
	Pwd      string `json:"pwd"`
}

// ID - for mongo
type ID struct {
	ID primitive.ObjectID `bson:"_id"`
}

// Register - user register
func (user *User) Register() (id *mongo.InsertOneResult, err error) {
	// DB connect
	db := con.ConnectDB()
	// set use database and collection
	collection := db.Database("goBase").Collection("user")
	//設定等待時間兩秒
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// 對指定DB的table進行insert一筆資料
	id, fail := collection.InsertOne(ctx, user)
	fmt.Println("Register new user : " + user.Username)
	if fail != nil { //儲存錯誤的變數不為空值，代表有錯
		err = fail
		return
	}
	return
}
