package conf

import (
	"encoding/json"
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

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ReadConfig(fileName string, Cfg interface{}) error {
	filePath := "conf/" + fileName + ".json"
	fileExists, err := FileExists(filePath)
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
