package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// type item struct {
// 	name     string
// 	provider string
// }

type menu struct {
	ItemList []string `json: itemList`
}

// func (menu *menu) addItemToMenu(item string) {
// 	menu.itemList <- item
// }

func main() {
	menu := &menu{ItemList: []string{"Test"}}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// item := "Test"
		// go menu.addItemToMenu(item)
		// menu.itemList <- "Hello"
		menuJson, err := json.Marshal(menu)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(menuJson)
	})
	// http.HandleFunc("/getMenu", handleAddItemToMenu)

	log.Fatal(http.ListenAndServe(":3001", nil))
}
