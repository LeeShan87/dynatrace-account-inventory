package main

import (
	"context"
	"fmt"

	auth "github.com/LeeShan87/dynatrace-account-inventory/dynatrace"
	"github.com/LeeShan87/dynatrace-account-inventory/generated/account"
	"github.com/LeeShan87/dynatrace-account-inventory/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// fetch permissions for groups
	groups, err := fetchGroupsWithoutPermsInMongo(db)
	utils.CheckError(err)
	fmt.Println("Groups permissions will be updated: ", len(groups))
	for _, group := range groups {
		groupID := group.Uuid
		fmt.Println("Fetching permissions for group: ", *groupID)
		permissions, err := getPersmissionsForGroup(apiClient, utils.AccountID, *groupID)
		utils.CheckError(err)
		err = saveGroupPermissionsInMongo(db, group, *permissions)
		utils.CheckError(err)
	}
}

func getPersmissionsForGroup(client *account.APIClient, accountID, group string) (*account.PermissionsGroupDto, error) {
	permissions, _, err := client.PermissionManagementAPI.PermissionsControllerGetGroupPermissions(context.Background(), accountID, group).Execute()
	return permissions, err
}

func saveGroupPermissionsInMongo(db *mongo.Database, group account.GetGroupDto, permissions account.PermissionsGroupDto) error {
	permissionsCollection := db.Collection("permissions")
	groupCollection := db.Collection("groups")
	permissionsBson := make([]interface{}, len(permissions.Permissions))
	for i, permission := range permissions.Permissions {
		mongoPermission := auth.ToMongoPermission(permission)
		mongoPermission.GroupID = *group.Uuid
		bsonPermission, _ := bson.Marshal(mongoPermission)
		permissionsBson[i] = bsonPermission
	}
	insertManyResult, err := permissionsCollection.InsertMany(context.Background(), permissionsBson)
	if err != nil {
		return err
	}
	permissionIDs := make([]primitive.ObjectID, len(insertManyResult.InsertedIDs))
	for i, id := range insertManyResult.InsertedIDs {
		permissionIDs[i] = id.(primitive.ObjectID)
	}
	groupID, _ := utils.FindDocumentIDByUUID(context.Background(), groupCollection, "uuid", *group.Uuid)
	filter := bson.M{"_id": groupID}
	update := bson.M{"$set": bson.M{"permissionRefs": permissionIDs}}
	_, err = groupCollection.UpdateOne(context.Background(), filter, update)
	return err
}

func fetchGroupsWithoutPermsInMongo(db *mongo.Database) ([]account.GetGroupDto, error) {
	collection := db.Collection("groups")
	filter := bson.D{{"permissionRefs", bson.D{{"$exists", false}}}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var groups []account.GetGroupDto
	if err = cursor.All(context.Background(), &groups); err != nil {
		return nil, err
	}
	return groups, nil
}
