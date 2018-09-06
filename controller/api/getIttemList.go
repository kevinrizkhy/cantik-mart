package api

import (
	"encoding/json"
	item_model "github.com/pardev/cantik-mart/model/item"
	"net/http"
)

func GetItemList(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query()["search"][0]
	item_arr := item_model.GetItemSearch(item)
	if len(item_arr) == 0 {
		w.WriteHeader(400)
	}
	jsonData, _ := json.Marshal(item_arr)
	w.Header().Add("Access-Control-Allow-Origin", "127.0.0.1:8080")
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
