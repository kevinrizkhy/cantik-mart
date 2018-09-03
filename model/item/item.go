package item

import (
	"fmt"
	database "github.com/pardev/cantik-mart/model/db"
)

func GetItemMap(id string) map[int](map[string]string) {
	rows, err := database.ExecuteQuery("select t_item.id, t_item.name, t_item_cat.name AS category, t_item.sell from t_item left join t_item_cat ON t_item.category=t_item_cat.id WHERE t_item.id = $1 ORDER BY name ASC", id)
	if len(rows) > 0 && err == nil {
		return rows
	} else {
		if err != nil {
			fmt.Println("ERR : GetItemMap - ", err)
		}
		return nil
	}
}

func GetItem(id string) [][]string {
	rows := GetItemMap(id)
	var data [][]string
	if rows == nil {
		return nil
	} else {
		for i := 0; i < len(rows); i++ {
			id := rows[i]["id"]
			name := rows[i]["name"]
			category := rows[i]["category"]
			sell := rows[i]["sell"]
			tempdata := [][]string{{id, name, category, sell}}
			data = append(data, tempdata...)
		}
		return data
	}
}
