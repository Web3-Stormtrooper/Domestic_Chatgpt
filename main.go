package main

import (
	"chatgpt/model"
	"chatgpt/routes"
)

func main() {
	///连接数据库
	model.ConnectDB()
	//初始化路由
	routes.InitRouter()
}
