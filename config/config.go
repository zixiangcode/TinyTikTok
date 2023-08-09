package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	Mysql string `json:"Mysql"`
}

var Config Configuration // 实例化一个 Configuration 类对象

// ReadConfig 读取配置文件
func ReadConfig(fileName string) {
	configFile, err := ioutil.ReadFile(fileName) // 读取文件内容并将其作为字节切片返回
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	err = json.Unmarshal(configFile, &Config) // 将解析后的 JSON 式数据存储在 Config 对象中
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
}
