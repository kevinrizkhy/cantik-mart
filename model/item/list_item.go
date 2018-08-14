package item

import (
	"fmt"
	database "github.com/pardev/cantik-mart/model/db"
)

func GetItemListMap() map[int](map[string]string) {
	rows, err := database.ExecuteQuery("select t_item.id, t_item.name, t_item_cat.name AS category, t_item.buy, t_item.margin, t_item.sell, t_item.description from t_item left join t_item_cat ON t_item.category=t_item_cat.id ORDER BY name ASC")
	if len(rows) > 0 && err == nil {
		return rows
	} else {
		if err != nil {
			fmt.Println("ERR : GetItemListMap - ", err)
		}
		return nil
	}
}

func GetItemList() string {
	rows := GetItemListMap()
	if rows == nil {
		return ""
	} else {
		js_str := "["
		for i := 0; i < len(rows); i++ {
			js_str += "['"
			js_str += rows[i]["id"] + "','" + rows[i]["name"] + "','" + rows[i]["category"] + "','" + rows[i]["buy"] + "','" + rows[i]["sell"] + "','" + rows[i]["margin"] + "','" + rows[i]["description"]
			js_str += "','<button onclick=\"detailItem(this)\" class=\"btn btn-info\">Detail</button>&nbsp&nbsp&nbsp<button onclick=\"editItem(this)\" class=\"btn btn-warning\">Edit</button>&nbsp&nbsp&nbsp<button onclick=\"deleteItem(this)\" class=\"btn btn-danger\">Delete</button>']"
			if (i + 1) != len(rows) {
				js_str += ","
			}
		}
		js_str += "]"
		return js_str
	}
}
