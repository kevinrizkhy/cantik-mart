package user

import (
	"fmt"
	database "github.com/pardev/cantik-mart/model/db"
)

func GetMemberMap() map[int](map[string]string) {
	rows, err := database.ExecuteQuery("select * from t_member ORDER BY name ASC")
	if len(rows) > 0 && err == nil {
		return rows
	} else {
		if err != nil {
			fmt.Println("ERR : GetMemberMap - ", err)
		}
		return nil
	}
}

func GetMemberList() [][]string {
	rows := GetMemberMap()
	var result [][]string
	if rows == nil {
		return result
	} else {
		for i := 0; i < len(rows); i++ {
			id := fmt.Sprint(rows[i]["id"])
			name := fmt.Sprint(rows[i]["name"])
			address := fmt.Sprint(rows[i]["address"])
			phone := fmt.Sprint(rows[i]["phone"])
			sex := fmt.Sprint(rows[i]["sex"])
			point := fmt.Sprint(rows[i]["point"])
			birthday := fmt.Sprint(rows[i]["birthday"])
			card := fmt.Sprint(rows[i]["card"])
			ktp := fmt.Sprint(rows[i]["ktp"])
			button := fmt.Sprint("<button onclick=\"location.href = '/detail/member?id=" + id + "';\" class=\"btn btn-info\">Detail</button>&nbsp&nbsp&nbsp<button onclick=\"editMember(this)\" class=\"btn btn-warning\">Edit</button>")
			temp := []string{id, card, ktp, name, sex, address, phone, birthday, point, button}
			result = append(result, temp)
		}
		return result
	}
}
