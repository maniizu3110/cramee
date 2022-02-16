package models

import "time"

//go:generate go run ../codegen/main.go -file ${GOFILE} -dest ..
type Lecture struct {
	Model
	TeacherID uint      `json:"teacher_id" gorm:"default:0"`
	StudentID uint      `json:"student_id" gorm:"default:0"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status" gorm:"type: enum('empty','pending','reserved','finish','absent'); default:'empty'"`
	ZoomLink  string    `json:"zoom_link"`
}
