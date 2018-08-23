package db

import ()

func UpdateOrder(id, number, order_by, order_to, total, note, status string) bool {
	_, err := ExecuteQuery("UPDATE t_order SET number=$1, order_by=$2, order_to=$3, total=$4 note=$5, status=$6 WHERE id=$7;", number, order_by, order_to, total, note, status, id) //returning id
	if err == nil {
		return true
	} else {
		return false
	}
}
