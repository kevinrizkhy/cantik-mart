package db

import (
	"fmt"
)

func UpdateStore(id, name, address, phone, status string) bool {
	_, err := ExecuteQuery("UPDATE t_store SET name=$1, address=$2, phone=$3,status=$4 WHERE id=$5;", name, address, phone, status, id)
	if err != nil {
		fmt.Println("Err : UpdateStore - ", err)
		return false
	}
	return true
}
