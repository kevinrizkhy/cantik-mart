package db

import ()

func InsertSupplier(name, address, phone, description string) bool {
	_, err := ExecuteQuery("INSERT INTO t_supplier (name, address, phone, description) VALUES ($1,$2,$3,$4);", name, address, phone, description) //returning id
	if err == nil {
		return true
	} else {
		return false
	}
}
