package db

import ()

func UpdateSupplier(id, name, address, phone, email, status string) bool {
	_, err := ExecuteQuery("UPDATE t_supplier SET name=$1, address=$2, phone=$3, email=$4, status=$5 WHERE id=$6;", name, address, phone, email, status, id) //returning id
	if err == nil {
		return true
	} else {
		return false
	}
}
