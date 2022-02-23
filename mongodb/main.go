package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	ping(client, ctx)

	document := bson.D{
		{"rollNo", 175},
		{"maths", 80},
		{"science", 90},
		{"computer", 95},
	}

	insertOneResult, err := insertOne(client, ctx, "gfg", "marks", document)

	if err != nil {
		panic(err)
	}

	fmt.Println("Result of InsertOne")
	fmt.Println(insertOneResult.InsertedID)

	documents := []interface{}{
		bson.D{
			{"rollNo", 153},
			{"maths", 65},
			{"science", 59},
			{"computer", 55},
		},
		bson.D{
			{"rollNo", 162},
			{"maths", 86},
			{"science", 80},
			{"computer", 69},
		},
	}

	insertManyResult, err := insertMany(client, ctx, "gfg", "marks", documents)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result of InsertMany")

	for id := range insertManyResult.InsertedIDs {
		fmt.Println(id)
	}
}
