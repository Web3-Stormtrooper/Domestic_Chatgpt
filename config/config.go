package config

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

/*
从配置文件解析出配置信息
*/
type DbInfo struct {
	Mysql struct {
		User     string `yaml: "user"`
		Port     string `yaml: "port"`
		Password string `yaml: "password"`
		Host     string `yaml: "host"`
		Dbname   string `yaml: "dbname"`
	}
}
type ApiOpenaiInfo struct {
	Openai struct {
		Url           string `yaml: "url"`
		ConetType     string `yaml: "content-type"`
		Authorization string `yaml: "authorization"`
		Model         string `yaml: "model"`
	}
}

func GetDbInfo() *DbInfo {
	var dbInfo DbInfo
	File, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Printf("读取配置文件失败 #%v", err)
	}
	err = yaml.Unmarshal(File, &dbInfo)
	if err != nil {
		log.Fatalf("解析失败: %v", err)
	}
	return &dbInfo
}

func GetOpenaiInfo() *ApiOpenaiInfo {
	var openaiInfo ApiOpenaiInfo
	File, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Printf("读取配置文件失败 #%v", err)
	}
	err = yaml.Unmarshal(File, &openaiInfo)
	if err != nil {
		log.Fatalf("解析失败: %v", err)
	}
	return &openaiInfo
}
