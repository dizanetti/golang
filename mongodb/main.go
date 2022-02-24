package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	client, ctx, cancel, err := connect()
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)

	insert(client, ctx)
	update(client, ctx)
	find(client, ctx)
	delete(client, ctx)
}

func insert(client *mongo.Client, ctx context.Context) {
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

	log.Println("Result of InsertOne")
	log.Println(insertOneResult.InsertedID)

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

	log.Println("Result of InsertMany")

	for id := range insertManyResult.InsertedIDs {
		log.Println(id)
	}
}

func update(client *mongo.Client, ctx context.Context) {
	var filter interface{}

	filter = bson.D{
		{"maths", bson.D{{"$lt", 100}}},
	}

	update := bson.D{
		{"$set", bson.D{
			{"maths", 100},
		}},
	}

	resultUpdate, err := updateOne(client, ctx, "gfg", "marks", filter, update)
	if err != nil {
		panic(err)
	}

	log.Println("update single document")
	log.Println(resultUpdate.ModifiedCount)

	filter = bson.D{
		{"computer", bson.D{{"$lt", 100}}},
	}

	update = bson.D{
		{"$set", bson.D{
			{"computer", 100},
		}},
	}

	resultUpdate, err = updateMany(client, ctx, "gfg", "marks", filter, update)
	if err != nil {
		panic(err)
	}

	log.Println("update multiple document")
	log.Println(resultUpdate.ModifiedCount)
}

func find(client *mongo.Client, ctx context.Context) {
	var filter, option interface{}

	filter = bson.D{
		{"maths", bson.D{{"$gt", 70}}},
	}

	option = bson.D{{"_id", 0}}

	cursor, err := query(client, ctx, "gfg",
		"marks", filter, option)
	if err != nil {
		panic(err)
	}

	var results []bson.D

	if err := cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	log.Println("Query Reult")
	for _, doc := range results {
		log.Println(doc)
	}
}

func delete(client *mongo.Client, ctx context.Context) {

	query := bson.D{
		{"maths", bson.D{{"$gt", 60}}},
	}

	resultDelete, err := deleteOne(client, ctx, "gfg", "marks", query)
	if err != nil {
		panic(err)
	}

	log.Println("No.of rows affected by DeleteOne()")
	log.Println(resultDelete.DeletedCount)

	query = bson.D{
		{"science", bson.D{{"$gt", 0}}},
	}

	resultDelete, err = deleteMany(client, ctx, "gfg", "marks", query)
	if err != nil {
		panic(err)
	}

	log.Println("No.of rows affected by DeleteMany()")
	log.Println(resultDelete.DeletedCount)
}
