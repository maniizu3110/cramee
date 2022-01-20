package token

import "time"

//go:generate mockgen -source=$GOFILE -destination=${GOPACKAGE}_mock/${GOFILE}.mock.go -package=${GOPACKAGE}_mock
type Maker interface{
	CreateToken(id int64, isTeacher bool,duration time.Duration)(string,error)
	VerifyToken(token string)(*Payload,error)
}