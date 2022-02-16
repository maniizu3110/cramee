package models

import "time"

//go:generate go run ../codegen/main.go -file ${GOFILE} -dest ..
type Student struct {
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
	StripeID                string                   `json:"stripe_id"`
	PasswordChangedAt       time.Time                `json:"password_changed_at"`
	StudentLectureSchedules []StudentLectureSchedule `json:"student_lecture_schedules"`
}

type LimitedStudentInfo struct {
	ID          uint   `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

//必要最低限の情報のみ抽出して返す
func (s *Student) GetLimitedInfo() *LimitedStudentInfo {
	return &LimitedStudentInfo{
		ID:          s.ID,
		PhoneNumber: s.PhoneNumber,
		Email:       s.Email,
	}
}

func (m *Student) SetPasswordChangedAt(t time.Time) {
	m.PasswordChangedAt = t.UTC()
}
