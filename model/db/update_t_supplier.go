package db

import ()

func UpdateSupplier(id, name, address, phone, description string) bool {
	_, err := ExecuteQuery("UPDATE t_supplier SET name=$1, address=$2, phone=$3, description=$4 WHERE id=$5;", name, address, phone, description, id) //returning id
	if err == nil {
		return true
	} else {
		return false
	}
}
