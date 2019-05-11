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
		switch r.Method {
		case "GET":
			handleGetRequest(w, r, m)
		case "POST":
			handlePostRequest(w, r, m)
		case "PUT":
			handlePutRequest(w, r, m)
		default:
			fmt.Fprint(w, "Invalid request")
		}
	}
}

// Only one get for now
func handleGetRequest(w http.ResponseWriter, r *http.Request, m *db.Menu) {
	handleGetMenu(w, r, m)
}

// Only have one post for now
func handlePostRequest(w http.ResponseWriter, r *http.Request, m *db.Menu) {
	apiPrefix := "/api/"
	pathWithoutAPIPrefix := r.URL.Path[len(apiPrefix):]

	switch pathWithoutAPIPrefix {
	case "add":
		handleAddItemToMenu(w, r, m)
	case "delete":
		handleDeleteItem(w, r, m)
	default:
		w.WriteHeader(400)
	}
}

func handlePutRequest(w http.ResponseWriter, r *http.Request, m *db.Menu) {
	handleUpdateItem(w, r, m)
}

func handleGetMenu(w http.ResponseWriter, r *http.Request, m *db.Menu) {
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
func handleAddItemToMenu(w http.ResponseWriter, r *http.Request, m *db.Menu) {
	// TODO: Remove this mock id handler
	id := strconv.Itoa(rand.Int())
	// Initialize item
	item := db.Item{ID: id}

	err := parseBodyForJSON(w, r, &item)

	if err != nil {
		fmt.Fprint(w, "Failed to parse")
	}

	m.AddItem(item)
	handleGetMenu(w, r, m)
}

// Only name, one thing at a time for now haha
func handleUpdateItem(w http.ResponseWriter, r *http.Request, m *db.Menu) {
	updateItem := db.Item{}

	err := parseBodyForJSON(w, r, &updateItem)

	if err != nil {
		fmt.Fprintf(w, "Failed to update item")
	}

	m.UpdateItem(updateItem.ID, updateItem.Name)
	handleGetMenu(w, r, m)
}

func handleDeleteItem(w http.ResponseWriter, r *http.Request, m *db.Menu) {
	deleteItem := db.Item{}

	err := parseBodyForJSON(w, r, &deleteItem)

	if err != nil {
		fmt.Fprintf(w, "Failed to delete item")
	}

	m.DeleteItem(deleteItem.ID)
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
