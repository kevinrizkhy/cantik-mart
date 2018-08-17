package db

import ()

func InsertUser(name, email, address, phone, role, store, password string) bool {
	_, err := ExecuteQuery("INSERT INTO t_user (name, email, address, phone, role, store,password) VALUES ($1,$2,$3,$4,$5,$6,$7);", email, name, address, phone, role, store, password)
	if err == nil {
		return true
	} else {
		return false
	}
}
