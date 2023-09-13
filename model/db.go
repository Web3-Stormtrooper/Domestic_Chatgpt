package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"chatgpt/config"
	"strings"
	"io/ioutil"
	"log"
)

/*
完成数据库启动以及一些数据库相关的操作
*/
var DB *sql.DB

func ConnectDB() {
    var err error
    
    Mysqlinfo := config.GetDbInfo()
    dbURL := Mysqlinfo.Mysql.User + ":" + Mysqlinfo.Mysql.Password + "@tcp(" + Mysqlinfo.Mysql.Host + ":" + Mysqlinfo.Mysql.Port + ")/" + Mysqlinfo.Mysql.Dbname
    DB, err = sql.Open("mysql", dbURL)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("connect success")

    // 读取 SQL 脚本内容
	sqlScript, err := ioutil.ReadFile("database/db.sql")
	if err != nil {
		log.Fatal(err)
	}

    // 执行 SQL 脚本

	statements := strings.Split(string(sqlScript), ";")

	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt != "" {
			_, err := DB.Exec(stmt)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

    fmt.Println("SQL script executed successfully.")
}
