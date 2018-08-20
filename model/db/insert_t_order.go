package db

import ()

func InsertOrder(number, total, order_by, sent_to, note string) bool {
	_, err := ExecuteQuery("INSERT INTO t_order(number, total, order_by, sent_to, note) VALUES ($1,$2,$3,$4,$5);", number, total, order_by, sent_to, note) //returning id
	if err == nil {
		return true
	} else {
		return false
	}
}
