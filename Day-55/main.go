package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name  string
	Email string
}

type StartLog struct {
	Hostname       string
	Pid            int
	StartTimeLocal string
}

func main() {
	// 設定身份驗證資訊
	auth := options.Credential{
		Username: "root",
		Password: "root",
	}

	// 建立與MongoDB的連線
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(auth))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	// 選擇要使用的資料庫和集合
	db := client.Database("local")
	collection := db.Collection("startup_log")

	// 查詢單筆資料
	var logResult StartLog
	err = collection.FindOne(ctx, bson.M{"pid": 28}).Decode(&logResult)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("查詢的Start Log: %+v\n", logResult)

	fmt.Printf("處理People\n")

	collection = db.Collection("people")

	// 建立一個新的Person物件
	person := Person{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	// 插入資料
	insertResult, err := collection.InsertOne(ctx, person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("插入的資料ID:", insertResult.InsertedID)

	// 查詢資料
	var result Person
	err = collection.FindOne(ctx, bson.M{"name": "John Doe"}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("查詢的資料: %+v\n", result)

	// 更新資料
	update := bson.M{"$set": bson.M{"email": "newemail@example.com"}}
	updateResult, err := collection.UpdateOne(ctx, bson.M{"name": "John Doe"}, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("更新的文件數量: %d\n", updateResult.ModifiedCount)

	// 刪除資料
	deleteResult, err := collection.DeleteOne(ctx, bson.M{"name": "John Doe"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("刪除的文件數量: %d\n", deleteResult.DeletedCount)
}
