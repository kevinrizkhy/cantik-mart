package db

import ()

func InsertSupplier(name, address, phone, email string) bool {
	_, err := ExecuteQuery("INSERT INTO t_supplier (name, address, phone, email) VALUES ($1,$2,$3,$4);", name, address, phone, email) //returning id
	if err == nil {
		return true
	} else {
		return false
	}
}
