package main

import (
	"fmt"
	"time"
	"context"
	// Install mongoDB go driver by "go get go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

func main(){

	// Create a context for mongoDB's timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	// Connect to MongoDB by mogo.Connect.
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017/"))
	if err != nil {
		panic(err)
	}

	// Locate the collection in the database
	collection := client.Database("test").Collection("tasks")
	
	// Use Find() to retrieve all data the collections as a cursor.
	cur, err := collection.Find(context.Background(), bson.M{} )
	if err != nil {
		panic(err)
	}

	// Loop through the data cursor and print out.
	for cur.Next(context.Background()) {

		// Print out the raw bson as string format
		fmt.Println("--MongoDB bson raw data printing out--")
		fmt.Println(cur.Current)

		// Decode the bson to struct and print out.
		fmt.Println("--Decoded to struct printing out--")
		result := struct {
			Name string
			Age int
		}{}

		err := cur.Decode(&result)
		if err != nil {
			fmt.Println(err)
		}
		
		fmt.Println("name: ", result.Name)
		fmt.Println("age: ", result.Age)
	}
}