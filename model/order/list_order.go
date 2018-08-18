package order

import (
	"fmt"
	database "github.com/pardev/cantik-mart/model/db"
)

func GetOrderListMap() map[int](map[string]string) {
	rows, err := database.ExecuteQuery("select t_order.id, t_order.name, t_supplier.name AS supplier, t_order.total, t_order.note from t_order left join t_supplier ON t_order.id =t_supplier.id ORDER BY t_order.date_created DESC")
	if len(rows) > 0 && err == nil {
		return rows
	} else {
		if err != nil {
			fmt.Println("ERR : GetOrderListMap - ", err)
		}
		return nil
	}
}

func GetOrder() string {
	rows := GetOrderListMap()
	if rows == nil {
		return ""
	} else {
		js_str := "["
		for i := 0; i < len(rows); i++ {
			js_str += "['"
			js_str += rows[i]["id"] + "','" + rows[i]["name"] + "','" + rows[i]["supplier"] + "','" + rows[i]["total"] + "','" + rows[i]["note"] + "','<button onclick=\"detailOrder(this)\" class=\"btn btn-info\"><i class=\"fa fa-info-circle\"></i></button>&nbsp&nbsp&nbsp<button onclick=\"editOrder(this)\" class=\"btn btn-warning\"><i class=\"fa fa-pencil-square-o\"></i></button>&nbsp&nbsp&nbsp<button onclick=\"deleteOrder(this)\" class=\"btn btn-danger\"><i class=\"fa fa-trash-o\"></i></button>']"
			if (i + 1) != len(rows) {
				js_str += ","
			}
		}
		js_str += "]"
		return js_str
	}
}
