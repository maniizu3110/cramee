package models

import "time"

//go:generate go run ../codegen/main.go -file ${GOFILE} -dest ..
type TeacherLectureSchedule struct {
	Model
	TeacherID  int64     `json:"teacher_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	IsReserved bool      `json:"is_reserved"`
	Teacher    Teacher   `json:"teacher"`
}
