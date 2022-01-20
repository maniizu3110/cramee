package models

//go:generate go run ../codegen/main.go -file ${GOFILE} -dest ..
type Lecture struct {
	Model
	TeacherLectureScheduleID int64 `json:"teacher_lecture_schedule_id"`
	StudentLectureScheduleID int64 `json:"student_lecture_schedule_id"`
}

