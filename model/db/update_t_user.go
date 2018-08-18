package db

import (
	"fmt"
)

func UpdateUserDetail(id, name, email, address, phone, role, store, status string) bool {
	_, err := ExecuteQuery("UPDATE t_user SET email=$1, name=$2, address=$3, phone=$4,role=$5,store=$6,status=$7 WHERE id=$8;", name, email, address, phone, role, store, status, id)
	if err != nil {
		fmt.Println("Err : UpdateUser - ", err)
		return false
	}
	return true
}

func UpdateUserSession(session_token, email, password string) error {
	_, err := ExecuteQuery("UPDATE t_user SET session_token=$1 WHERE email=$2 AND password=$3;", session_token, email, password)
	if err != nil {
		fmt.Println("Err : UpdateUserSession - ", err)
	}
	return err
}
