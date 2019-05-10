package db

type Item struct {
	Type     string
	Provider string
}

type Menu struct {
	ItemList []Item
}

func (menu *Menu) addItemToMenu(item Item) {
	menu.ItemList.append(item)
}
