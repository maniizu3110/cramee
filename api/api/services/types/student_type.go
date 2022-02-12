package types

import "cramee/models"

type CreateStudent struct {
}

type LoginStudentRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginStudentResponse struct {
	AccessToken string                     `json:"access_token"`
	Student     *models.LimitedStudentInfo `json:"student"`
}

type ChargeWithIDParams struct {
	Amount    int64 `json:"amount"`
	StudentID uint  `json:"student_id"`
}
