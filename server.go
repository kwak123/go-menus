package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"go-menus/internal/db"

	"github.com/gorilla/mux"

	"github.com/google/uuid"
)

type addBody struct {
	MenuID string `json:"menuId"`
}

type modifierBody struct {
	MenuID string  `json:"menuId"`
	Item   db.Item `json:"item"`
}

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
	apiPrefix := "/api/"
	menuID := r.URL.Path[len(apiPrefix):]
	if menuID == "" {
		handleGetAllMenus(w, r)
	} else {
		handleGetMenu(w, r, menuID)
	}
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

func handleGetMenu(w http.ResponseWriter, r *http.Request, menuID string) {
	m := db.GetMenu(menuID)
	menuJSON, err := json.Marshal(m)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(menuJSON)
}

func handleGetAllMenus(w http.ResponseWriter, r *http.Request) {
	menus := db.GetAllMenus()
	menuJSON, err := json.Marshal(menus)

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
	addBody := &addBody{}

	// Create uuid for the new item
	newID := uuid.New()
	stringifiedID := newID.String()
	item := db.Item{ID: stringifiedID}

	err := parseBodyForJSON(w, r, &addBody)

	if err != nil {
		fmt.Fprint(w, "Failed to parse")
	}

	db.AddItemToMenu(addBody.MenuID, item)
	handleGetMenu(w, r, addBody.MenuID)
}

func handleDeleteItem(w http.ResponseWriter, r *http.Request) {
	deleteBody := modifierBody{}

	err := parseBodyForJSON(w, r, &deleteBody)

	if err != nil {
		fmt.Fprintf(w, "Failed to delete item")
	}

	db.DeleteItemFromMenu(deleteBody.MenuID, deleteBody.Item.ID)
	handleGetMenu(w, r, deleteBody.MenuID)
}

// Only name, one thing at a time for now haha
func handleUpdateItem(w http.ResponseWriter, r *http.Request) {
	updateBody := modifierBody{}

	err := parseBodyForJSON(w, r, &updateBody)

	if err != nil {
		fmt.Fprintf(w, "Failed to update item")
	}

	db.UpdateItemInMenu(updateBody.MenuID, updateBody.Item)
	handleGetMenu(w, r, updateBody.MenuID)
}

func clientRouterHandler(w http.ResponseWriter, r *http.Request) {
	entryPoint := "dist/index.html"
	http.ServeFile(w, r, entryPoint)
}

func main() {
	db.InitializeDb()
	routeHandler := makeRouteHandler()

	routeMux := mux.NewRouter()
	routeMux.PathPrefix("/api/").HandlerFunc(routeHandler)

	fs := http.FileServer(http.Dir("dist"))
	// Catch-all for react routing
	routeMux.PathPrefix("/app/").HandlerFunc(clientRouterHandler)
	routeMux.PathPrefix("/").Handler(fs)

	log.Fatal(http.ListenAndServe(":3001", routeMux))
}
