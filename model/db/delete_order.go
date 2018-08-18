package db

import ()

func DeleteOrder(id, name, supplier, total, note string) bool {
	_, err := ExecuteQuery("DELETE FROM t_order WHERE id=$1 AND name=$2 AND supplier=$3 AND total=$4 AND note =$5;", id, name, supplier, total, note) //returning id
	if err == nil {
		return true
	} else {
		return false
	}
}
