package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"go-menus/internal/db"
)

// This closure is probably unnecessary once converting to db
func makeRouteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handleGetRequest(w, r)
		case "POST":
			handlePostRequest(w, r)
		case "PUT":
			handlePutRequest(w, r)
		default:
			fmt.Fprint(w, "Invalid request")
		}
	}
}

// Only one get for now
func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	handleGetMenu(w, r)
}

// Only have one post for now
func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	apiPrefix := "/api/"
	pathWithoutAPIPrefix := r.URL.Path[len(apiPrefix):]

	switch pathWithoutAPIPrefix {
	case "add":
		handleAddItemToMenu(w, r)
	case "delete":
		handleDeleteItem(w, r)
	default:
		w.WriteHeader(400)
	}
}

func handlePutRequest(w http.ResponseWriter, r *http.Request) {
	handleUpdateItem(w, r)
}

func handleGetMenu(w http.ResponseWriter, r *http.Request) {
	m := db.GetMenu("stub")
	menuJSON, err := json.Marshal(m)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(menuJSON)
}

func parseBodyForJSON(w http.ResponseWriter, r *http.Request, v interface{}) error {
	JSON, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Error parsing body: %s", err)
	}

	err = json.Unmarshal(JSON, v)

	if err != nil {
		fmt.Fprintf(w, "Error unmarshalling JSON")
	}

	return err
}

// TODO: Expand to add an item
func handleAddItemToMenu(w http.ResponseWriter, r *http.Request) {
	// TODO: Remove this mock id handler
	id := strconv.Itoa(rand.Int())
	// Initialize item
	item := db.Item{ID: id}

	err := parseBodyForJSON(w, r, &item)

	if err != nil {
		fmt.Fprint(w, "Failed to parse")
	}

	db.AddItemToMenu("stub", item)
	handleGetMenu(w, r)
}

// Only name, one thing at a time for now haha
func handleUpdateItem(w http.ResponseWriter, r *http.Request) {
	updateItem := db.Item{}

	err := parseBodyForJSON(w, r, &updateItem)

	if err != nil {
		fmt.Fprintf(w, "Failed to update item")
	}

	db.UpdateItemInMenu("stub", updateItem)
	handleGetMenu(w, r)
}

func handleDeleteItem(w http.ResponseWriter, r *http.Request) {
	deleteItem := db.Item{}

	err := parseBodyForJSON(w, r, &deleteItem)

	if err != nil {
		fmt.Fprintf(w, "Failed to delete item")
	}

	db.DeleteItemFromMenu("stub", deleteItem.ID)
	handleGetMenu(w, r)
}

func main() {
	db.InitializeDb()
	routeHandler := makeRouteHandler()

	fs := http.FileServer(http.Dir("dist"))
	http.HandleFunc("/api/", routeHandler)
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(":3001", nil))
}
