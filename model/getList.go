package model

import (
	"log"
)

type List_req struct {
	User_id string `json:"user_id"`
}

type List1 struct {
	Title           string
	Session_id      int
	CreateTimestamp string
	UpdateTimestamp string
}

type List_res struct {
	Ret  int
	Meg  string
	Data []List1
}

func GetList(user_id string) ([]List1, error) {
    rows, err := DB.Query("SELECT session_id, title, create_timestamp, update_timestamp FROM t_session_info where user_id=? AND  session_status=1", user_id)
    if err != nil {
        log.Printf("Error executing query: %v", err)
        return nil, err
    }
    defer rows.Close()

	var data []List1
	for rows.Next() {
		var item List1
		err := rows.Scan(&item.Session_id, &item.Title, &item.CreateTimestamp, &item.UpdateTimestamp)
		if err != nil {
			log.Printf("Error2: %v", err)
			return nil, err
		}
		data = append(data, item)
	}

	return data, nil
}
