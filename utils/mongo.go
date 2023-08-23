package utils

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindDocumentIDByUUID(ctx context.Context, collection *mongo.Collection, uuidKey, uuidValue string) (string, error) {
	filter := bson.M{uuidKey: uuidValue}
	var result map[string]interface{}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return "", err
	}
	documentID, ok := result["_id"]
	if !ok {
		return "", fmt.Errorf("document ID not found for UUID: %s", uuidValue)
	}
	return documentID.(string), nil
}

func ConnecToMongoDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	return client, err
}
