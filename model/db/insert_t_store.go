package db

import ()

func InsertStore(id, name, address, phone string) bool {
	_, err := ExecuteQuery("INSERT INTO t_user (id, name, email, address, phone) VALUES ($1,$2,$3,$4);", id, name, address, phone)
	if err == nil {
		return true
	} else {
		return false
	}
}
