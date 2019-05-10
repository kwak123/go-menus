package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Types
type item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// provider string
}

type menu struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ItemList []item `json:"itemList"`
}

func (menu *menu) addItemToMenu(i item) {
	menu.ItemList = append(menu.ItemList, i)
}

// This closure is probably unnecessary once converting to db
func makeRouteHandler() http.HandlerFunc {
	m := &menu{ID: "123", Name: "Test"}

	return func(w http.ResponseWriter, r *http.Request) {
		apiPrefix := "/api/"
		pathWithoutAPIPrefix := r.URL.Path[len(apiPrefix):]

		if r.Method == "GET" {
			handleGetMenu(w, r, m)
		} else if r.Method == "POST" {
			switch pathWithoutAPIPrefix {
			case "add":
				handleAddItemToMenu(w, r, m)
			default:
				w.WriteHeader(400)
			}
		} else {
			fmt.Fprint(w, "Fai")
		}
	}
}

// Only one get for now
func handleGetRequest(w http.ResponseWriter, r *http.Request, m *menu) {
	handleGetMenu(w, r, m)
}

// Only have one post for now
func handlePostRequest(w http.ResponseWriter, r *http.Request, m *menu) {
	handleAddItemToMenu(w, r, m)
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
	// TODO: Remove this mock id handler
	id := strconv.Itoa(rand.Int())
	// Initialize item
	item := item{ID: id}

	// Try to read the body
	itemJSON, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Error parsing body: %s", err)
	}

	err = json.Unmarshal(itemJSON, &item)

	if err != nil {
		fmt.Fprintf(w, "Error converting JSON to item: %s", err)
	}

	m.addItemToMenu(item)
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
