package user

import (
	"fmt"
	database "github.com/pardev/cantik-mart/model/db"
)

func GetUserDetail(id string) map[string]interface{} {
	rows, err := database.ExecuteQuery("SELECT t_user.id, t_user.name, t_user.email, t_user.address, t_user.phone, t_role.name AS role, t_store.name AS store_name, t_store.phone AS store_phone, t_store.address AS store_address FROM t_user LEFT JOIN t_role ON t_user.role=t_role.id LEFT JOIN t_store ON t_user.store=t_store.id WHERE t_user.id = $1", id)
	if len(rows) > 0 && err == nil {
		result := make(map[string]interface{})
		result["id"] = rows[0]["id"]
		result["name"] = rows[0]["name"]
		result["email"] = rows[0]["email"]
		result["address"] = rows[0]["company"]
		result["role"] = rows[0]["role"]
		result["phone"] = rows[0]["phone"]
		result["store_name"] = rows[0]["store_name"]
		result["store_phone"] = rows[0]["store_phone"]
		result["store_address"] = rows[0]["store_address"]
		return result
	} else {
		if err != nil {
			fmt.Println("ERR : GetUserDetail - ", err)
		}
		return nil
	}
}
