package model

import (
	"bytes"
	"chatgpt/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Ask_req struct {
	User_id    string `json:"user_id"`
	Session_id int
	Question   string
}

type Ask_res struct {
	Ret  int
	Meg  string
	Data []struct {
		Answer string
	}
}
type Session struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Json struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index        int `json:"index"`
		Message      Session
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func fetchSessionData(session_id int) ([]Session, error) {

	rows, err := DB.Query("SELECT question, answer FROM t_session_detail WHERE session_id = ?", session_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessionData []Session
	for rows.Next() {
		var question, answer string
		err := rows.Scan(&question, &answer)
		if err != nil {
			return nil, err
		}
		sessionData = append(sessionData, Session{Role: "user", Content: question})
		sessionData = append(sessionData, Session{Role: "assistant", Content: answer})
	}

	return sessionData, nil
}

func Ask(requestData *Ask_req) (string, error) {

	user_id := requestData.User_id
	session_id := requestData.Session_id
	question := requestData.Question

	sessionData, err := fetchSessionData(session_id)
	if err != nil {
		return "", err
	}
	sessionData = append(sessionData, Session{Role: "user", Content: question})

	var message []Session
	message = append(message, Session{Role: "system", Content: "You are a helpful assistant."}) //这两行的顺序是否颠倒了？？？
	message = append(message, sessionData...)
	JSON_str, err := json.Marshal(message)
	if err != nil {
		return "", err
	}

	answer, err := api(string(JSON_str))
	if err != nil {
		return "", err
	}

	msgSize := len(question) + len(answer)

	_, err = DB.Exec("INSERT INTO t_session_detail (session_id, user_id, question, answer, msgSize) VALUES (?, ?, ?, ?, ?)",
		session_id, user_id, question, answer, msgSize)
	if err != nil {
		fmt.Println("Failed to insert data:", err)
		return "", err
	}

	return answer, nil

}
func api(message string) (string, error) {
	postInfo := config.GetOpenaiInfo()
	url := postInfo.Openai.Url
	fmt.Println("testConfig", postInfo.Openai.Authorization)
	var messages []Session
	if message != "" {
		err := json.Unmarshal([]byte(message), &messages)
		if err != nil {
			return "", err
		}
	}
	headers := make(map[string]string)
	headers["Content-Type"] = postInfo.Openai.ConetType
	headers["Authorization"] = postInfo.Openai.Authorization
	body := map[string]interface{}{
		"model":    postInfo.Openai.Model,
		"messages": messages,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("errtest1:", err)
		return "", err
	}
	defer resp.Body.Close()

	var result Json
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	answer := result.Choices[0].Message.Content
	return answer, err
}

