package db

import ()

func UpdateOrder(id, number, total, order_by, sent_to, note, status string) bool {
	_, err := ExecuteQuery("UPDATE t_order SET number=$1, total=$2, order_by=$3, sent_to=$4, note=$5, status=$6 WHERE id=$7;", number, total, order_by, sent_to, note, status, id) //returning id
	if err == nil {
		return true
	} else {
		return false
	}
}
