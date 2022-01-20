package main

import (
	"cramee/util"
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	///mysqlに接続する
	config, err := util.LoadConfig(".")
	if err != nil {
		logrus.Panic(err)
	}
	db := util.InitDatabase(config)

	//crameeデータベースから全てのテーブルを削除
	var tables []string
	{
		rows, err := db.Raw("select table_name from information_schema.tables where table_schema = ?", "cramee").Rows()
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var name string
			rows.Scan(&name)
			tables = append(tables, name)
		}
	}
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Info(err)
	}
	if _, err := sqlDB.Exec("SET FOREIGN_KEY_CHECKS=0"); err != nil {
		panic(err)
	}
	for _, name := range tables {
		fmt.Println(name)
		if _, err := sqlDB.Exec("DROP TABLE " + name); err != nil {
			panic(err)
		}
	}

}
