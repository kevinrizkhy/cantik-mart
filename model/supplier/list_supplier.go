package supplier

import (
	"fmt"
	database "github.com/pardev/cantik-mart/model/db"
)

//name, address, phone, email

func GetSupplierListMap() map[int](map[string]string) {
	rows, err := database.ExecuteQuery("SELECT t_supplier.id, t_supplier.name, t_supplier.address, t_supplier.phone, t_supplier.email, t_supplier.status FROM t_supplier ORDER BY name ASC")
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
			status := ""
			if rows[i]["status"] == "true" {
				status = "Aktif"
			} else {
				status = "Tidak Aktif"
			}
			js_str += "["
			js_str += rows[i]["id"] + ",'" + rows[i]["name"] + "','" + rows[i]["address"] + "','" + rows[i]["phone"] + "','" + rows[i]["email"] + "','" + status
			js_str += "','<button onclick=\"editSupplier(this)\" class=\"btn btn-warning\">Edit</button']"
			if (i + 1) != len(rows) {
				js_str += ","
			}
		}
		js_str += "]"
		return js_str
	}
}

func GetSupplier() [][]string {
	rows := GetSupplierListMap()
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
