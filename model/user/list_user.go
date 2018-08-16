package user

import (
	"fmt"
	database "github.com/pardev/cantik-mart/model/db"
)

//id, name, email, address, phone, role, store, password

func GetUserListMap() map[int](map[string]string) {
	rows, err := database.ExecuteQuery("SELECT t_user.id, t_user.name, t_user.email, t_user.address, t_user.phone, t_role.name AS role, t_store.name AS store FROM t_user LEFT JOIN t_role ON t_user.role=t_role.id LEFT JOIN t_store ON t_user.store = t_store.id ORDER BY t_user.role,t_user.name ASC")
	if len(rows) > 0 && err == nil {
		return rows
	} else {
		if err != nil {
			fmt.Println("ERR : GetUserWebList - ", err)
		}
		return nil
	}
}

func GetUserList() string {
	rows := GetUserListMap()
	if rows == nil {
		return ""
	} else {
		js_str := "["
		for i := 0; i < len(rows); i++ {
			js_str += "["
			js_str += rows[i]["id"] + ",'" + rows[i]["name"] + "','" + rows[i]["email"] + "','" + rows[i]["address"] + "','" + rows[i]["phone"] + "','" + rows[i]["role"] + "','" + rows[i]["store"]
			js_str += "','<button onclick=\"detailUser(this)\" class=\"btn btn-info\">Detail</button>&nbsp&nbsp&nbsp<button onclick=\"editUser(this)\" class=\"btn btn-warning\">Edit</button>&nbsp&nbsp&nbsp<button onclick=\"deleteUser(this)\" class=\"btn btn-danger\">Delete</button>']"
			if (i + 1) != len(rows) {
				js_str += ","
			}
		}
		js_str += "]"
		return js_str
	}
}
