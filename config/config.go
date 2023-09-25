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

var configFileContent []byte

func init() {
	var err error
	configFileContent, err = ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("读取配置文件失败 #%v", err)
	}
}

func GetDbInfo() *DbInfo {
	var dbInfo DbInfo
	err := yaml.Unmarshal(configFileContent, &dbInfo)
	if err != nil {
		log.Fatalf("解析失败: %v", err)
	}
	return &dbInfo
}

func GetOpenaiInfo() *ApiOpenaiInfo {
	var openaiInfo ApiOpenaiInfo
	err := yaml.Unmarshal(configFileContent, &openaiInfo)
	if err != nil {
		log.Fatalf("解析失败: %v", err)
	}
	return &openaiInfo
}

