package services

import (
	"cramee/api/services/types"
	"cramee/models"
	"cramee/myerror"
	"cramee/token"
	"cramee/util"
)

type SignTeacherService interface {
	CreateTeacher(params *models.Teacher) (*models.LimitedTeacherInfo, error)
	LoginTeacher(params *types.LoginTeacherRequest) (*types.LoginTeacherResponse, error)
}

type signTeacherServiceImpl struct {
	repo       TeacherRepository
	config     util.Config
	tokenMaker token.Maker
}

func NewSignTeacherService(repository TeacherRepository, config util.Config, tokenMaker token.Maker) SignTeacherService {
	res := &signTeacherServiceImpl{}
	res.repo = repository
	res.config = config
	res.tokenMaker = tokenMaker
	return res
}

func (s *signTeacherServiceImpl) CreateTeacher(params *models.Teacher) (*models.LimitedTeacherInfo, error) {
	hashedPassword, err := util.HashPassword(params.HashedPassword)
	if err != nil {
		return nil, myerror.NewPublic(myerror.ErrChangePasswordToHash, err)
	}
	params.HashedPassword = hashedPassword
	teacher, err := s.repo.Create(params)
	if err != nil {
		return nil, myerror.NewPublic(myerror.ErrCreate, err)
	}
	return teacher.GetLimitedInfo(), nil
}

func (s *signTeacherServiceImpl) LoginTeacher(params *types.LoginTeacherRequest) (*types.LoginTeacherResponse, error) {
	teacher, err := s.repo.GetByEmail(params.Email)
	if err != nil {
		return nil, myerror.NewPublic(myerror.ErrGet, err)
	}
	err = util.CheckPassword(params.Password, teacher.HashedPassword)
	if err != nil {
		return nil, myerror.NewPublic(myerror.ErrCheckPassword, err)
	}
	accessToken, err := s.tokenMaker.CreateToken(int64(teacher.ID), false, s.config.AccessTokenDuration)
	if err != nil {
		return nil, myerror.NewPublic(myerror.ErrCheckPassword, err)
	}
	res := &types.LoginTeacherResponse{
		AccessToken: accessToken,
		Teacher:     teacher.GetLimitedInfo(),
	}
	return res, nil
}
