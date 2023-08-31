package api

import (
	"chatgpt/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 提问
func Ask(c *gin.Context) {

	var requestData model.Ask_req
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	answer, err := model.Ask(&requestData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	res := model.Ask_res{
		Ret: 0,
		Meg: "success",
		Data: []struct{ Answer string }{
			{Answer: answer}},
	}
	c.JSON(http.StatusOK, res)

}

// 创建会话
func Create(c *gin.Context) {
	var requestData *model.Create_req
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session_id, err := model.Create(requestData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// 构造并返回响应数据
	res := model.Create_res{
		Ret: 0,
		Meg: "success",
		Data: []struct{ Session_id int64 }{
			{Session_id: session_id}},
	}
	c.JSON(http.StatusOK, res)
}

func Delete(c *gin.Context) {
	var requestData model.Delete_req
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = model.Delete(&requestData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	res := model.Delete_res{
		Ret: 0,
		Meg: "success",
		Data: []struct {
			Session_id int
			User_id    string
		}{{Session_id: requestData.Session_id, User_id: requestData.User_id}},
	}
	c.JSON(http.StatusOK, res)
}

func Detail(c *gin.Context) {
	var requestData model.Detail_req
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_id := requestData.User_id
	session_id := requestData.Session_id

	lists, err := model.Detail(session_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch session detail"})
		return
	}

	// 构造并返回响应数据
	res := model.Detail_res{
		Ret: 0,
		Meg: "success",
		Data: []struct {
			User_id    string
			Session_id int64
			List       []model.Detail_list
		}{
			{User_id: user_id, Session_id: session_id, List: lists},
		},
	}
	c.JSON(http.StatusOK, res)
}
func GetList(c *gin.Context) {
	var requestData model.List_req
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_id := requestData.User_id

	data, err := model.GetList(user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	res := model.List_res{
		Ret:  0,
		Meg:  "success",
		Data: data,
	}
	c.JSON(http.StatusOK, res)
}

func Update(c *gin.Context) {
	var requestData model.Json_req
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session_id := requestData.Session_id
	title := requestData.Title

	err = model.Update(session_id, title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update session title"})
		return
	}

	res := model.Json_res{
		Ret: 0,
		Meg: "success",
		Data: []struct {
			Title string
		}{{Title: title}},
	}
	c.JSON(http.StatusOK, res)
}
