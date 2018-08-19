package db

import ()

func UpdateOrder(id, name, supplier, note, total string) bool {
	_, err := ExecuteQuery("UPDATE t_store SET name=$1, supplier=$2, total=$3, note=$4, WHERE id=$5;", name, supplier, total, note, id) //returning id
	if err == nil {
		return true
	} else {
		return false
	}
}
