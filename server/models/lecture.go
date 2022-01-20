package models

import "time"

type Lecture struct {
	TeacherLectureScheduleID int64     `json:"teacher_lecture_schedule_id"`
	StudentLectureScheduleID int64     `json:"student_lecture_schedule_id"`
	TeacherID                int64     `json:"teacher_id"`
	StudentID                int64     `json:"student_id"`
	StartTime                time.Time `json:"start_time"`
	EndTime                  time.Time `json:"end_time"`
}
