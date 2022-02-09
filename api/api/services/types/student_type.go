package types


type CreateStudent struct{

}

type LoginStudentRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginStudentResponse struct {
	AccessToken string          `json:"access_token"`
	Client      StudentResponse `json:"student"`
}

type StudentResponse struct {
	ID          int64  `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}