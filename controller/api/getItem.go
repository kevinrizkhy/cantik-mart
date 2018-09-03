package api

import (
	"encoding/json"
	item_model "github.com/pardev/cantik-mart/model/item"
	"net/http"
)

func GetItem(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query()["id"][0]
	item_arr := item_model.GetItem(item)
	jsonData, _ := json.Marshal(item_arr)
	w.Header().Add("Access-Control-Allow-Origin", "127.0.0.1:8080")
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
