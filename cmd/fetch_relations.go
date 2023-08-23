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

	users, err := fetchUsersInMongo(db)
	utils.CheckError(err)
	fmt.Println("Users relationships will be updated: ", len(users))
	for _, user := range users {
		// Fetch user groups
		id := user.Uid
		email := user.Email
		fmt.Println("Fetching groups for user: ", id)
		groups := getUserGroups(apiClient, utils.AccountID, email)
		err = updateUserGroupsInMongo(db, id, groups)
		utils.CheckError(err)
	}
	groups, err := fetchGroupsInMongo(db)
	utils.CheckError(err)
	fmt.Println("Groups relationships will be updated: ", len(groups))
	for _, group := range groups {
		// Fetch users in group
		groupId := group.Uuid
		fmt.Println("Fetching users for group: ", groupId)
		users := getUsersInGroup(apiClient, utils.AccountID, *groupId)
		err = updateGroupsUsersInMongo(db, *groupId, *users)
		utils.CheckError(err)
	}
}

func fetchUsersInMongo(db *mongo.Database) ([]*account.UserDto, error) {
	collection := db.Collection("users")
	filter := bson.D{{"groupRefs", bson.D{{"$exists", false}}}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var users []*account.UserDto
	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

func fetchGroupsInMongo(db *mongo.Database) ([]account.GetGroupDto, error) {
	collection := db.Collection("groups")
	filter := bson.D{{"userRefs", bson.D{{"$exists", false}}}}
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

func getUserGroups(client *account.APIClient, accountID, email string) []account.AccountGroupDto {
	groups, _, err := client.UserManagementAPI.UsersControllerGetUserGroups(context.Background(), accountID, email).Execute()
	utils.CheckError(err)
	return groups.Groups
}

func getUsersInGroup(client *account.APIClient, accountID, groupId string) *account.GroupUserListDto {
	users, _, err := client.GroupManagementAPI.GroupsControllerGetUsersForGroup(context.Background(), accountID, groupId).Execute()
	utils.CheckError(err)
	return users
}

func updateGroupsUsersInMongo(db *mongo.Database, group string, users account.GroupUserListDto) error {
	userCollection := db.Collection("users")
	groupCollection := db.Collection("groups")
	groupID, _ := utils.FindDocumentIDByUUID(context.Background(), groupCollection, "uuid", group)
	userIDs := make([]primitive.ObjectID, len(users.Items))
	for i, user := range users.Items {
		userID, err := utils.FindDocumentIDByUUID(context.Background(), userCollection, "uid", user.Uid)
		if err != nil {
			fmt.Printf("User not found: %s\n", user.Uid)
		}
		userIDs[i] = userID
	}
	filter := bson.M{"_id": groupID}
	update := bson.M{"$set": bson.M{"userRefs": userIDs}}
	_, err := groupCollection.UpdateOne(context.Background(), filter, update)
	return err
}

func updateUserGroupsInMongo(db *mongo.Database, id string, groups []account.AccountGroupDto) error {
	userCollection := db.Collection("users")
	groupCollection := db.Collection("groups")
	userID, _ := utils.FindDocumentIDByUUID(context.Background(), userCollection, "uid", id)
	groupIDs := make([]primitive.ObjectID, len(groups))
	for i, group := range groups {
		groupID, err := utils.FindDocumentIDByUUID(context.Background(), groupCollection, "uuid", group.Uuid)
		if err != nil {
			fmt.Printf("Group not found: %s\n", group.Uuid)
		}
		groupIDs[i] = groupID
	}
	filter := bson.M{"_id": userID}
	update := bson.M{"$set": bson.M{"groupRefs": groupIDs}}
	_, err := userCollection.UpdateOne(context.Background(), filter, update)
	return err
}
