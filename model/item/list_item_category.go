package item

import (
	"fmt"
	database "github.com/pardev/cantik-mart/model/db"
)

func GetItemCategoryListMap() map[int](map[string]string) {
	rows, err := database.ExecuteQuery("SELECT * FROM t_item_cat")
	if len(rows) > 0 && err == nil {
		return rows
	} else {
		if err != nil {
			fmt.Println("ERR : GetItemListMap - ", err)
		}
		return nil
	}
}

func GetItemCategoryListArray() [][]string {
	rows := GetItemCategoryListMap()
	var result [][]string
	if rows == nil {
		return result
	} else {
		for i := 0; i < len(rows); i++ {
			id := fmt.Sprint(rows[i]["id"])
			name := fmt.Sprint(rows[i]["name"])
			temp := []string{id, name}
			result = append(result, temp)
		}
		return result
	}
}
