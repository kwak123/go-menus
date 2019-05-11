package db

import (
	"fmt"
)

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
