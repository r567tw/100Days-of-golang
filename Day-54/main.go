package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	if _, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017").SetConnectTimeout(5*time.Second)); err != nil {
		fmt.Print(err)
		return
	} else {
		fmt.Println("MongoDB connect success")
	}
}

// Ref: https://zhuanlan.zhihu.com/p/144308830
