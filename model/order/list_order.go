package order

import (
	"fmt"
	database "github.com/pardev/cantik-mart/model/db"
)

//number, total, order_by, sent_to, note, status,
func GetOrderListMap() map[int](map[string]string) {
	rows, err := database.ExecuteQuery("select t_order.id, t_order.number, t_user.name AS order_by, t_supplier.name AS order_to, t_order.total, t_order.note, t_order_status.name AS status from t_order LEFT JOIN t_user ON t_order.order_by = t_user.id LEFT JOIN t_order_status ON t_order.status = t_order_status.id LEFT JOIN t_supplier ON t_order.order_to = t_supplier.id ORDER BY t_order.date_created DESC")
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
			js_str += rows[i]["id"] + "','" + rows[i]["number"] + "','" + rows[i]["order_by"] + "','" + rows[i]["order_to"] + "','" + rows[i]["total"] + "','" + rows[i]["note"] + "','" + rows[i]["status"] + "','<button onclick=\"detailOrder(this)\" class=\"btn btn-info\"><i class=\"fa fa-info-circle\"></i></button>&nbsp&nbsp&nbsp<button onclick=\"editOrder(this)\" class=\"btn btn-warning\"><i class=\"fa fa-pencil-square-o\"></i></button>']"
			if (i + 1) != len(rows) {
				js_str += ","
			}
		}
		js_str += "]"
		return js_str
	}
}
