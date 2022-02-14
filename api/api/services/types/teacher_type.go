package types

import "cramee/models"

type CreateTeacher struct {
}

type LoginTeacherRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginTeacherResponse struct {
	AccessToken string                     `json:"access_token"`
	Teacher     *models.LimitedTeacherInfo `json:"teacher"`
}
