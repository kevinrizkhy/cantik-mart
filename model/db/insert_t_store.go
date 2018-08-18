package db

import ()

<<<<<<< HEAD
func InsertStore(id, name, address, phone string) bool {
	_, err := ExecuteQuery("INSERT INTO t_store (id, name, email, address, phone) VALUES ($1,$2,$3,$4);", id, name, address, phone)
=======
func InsertStore(name, address, phone string) bool {
	_, err := ExecuteQuery("INSERT INTO t_store (name, address, phone) VALUES ($1,$2,$3);", name, address, phone)
>>>>>>> 7ae367e299f05856dc97503e5710535653feb006
	if err == nil {
		return true
	} else {
		return false
	}
}
