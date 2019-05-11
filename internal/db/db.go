package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database
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
		database = client.Database("test")
		refreshDb()
		cancel()
	} else {
		println("Already connected")
	}
}

func refreshDb() {
	err := database.Collection("menu").Drop(context.TODO())
	if err != nil {
		println("Failed to refresh db: %s", err)
	} else {
		println("DB refreshed")
	}
	database.Collection("menu")
}

var menu = Menu{ID: "123", Name: "Test"}

// GetMenu finds desired menu by menuID
func GetMenu(menuID string) Menu {
	fmt.Printf("Need to implement using the id: %s", menuID)
	return menu
}

// AddItemToMenu adds Item to specified Menu.ItemList
func AddItemToMenu(menuID string, newItem Item) {
	fmt.Printf("Need to implement using the id: %s", menuID)
	menu.AddItem(newItem)
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
