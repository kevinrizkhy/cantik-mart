package db

import ()

func InsertStore(name, address, phone string) bool {
	_, err := ExecuteQuery("INSERT INTO t_store (name, address, phone) VALUES ($1,$2,$3);", name, address, phone)
	if err == nil {
		return true
	} else {
		return false
	}
}
