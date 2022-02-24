package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func updateOne(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {
	collection := client.Database(dataBase).Collection(col)

	result, err = collection.UpdateOne(ctx, filter, update)

	return result, err
}

func updateMany(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {
	collection := client.Database(dataBase).Collection(col)

	result, err = collection.UpdateMany(ctx, filter, update)
	return result, err
}
