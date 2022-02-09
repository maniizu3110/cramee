package services

import (
	"cramee/myerror"
	"cramee/token"
	"cramee/util"
)

type SignStudentService interface {
	CreateStudent(params *repository.CreateStudentParams) (*repository.Student, error)
	LoginStudent(params *LoginStudentParams) (*LoginStudentResponse, error)
}

type SignStudentRepostor interface {
	CreateStudent()
}

type signStudentServiceImpl struct {
	repo StudentRepository
	config     util.Config
	tokenMaker token.Maker
}


func NewSignStudentService(config util.Config, tokenMaker token.Maker) SignStudentService {
	res := &signStudentServiceImpl{}
	res.repo = repository
	res.config = config
	res.tokenMaker = tokenMaker
	return res
}


func (s *signStudentServiceImpl) CreateStudent(params *repository.CreateStudentParams) (*repository.Student, error) {
	hashedPassword, err := util.HashPassword(params.HashedPassword)
	if err != nil {
		return nil, myerror.NewPublic(myerror.ErrChangePasswordToHash, err)
	}
	params.HashedPassword = hashedPassword
	student, err := s.repository.CreateStudent(params)
	if err != nil {
		return nil, myerror.NewPublic(myerror.ErrCreate, err)
	}
	return &student, nil
}