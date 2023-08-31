package model

type Delete_req struct {
	User_id    string `json:"user_id"`
	Session_id int
}
type Delete_res struct {
	Ret  int
	Meg  string
	Data []struct {
		Session_id int
		User_id    string
	}
}

func Delete(requestData *Delete_req) error {

	//user_id := requestData.User_id
	session_id := requestData.Session_id

	session_status := 2
	_, err := DB.Exec("UPDATE t_session_info SET session_status = ? WHERE session_id = ?", session_status, session_id)
	if err != nil {
		return err
	}

	return nil

}
