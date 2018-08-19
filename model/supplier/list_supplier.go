package supplier

import (
	"fmt"
	database "github.com/pardev/cantik-mart/model/db"
)

//name, address, phone, description

func GetSupplierListMap() map[int](map[string]string) {
	rows, err := database.ExecuteQuery("SELECT t_supplier.id, t_supplier.name, t_supplier.address, t_supplier.description FROM t_supplier ORDER BY name ASC")
	if len(rows) > 0 && err == nil {
		return rows
	} else {
		if err != nil {
			fmt.Println("ERR : GetSupplierWebList - ", err)
		}
		return nil
	}
}

func GetSupplierList() string {
	rows := GetSupplierListMap()
	if rows == nil {
		return ""
	} else {
		js_str := "["
		for i := 0; i < len(rows); i++ {
			js_str += "["
			js_str += rows[i]["id"] + ",'" + rows[i]["name"] + "','" + rows[i]["address"] + "','" + rows[i]["phone"] + "','" + rows[i]["description"]
			js_str += "','<button onclick=\"detailSupplier(this)\" class=\"btn btn-info\">Detail</button>&nbsp&nbsp&nbsp<button onclick=\"editSupplier(this)\" class=\"btn btn-warning\">Edit</button>&nbsp&nbsp&nbsp<button onclick=\"deleteSupplier(this)\" class=\"btn btn-danger\">Delete</button>']"
			if (i + 1) != len(rows) {
				js_str += ","
			}
		}
		js_str += "]"
		return js_str
	}
}
