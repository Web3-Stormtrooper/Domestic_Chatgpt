package model

type Create_req struct {
	User_id string `json:"user_id"`
	Title   string
}

type Create_res struct {
	Ret  int
	Meg  string
	Data []struct {
		Session_id int64
	}
}

func Create(requestData *Create_req) (int64, error) {

	user_id := requestData.User_id
	title := requestData.Title

	result, err := DB.Exec("INSERT INTO t_session_info (user_id, title) VALUES (?, ?)", user_id, title)
	if err != nil {
		return 0, err
	}

	session_id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return session_id, nil

}
