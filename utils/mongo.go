package utils

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindDocumentIDByUUID(ctx context.Context, collection *mongo.Collection, uuidKey, uuidValue string) (primitive.ObjectID, error) {
	filter := bson.M{uuidKey: uuidValue}
	var result map[string]interface{}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return primitive.NilObjectID, err
	}
	documentID, ok := result["_id"].(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("document ID not found for UUID: %s", uuidValue)
	}
	return documentID, nil
}
