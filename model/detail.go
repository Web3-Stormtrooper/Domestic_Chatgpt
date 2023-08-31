package model

import (
	"log"
)

type Detail_req struct {
	User_id    string `json:"user_id"`
	Session_id int64
}

type Detail_list struct {
	Question string
	Answer   string
}

type Detail_res struct {
	Ret  int
	Meg  string
	Data []struct {
		User_id    string
		Session_id int64
		List       []Detail_list
	}
}

func Detail(session_id int64) ([]Detail_list, error) {
	rows, err := DB.Query("SELECT question, answer FROM t_session_detail WHERE session_id = ? ", session_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var lists []Detail_list
	for rows.Next() {
		var item Detail_list
		err := rows.Scan(&item.Question, &item.Answer)
		if err != nil {
			log.Fatal(err)
		}
		lists = append(lists, item)
	}

	return lists, nil
}
