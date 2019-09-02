package conf

import (
	"encoding/json"
	"kubewatch/util"
	"log"
	"os"
)

type CoreConfig struct {
	MysqlDbHost   string `json:"mysqlDbHost"`
	MysqlDbPort   string `json:"mysqlDbPort"`
	MysqlDbName   string `json:"mysqlDbName"`
	MysqlUserName string `json:"mysqlUserName"`
	MysqlDbPass   string `json:"mysqlDbPass"`
}

func ReadConfig(fileName string, Cfg interface{}) error {
	filePath := "conf/" + fileName + ".json"
	fileExists, err := util.FileExists(filePath)
	if !fileExists {
		return err
	}
	file, _ := os.Open(filePath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err != nil {
		log.Fatal("Read config error: ", err)
	}
	err = decoder.Decode(Cfg)
	if err != nil {
		log.Fatal("Json decode error: ", err)
	}
	return nil
}

func GetCoreConfig() *CoreConfig {
	clientPointer := new(CoreConfig)
	err := ReadConfig("config", clientPointer)
	if err != nil {
		log.Fatal("Read Config error: ", err)
	}
	return clientPointer
}
