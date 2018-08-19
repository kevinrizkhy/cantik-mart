package db

import (
	"fmt"
)

func DeleteSupplier(id string) error {
	_, err := ExecuteQuery("DELETE FROM t_supplier WHERE id=$1;", id)
	if err != nil {
		fmt.Println("Err : DeleteSupplier - ", err)
		return err
	} else {
		return err
	}
}
