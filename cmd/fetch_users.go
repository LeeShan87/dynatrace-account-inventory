package main

import (
	"context"
	"fmt"

	auth "github.com/LeeShan87/dynatrace-account-inventory/dynatrace"
	"github.com/LeeShan87/dynatrace-account-inventory/generated/account"
	"github.com/LeeShan87/dynatrace-account-inventory/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	client, err := utils.ConnecToMongoDB()
	utils.CheckError(err)
	defer client.Disconnect(context.Background())

	db := client.Database("account_api")

	authClient := auth.NewDTAuthClient()
	apiClient, err := auth.GetApiClient(*authClient, utils.ApiURL)
	utils.CheckError(err)

	// Fetch users
	collection := db.Collection("users")
	users := fetchUsers(apiClient, utils.AccountID)
	insertManyResult, err := saveUsersInMongo(users, collection)
	utils.CheckError(err)
	fmt.Println("Inserted IDs:", len(insertManyResult.InsertedIDs))
}

func fetchUsers(client *account.APIClient, accountID string) *account.UserListDto {
	users, _, err := client.UserManagementAPI.UsersControllerGetUsers(context.Background(), accountID).Execute()
	utils.CheckError(err)
	return users
}

func saveUsersInMongo(users *account.UserListDto, collection *mongo.Collection) (*mongo.InsertManyResult, error) {
	bsonUsers := make([]interface{}, len(users.Items))
	for i, user := range users.Items {
		bsonUser, _ := bson.Marshal(auth.ToMongoUser(user))
		bsonUsers[i] = bsonUser
	}
	insertManyResult, err := collection.InsertMany(context.Background(), bsonUsers)
	return insertManyResult, err
}
