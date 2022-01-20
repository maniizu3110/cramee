package main

import (
	"cramee/models"
	"cramee/util"

	"github.com/sirupsen/logrus"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		logrus.Panic(err)
	}
	db := util.InitDatabase(config)

	db.AutoMigrate(
		&models.Student{},
		&models.Lecture{},
		&models.StudentLectureSchedule{},
		&models.Teacher{},
		&models.TeacherLectureSchedule{},
	)
}
