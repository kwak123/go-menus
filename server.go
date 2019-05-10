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

// This closure is probably unnecessary once converting to db
func makeRouteHandler(m *menu) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path[1:] {
		case "":
			handleGetMenu(w, r, m)
		case "add":
			handleAddItemToMenu(w, r, m)
		}
	}
}

func handleGetMenu(w http.ResponseWriter, r *http.Request, m *menu) {
	menuJSON, err := json.Marshal(m)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(menuJSON)
}

// TODO: Expand to add an item
func handleAddItemToMenu(w http.ResponseWriter, r *http.Request, m *menu) {
	item := item{Name: "New Item"}
	m.addItemToMenu(item)
	handleGetMenu(w, r, m)
}

func main() {
	menu := new(menu)
	routeHandler := makeRouteHandler(menu)

	http.HandleFunc("/", routeHandler)
	// http.HandleFunc("/getMenu", handleAddItemToMenu)

	log.Fatal(http.ListenAndServe(":3001", nil))
}
