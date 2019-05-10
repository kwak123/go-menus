package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type item struct {
	Name string `json:"name"`
	// provider string
}

type menu struct {
	ItemList []item `json:"itemList"`
}

func (menu *menu) addItemToMenu(i item) {
	menu.ItemList = append(menu.ItemList, i)
}

func main() {
	menu := &menu{ItemList: []item{{Name: "Hello"}}}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// item := "Test"
		// go menu.addItemToMenu(item)
		// menu.itemList <- "Hello"
		item := item{Name: "New Item"}
		menu.addItemToMenu(item)
		menuJSON, err := json.Marshal(menu)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(menuJSON)
	})
	// http.HandleFunc("/getMenu", handleAddItemToMenu)

	log.Fatal(http.ListenAndServe(":3001", nil))
}
