package dao

import (
	"time"
	"vermouth/conf"
	"vermouth/model"
	"vermouth/util"
)

func Init() {
	// Init Mysql and Migrate the schema
	config := conf.GetCoreConfig()
	mysqlClient := util.GetConnect(config.MysqlUserName, config.MysqlDbPass, config.MysqlDbHost, config.MysqlDbPort, config.MysqlDbName)
	mysqlClient.Db.DB().SetMaxOpenConns(config.MysqlMaxOpenConns)
	mysqlClient.Db.DB().SetConnMaxLifetime(10 * time.Minute)
	mysqlClient.Db.Exec("DROP TABLE user")
	mysqlClient.Db.Exec("DROP TABLE pay_record")
	mysqlClient.Db.AutoMigrate(&model.User{}, &model.PayRecord{})

	//user := model.User{
	//	Nickname:   "轨迹",
	//	SessionKey: "C6+JtMIQwUltC3ezmdLDkg==",
	//	OpenID:     "oVRY_5YntTT7wIIUsNwjam7LRBlM",
	//	Avatar:     "https://wx.qlogo.cn/mmopen/vi_32/PiajxSqBRaEIaZpTg3By2IwOm4SKOaCdbibcMuiaRxEPV3SEibDs1uqr32OKlETvyDbh5RBrkFvVTCvzg1hg06agUQ/132",
	//	Province:   "湖北",
	//	Country:    "中国",
	//	Gender:     1,
	//}
	//mysqlClient.Db.Delete(&user)
}
