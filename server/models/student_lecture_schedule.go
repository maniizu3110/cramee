package models

import "time"

type StudentLectureSchedule struct {
	Model
	StudentID  int64     `json:"student_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	IsReserved bool      `json:"is_reserved"`
}
