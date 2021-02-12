package main

import (
	"fmt"
	"time"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

func main(){
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017/"))
	if err != nil {
		panic(err)
	}

	collection := client.Database("test").Collection("tasks")
	
	cur, err := collection.Find(context.Background(), bson.M{} )
	if err != nil {
		panic(err)
	}

	for cur.Next(context.Background()) {

		result := struct {
			Name string
			Age int
		}{}

		err := cur.Decode(&result)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("--Decoded to struct printing out--")
		fmt.Println("name: ", result.Name)
		fmt.Println("age: ", result.Age)

		fmt.Println("--MongoDB bson raw data printing out--")
		fmt.Println(cur.Current)
	}
}