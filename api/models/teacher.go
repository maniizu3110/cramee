package models

import "time"

//go:generate go run ../codegen/main.go -file ${GOFILE} -dest ..
type Teacher struct {
	Model
	FirstName               string                   `json:"first_name"`
	FirstNameKana           string                   `json:"first_name_kana"`
	LastName                string                   `json:"last_name"`
	LastNameKana            string                   `json:"last_name_kana"`
	PhoneNumber             string                   `json:"phone_number" gorm:"unique"`
	Email                   string                   `json:"email" gorm:"unique"`
	Address                 string                   `json:"address"`
	HashedPassword          string                   `json:"hashed_password" gorm:"unique"`
	Image                   string                   `json:"image"`
	PasswordChangedAt       time.Time                `json:"password_changed_at"`
	TeacherLectureSchedules []TeacherLectureSchedule `json:"teacher_lecture_schedules"`
}
