package db

import ()

func InsertItem(id, name, category, buy, sell, margin, description string) bool {
	_, err := ExecuteQuery("INSERT INTO t_item (id, name, category, buy, sell, margin, description) VALUES ($1,$2,$3,$4,$5,$6,$7);", id, name, category, buy, sell, margin, description) //returning id
	if err == nil {
		return true
	} else {
		return false
	}
}
