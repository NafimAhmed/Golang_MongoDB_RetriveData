package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"golang.org/x/vuln/client"
)

const URL = "mongodb://localhost:27017"

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI(URL))

	if err != nil {

		log.Fatal(err)

	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	quickStartDatabase := client.Database("First_Database")

	quickStartCollection := quickStartDatabase.Collection("First_Collection")

	cursor, err := quickStartCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var episodes []bson.M

	err = cursor.All(ctx, &episodes)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(episodes)

	for _, episode := range episodes {

		fmt.Println(episode["title"])
	}

}
