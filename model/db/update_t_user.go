package db

import (
	"fmt"
)

func UpdateUserDetail(id, name, phone, address, company, email, password, role string) error {
	_, err := ExecuteQuery("UPDATE t_user SET name=$1, phone=$2, address=$3, company=$4 WHERE id=$5;", name, phone, address, company, id)
	if err != nil {
		fmt.Println("Err : UpdateUser - ", err)
	}
	return err
}

func UpdateUserSession(session_token, email, password string) error {
	_, err := ExecuteQuery("UPDATE t_user SET session_token=$1 WHERE email=$2 AND password=$3;", session_token, email, password)
	if err != nil {
		fmt.Println("Err : UpdateUserSession - ", err)
	}
	return err
}
