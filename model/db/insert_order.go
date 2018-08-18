package db

import ()

func InsertOrder(name, supplier, total, note string) bool {
	_, err := ExecuteQuery("INSERT INTO t_order(name, supplier, total, note) VALUES ($1,$2,$3,$4);", name, supplier, total, note) //returning id
	if err == nil {
		return true
	} else {
		return false
	}
}
