package order

import (
	"fmt"
	database "github.com/pardev/cantik-mart/model/db"
)

//number, total, order_by, sent_to, note, status,
func GetOrderListMap() map[int](map[string]string) {
	rows, err := database.ExecuteQuery("select t_order.id, t_order.number, t_order.total, t_user.name, t_supplier.name, t_order.note, t_order_status.name from t_order LEFT JOIN t_user ON t_order.order_by = t_user.id LEFT JOIN t_order_status ON t_order.status = t_order_status.name ORDER BY t_order.date_created DESC")
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
			js_str += rows[i]["id"] + "','" + rows[i]["name"] + "','" + rows[i]["order_by"] + "','" + rows[i]["sent_to"] + "','" + rows[i]["note"] + "','" + rows[i]["status"] + "','<button onclick=\"detailOrder(this)\" class=\"btn btn-info\"><i class=\"fa fa-info-circle\"></i></button>&nbsp&nbsp&nbsp<button onclick=\"editOrder(this)\" class=\"btn btn-warning\"><i class=\"fa fa-pencil-square-o\"></i></button>&nbsp&nbsp&nbsp<button onclick=\"deleteOrder(this)\" class=\"btn btn-danger\"><i class=\"fa fa-trash-o\"></i></button>']"
			if (i + 1) != len(rows) {
				js_str += ","
			}
		}
		js_str += "]"
		return js_str
	}
}
