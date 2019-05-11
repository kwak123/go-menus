package db

import (
	"fmt"
)

var menu = new(Menu)

// GetMenu finds desired menu by menuID
func GetMenu(menuID string) Menu {
	fmt.Printf("Implement using the id: %s", menuID)
	return *menu
}

// AddItemToMenu adds Item to specified Menu.ItemList
func AddItemToMenu(menuID string, newItem Item) {
	menu.AddItem(newItem)
}

// DeleteItemFromMenu deletes specified ItemID from specified Menu
func DeleteItemFromMenu(menuID string, itemID string) {
	menu.DeleteItem(itemID)
}

// UpdateItemInMenu replaces specified Item from specific menu
// Item.ID will be kept the same
func UpdateItemInMenu(menuID string, updatedItem Item) {
	menu.UpdateItem(updatedItem)
}
