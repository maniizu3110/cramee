package services

import (
	//"cramee/api/services/types"
	"cramee/models"
	"cramee/myerror"
	"cramee/token"
	"cramee/util"
)

type SignStudentService interface {
	CreateStudent(params *models.Student) (*models.LimitedStudentInfo, error)
	//LoginStudent(params *types.LoginStudentRequest) (*types.LoginStudentResponse, error)
}

type signStudentServiceImpl struct {
	repo       StudentRepository
	config     util.Config
	tokenMaker token.Maker
}

func NewSignStudentService(repository StudentRepository, config util.Config, tokenMaker token.Maker) SignStudentService {
	res := &signStudentServiceImpl{}
	res.repo = repository
	res.config = config
	res.tokenMaker = tokenMaker
	return res
}

func (s *signStudentServiceImpl) CreateStudent(params *models.Student) (*models.LimitedStudentInfo, error) {
	hashedPassword, err := util.HashPassword(params.HashedPassword)
	if err != nil {
		return nil, myerror.NewPublic(myerror.ErrChangePasswordToHash, err)
	}
	params.HashedPassword = hashedPassword
	student, err := s.repo.Create(params)
	if err != nil {
		return nil, myerror.NewPublic(myerror.ErrCreate, err)
	}
	return student.GetLimitedInfo(), nil
}

//func (s *signStudentServiceImpl) LoginStudent(params *LoginStudentParams) (*LoginStudentResponse, error) {
//	student, err := s.store.GetStudentByEmail(context.Background(), params.Email)
//	if err != nil {
//		return nil, myerror.NewPublic(myerror.ErrGet, err)
//	}
//	err = util.CheckPassword(params.Password, student.HashedPassword)
//	if err != nil {
//		return nil, myerror.NewPublic(myerror.ErrCheckPassword, err)
//	}
//	accessToken, err := s.tokenMaker.CreateToken(student.ID, false, s.config.AccessTokenDuration)
//	if err != nil {
//		return nil, myerror.NewPublic(myerror.ErrCheckPassword, err)
//	}
//	createdStudent := newStudentResponse(student)
//	res := &LoginStudentResponse{
//		AccessToken: accessToken,
//		Client:      *createdStudent,
//	}
//	return res, nil
//}
