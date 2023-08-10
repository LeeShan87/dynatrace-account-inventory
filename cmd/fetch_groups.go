package main

import (
	"context"
	"fmt"
	"os"

	auth "github.com/LeeShan87/dynatrace-account-inventory/dynatrace"
	"github.com/LeeShan87/dynatrace-account-inventory/generated/account"
	"github.com/LeeShan87/dynatrace-account-inventory/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	apiURL = "https://api.dynatrace.com"
)

func main() {
	accountID := os.Getenv("ACOUNT_ID")
	mongoURI := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	utils.CheckError(err)
	defer client.Disconnect(context.Background())

	db := client.Database("account_api")

	authClient := auth.NewDTAuthClient()
	apiClient, err := auth.GetApiClient(*authClient, apiURL)
	utils.CheckError(err)

	// Fetch groups
	collection := db.Collection("groups")
	groups := fetchGroups(apiClient, accountID)
	insertManyResult, err := saveGroupsInMongo(groups, collection)
	utils.CheckError(err)
	fmt.Println("Inserted IDs:", len(insertManyResult.InsertedIDs))
}

func fetchGroups(client *account.APIClient, accountID string) *account.GroupListDto {
	groups, _, err := client.GroupManagementAPI.GroupsControllerGetGroups(context.Background(), accountID).Execute()
	utils.CheckError(err)
	return groups
}

func saveGroupsInMongo(groups *account.GroupListDto, collection *mongo.Collection) (*mongo.InsertManyResult, error) {
	bsonGroups := make([]interface{}, len(groups.Items))
	for i, group := range groups.Items {
		bsonGroup, _ := bson.Marshal(group)
		bsonGroups[i] = bsonGroup
	}
	insertManyResult, err := collection.InsertMany(context.Background(), bsonGroups)
	return insertManyResult, err
}
