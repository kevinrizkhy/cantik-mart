package db

import ()

func DeleteItem(id, name, category string) bool {
	_, err := ExecuteQuery("DELETE FROM t_item WHERE id=$1 AND name=$2 AND category=$3;", id, name, category) //returning id
	if err == nil {
		return true
	} else {
		return false
	}
}
