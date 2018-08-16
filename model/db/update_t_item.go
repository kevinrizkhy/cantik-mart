package db

import ()

func UpdateItem(id, name, category, buy, sell, margin, description string) bool {
	_, err := ExecuteQuery("UPDATE t_item SET name=$1, category=$2, buy=$3, sell=$4, margin=$5, description=$6 WHERE id=$7;", name, category, buy, sell, margin, description, id) //returning id
	if err == nil {
		return true
	} else {
		return false
	}
}
