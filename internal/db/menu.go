package db

// Item of food
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// provider string
}

// Menu with list of items
type Menu struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ItemList []Item `json:"itemList"`
}

func (menu *Menu) AddItem(i Item) {
	menu.ItemList = append(menu.ItemList, i)
}

// Just names for learning haha
func (menu *Menu) UpdateItem(itemID string, newName string) {
	for i := 0; i < len(menu.ItemList); i++ {
		if menu.ItemList[i].ID == itemID {
			menu.ItemList[i].Name = newName
		}
	}
}

// man this is inefficient
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
