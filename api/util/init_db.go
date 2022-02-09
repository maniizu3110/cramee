package util

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase(config Config) (db *gorm.DB) {
	logrus.Info("接続情報出力")
	logrus.Info(config.DBSource)
	db, err := gorm.Open(mysql.Open(config.DBSource), &gorm.Config{})
	if err != nil {
		logrus.Panic(err)
	}
	return db
}
