package config

import (
	"database/sql"
)

<<<<<<< Updated upstream
var Db *sql.DB

=======
type Configuration struct {
	Mysql string `json:"Mysql"`
}
>>>>>>> Stashed changes
const AccessKeyId = ""
const AccessKeySecret = ""
const  Endpoint = "https://oss-cn-hangzhou.aliyuncs.com" // OSS的访问域名   杭州

const BucketName = "web-tlias-amireux"
<<<<<<< Updated upstream

=======
>>>>>>> Stashed changes


<<<<<<< Updated upstream
=======
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


>>>>>>> Stashed changes
