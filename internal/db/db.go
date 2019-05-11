package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var menuCollectionName = "menus"

// Database references the initialized mongo database
var Database *mongo.Database
var hasInitialized = false

var menu = Menu{ID: "stub", Name: "Test", ItemList: []Item{}}

// InitializeDb starts the MongoDB client
func InitializeDb() {
	if !hasInitialized {
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

		if err != nil {
			println("Failed to initialize db client: %s", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)

		if err != nil {
			println("Failed to connect to db: %s", err)
		}

		hasInitialized = true
		println("Connected to db")
		Database = client.Database("test")
		refreshDb()
		cancel()
	} else {
		println("Already connected")
	}
}

func refreshDb() {
	err := Database.Collection(menuCollectionName).Drop(context.TODO())
	if err != nil {
		println("Failed to refresh db: %s", err)
	} else {
		println("DB refreshed")
	}
	Database.Collection(menuCollectionName).InsertOne(context.TODO(), menu)
}

// GetMenu finds desired menu by menuID
func GetMenu(menuID string) Menu {
	menu := Menu{}
	filter := bson.D{{"id", menuID}}
	err := Database.Collection(menuCollectionName).FindOne(context.TODO(), filter).Decode(&menu)

	if err != nil {
		println("Failed to find menu")
	}
	return menu
}

// AddItemToMenu adds Item to specified Menu.ItemList
func AddItemToMenu(menuID string, newItem Item) {
	// menu := Menu{}
	filter := bson.D{{"id", menuID}}

	updateOperation := bson.D{
		{"$push", bson.D{
			{"itemlist", newItem},
		}},
	}

	_, err := Database.Collection(menuCollectionName).UpdateOne(context.TODO(), filter, updateOperation)

	if err != nil {
		println("")
		fmt.Printf("Failed to add item to menu: %s", err.Error())
	}
}

// DeleteItemFromMenu deletes specified ItemID from specified Menu
func DeleteItemFromMenu(menuID string, itemID string) {
	fmt.Printf("Need to implement using the id: %s", menuID)
	menu.DeleteItem(itemID)
}

// UpdateItemInMenu replaces specified Item from specific menu
// Item.ID will be kept the same
func UpdateItemInMenu(menuID string, updatedItem Item) {
	fmt.Printf("Need to implement using the id: %s", menuID)
	menu.UpdateItem(updatedItem)
}
