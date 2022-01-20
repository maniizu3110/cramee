package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

//Different types of error returned by the VerifiToken function
var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

//Payload contains the payload data of the token
type Payload struct {
	UUID      uuid.UUID
	ID        int64
	IsTeacher bool
	Authority int64
	IssuedAt  time.Time
	ExpiredAt time.Time
}
//go:generate mockgen -source=$GOFILE -destination=${GOPACKAGE}_mock/${GOFILE}.mock.go -package=${GOPACKAGE}_mock

//NewPayload creates a token payload with a specific username and duration
func NewPayload(id int64, isTeacher bool, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		UUID:      tokenID,
		ID:        id,
		IsTeacher: isTeacher,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

//Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
