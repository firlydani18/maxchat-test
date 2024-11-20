package main

import (
	"encoding/json"
	"log"
	"net/http"

	"maxchat-backend/model"
	"maxchat-backend/utils"
)

var data []model.Item

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	for _, item := range data {
		if item.Code == code {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func FilterItems(w http.ResponseWriter, r *http.Request) {
	modelFilter := r.URL.Query().Get("model")
	techFilters := r.URL.Query()["tech"]

	var filtered []model.Item
	for _, item := range data {
		if modelFilter != "" && item.Model != modelFilter {
			continue
		}
		if len(techFilters) > 0 {
			match := false
			for _, tech := range techFilters {
				if utils.Contains(item.Tech, tech) {
					match = true
					break
				}
			}
			if !match {
				continue
			}
		}
		filtered = append(filtered, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}

func main() {
	var err error

	data, err = utils.LoadData("data.txt")
	if err != nil {
		log.Fatalf("Error loading data: %v", err)
	}

	http.HandleFunc("/items", GetAllItems)
	http.HandleFunc("/item", GetItem)
	http.HandleFunc("/filter", FilterItems)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
