package models

import "time"

type TeacherLectureSchedule struct {
	Model
	TeacherID  int64     `json:"teacher_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	IsReserved bool      `json:"is_reserved"`
}
