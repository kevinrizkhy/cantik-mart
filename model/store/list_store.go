package store

import (
	"fmt"
	database "github.com/pardev/cantik-mart/model/db"
)

func GetStoreListMap() map[int](map[string]string) {
	rows, err := database.ExecuteQuery("select * from t_store ORDER BY name ASC")
	if len(rows) > 0 && err == nil {
		return rows
	} else {
		if err != nil {
			fmt.Println("ERR : GetStoreList - ", err)
		}
		return nil
	}
}

func GetStore() [][]string {
	rows := GetStoreListMap()
	var result [][]string
	if rows == nil {
		return result
	} else {
		for i := 0; i < len(rows); i++ {
			id := fmt.Sprint(rows[i]["id"])
			name := fmt.Sprint(rows[i]["name"])
			address := fmt.Sprint(rows[i]["address"])
			phone := fmt.Sprint(rows[i]["phone"])
			button := fmt.Sprint("<button onclick=\"detailUser(this)\" class=\"btn btn-info\">Detail</button>&nbsp&nbsp&nbsp<button onclick=\"editUser(this)\" class=\"btn btn-warning\">Edit</button>&nbsp&nbsp&nbsp<button onclick=\"deleteUser(this)\" class=\"btn btn-danger\">Delete</button>")
			temp := []string{id, name, address, phone, button}
			result = append(result, temp)
		}
		return result
	}
}
