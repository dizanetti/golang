package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func insertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	collection := client.Database(dataBase).Collection(col)
	result, err := collection.InsertOne(ctx, doc)

	return result, err
}

func insertMany(client *mongo.Client, ctx context.Context, dataBase, col string, docs []interface{}) (*mongo.InsertManyResult, error) {
	collection := client.Database(dataBase).Collection(col)
	result, err := collection.InsertMany(ctx, docs)

	return result, err
}
