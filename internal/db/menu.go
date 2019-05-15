package db

// Item of food
type Item struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Provider string `json:"provider"`
	// provider string
}

// Menu with list of items
type Menu struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	ItemList []Item `json:"itemList"`
}

// AddItem appends Item to ItemList
func (menu *Menu) AddItem(i Item) {
	menu.ItemList = append(menu.ItemList, i)
}

// UpdateItem will find itemID in menu and replace it
func (menu *Menu) UpdateItem(updatedItem Item) {
	for i := 0; i < len(menu.ItemList); i++ {
		if menu.ItemList[i].ID == updatedItem.ID {
			menu.ItemList[i] = updatedItem
			return
		}
	}
}

// DeleteItem excises itemID from ItemList
func (menu *Menu) DeleteItem(itemID string) {
	indexToDelete := -1
	for i := 0; i < len(menu.ItemList); i++ {
		if menu.ItemList[i].ID == itemID {
			indexToDelete = i
		}
	}

	if indexToDelete > -1 {
		menu.ItemList = append(menu.ItemList[:indexToDelete], menu.ItemList[indexToDelete+1:]...)
	}
}
