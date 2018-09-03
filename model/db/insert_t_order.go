package db

import ()

func InsertOrder(number, order_by, order_to, total, note string) bool {
	_, err := ExecuteQuery("INSERT INTO t_order(number, order_by, order_to, total, note) VALUES ($1,$2,$3,$4,$5);", number, order_by, order_to, total, note) //returning id
	if err == nil {
		return true
	} else {
		return false
	}
}
