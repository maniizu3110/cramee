package models

import "time"

//go:generate go run ../codegen/main.go -file ${GOFILE} -dest ..

type StudentLectureSchedule struct {
	Model
	StudentID int64     `json:"student_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status" gorm:"type: enum('empty','pending','reserved','finish','absent'); default:'empty'"`
	Student   Student   `json:"student"`
}
