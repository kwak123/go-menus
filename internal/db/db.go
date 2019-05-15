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
	menu := Menu{ID: "stub", Name: "Test", Location: "somewhere", ItemList: []Item{}}
	err := Database.Collection(menuCollectionName).Drop(context.TODO())
	if err != nil {
		println("Failed to refresh db: %s", err)
	} else {
		println("DB refreshed")
	}
	Database.Collection(menuCollectionName).InsertOne(context.TODO(), menu)
}

// GetAllMenus fetches all available menus
func GetAllMenus() []*Menu {
	findOptions := options.Find()
	var menuList []*Menu

	cursor, err := Database.Collection(menuCollectionName).Find(context.TODO(), bson.D{}, findOptions)

	if err != nil {
		println("Failed to get menus")
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var m Menu
		err := cursor.Decode(&m)

		if err != nil {
			println("Failed to decode cursor")
		}
		menuList = append(menuList, &m)
	}

	return menuList
}

// Wrapper for adding menu to db
func AddMenu(m Menu) {
	_, err := Database.Collection(menuCollectionName).InsertOne(context.TODO(), m)
	if err != nil {
		println("Failed to add menu")
	}
}

// GetMenu finds desired menu by menuID
func GetMenu(menuID string) Menu {
	print(menuID)
	menu := Menu{}
	filter := bson.D{{"id", menuID}}
	err := Database.Collection(menuCollectionName).FindOne(context.TODO(), filter).Decode(&menu)

	if err != nil {
		println("Failed to find menu")
		println(err.Error())
	}
	return menu
}

// AddItemToMenu adds Item to specified Menu.ItemList
func AddItemToMenu(menuID string, newItem Item) {
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
	filter := bson.D{{"id", menuID}}

	itemToDelete := bson.D{{"id", itemID}}

	deleteOperation := bson.D{
		{"$pull", bson.D{
			{"itemlist", itemToDelete},
		}},
	}

	_, err := Database.Collection(menuCollectionName).UpdateOne(context.TODO(), filter, deleteOperation)

	if err != nil {
		println("")
		fmt.Printf("Failed to delete item from menu: %s", err.Error())
	}
}

// UpdateItemInMenu replaces specified Item from specific menu
// Item.ID will be kept the same
// TODO: Do this with MongoDB in one operation
func UpdateItemInMenu(menuID string, updatedItem Item) {
	menu := Menu{}
	filter := bson.D{{"id", menuID}}
	err := Database.Collection(menuCollectionName).FindOne(context.TODO(), filter).Decode(&menu)

	if err != nil {
		println("")
		fmt.Printf("Failed to find menu for update: %s", err.Error())
	}

	menu.UpdateItem(updatedItem)
	_, err = Database.Collection(menuCollectionName).ReplaceOne(context.TODO(), filter, menu)

	if err != nil {
		println("")
		fmt.Printf("Failed to update item in menu: %s", err.Error())
	}
}
