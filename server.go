package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"./internal/db"
)

// This closure is probably unnecessary once converting to db
func makeRouteHandler() http.HandlerFunc {
	m := &db.Menu{ID: "123", Name: "Test"}

	return func(w http.ResponseWriter, r *http.Request) {
		apiPrefix := "/api/"
		pathWithoutAPIPrefix := r.URL.Path[len(apiPrefix):]

		if r.Method == "GET" {
			handleGetMenu(w, r, m)
		} else if r.Method == "POST" {
			switch pathWithoutAPIPrefix {
			case "add":
				handleAddItemToMenu(w, r, m)
			case "delete":
				handleDeleteItem(w, r, m)
			default:
				w.WriteHeader(400)
			}
		} else if r.Method == "PUT" {
			handleUpdateItem(w, r, m)
		} else {
			fmt.Fprint(w, "Fai")
		}
	}
}

// Only one get for now
func handleGetRequest(w http.ResponseWriter, r *http.Request, m *db.Menu) {
	handleGetMenu(w, r, m)
}

// Only have one post for now
func handlePostRequest(w http.ResponseWriter, r *http.Request, m *db.Menu) {
	handleAddItemToMenu(w, r, m)
}

func handleGetMenu(w http.ResponseWriter, r *http.Request, m *db.Menu) {
	menuJSON, err := json.Marshal(m)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(menuJSON)
}

// TODO: Expand to add an item
func handleAddItemToMenu(w http.ResponseWriter, r *http.Request, m *db.Menu) {
	// TODO: Remove this mock id handler
	id := strconv.Itoa(rand.Int())
	// Initialize item
	item := db.Item{ID: id}

	// Try to read the body
	itemJSON, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Error parsing body: %s", err)
	}

	err = json.Unmarshal(itemJSON, &item)

	if err != nil {
		fmt.Fprintf(w, "Error converting JSON to item: %s", err)
	}

	m.AddItem(item)
	handleGetMenu(w, r, m)
}

// Only name, one thing at a time for now haha
func handleUpdateItem(w http.ResponseWriter, r *http.Request, m *db.Menu) {
	updateItem := db.Item{}
	updateJSON, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Error parsing body: %s", err)
	}

	err = json.Unmarshal(updateJSON, &updateItem)

	if err != nil {
		fmt.Fprintf(w, "Error converting JSON to item: %s", err)
	}

	m.UpdateItem(updateItem.ID, updateItem.Name)
	handleGetMenu(w, r, m)
}

func handleDeleteItem(w http.ResponseWriter, r *http.Request, m *db.Menu) {
	itemToDelete := db.Item{}
	deleteJSON, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Error parsing body: %s", err)
	}

	err = json.Unmarshal(deleteJSON, &itemToDelete)

	if err != nil {
		fmt.Fprintf(w, "Error converting JSON to item: %s", err)
	}

	m.DeleteItem(itemToDelete.ID)
	handleGetMenu(w, r, m)
}

func main() {
	routeHandler := makeRouteHandler()

	fs := http.FileServer(http.Dir("dist"))
	http.HandleFunc("/api/", routeHandler)
	http.Handle("/", fs)

	// http.HandleFunc("/getMenu", handleAddItemToMenu)

	log.Fatal(http.ListenAndServe(":3001", nil))
}
