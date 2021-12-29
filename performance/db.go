package performance

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
)

func GetDb(name string, configPath string) *sql.DB {
	config := loadConfig(configPath)
	db, err := sql.Open("mysql", config.Username+":"+config.Password+"@tcp("+config.Hostname+":"+config.Port+")/"+name+"?charset=utf8")

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

//用json配置测试
type Config struct {
	Username string `json:"username:`
	Password string `json:"password:`
	Hostname string `json:"hostname:`
	Port     string `json:"port:`
}

func loadConfig(configPath string) *Config {
	//不同的配置规则，解析复杂度不同
	f, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println("load config error: ", err)
		return nil
	}
	config := Config{}
	err = json.Unmarshal(f, &config)
	if err != nil {
		fmt.Println("Para config failed: ", err)
		return &config
	}
	return &config
}
