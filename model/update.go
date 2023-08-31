package model

type Json_req struct {
	User_id    string `json:"user_id"`
	Title      string
	Session_id int
}

type Json_res struct {
	Ret  int
	Meg  string
	Data []struct {
		Title string
	}
}

func Update(session_id int, title string) error {

	_, err := DB.Exec("UPDATE t_session_info SET title = ? WHERE session_id = ?", title, session_id)
	if err != nil {
		return err
	}

	return nil
}
