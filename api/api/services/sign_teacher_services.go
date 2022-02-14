package services

import (
	"cramee/api/services/types"
	"cramee/lib/zoom"
	"cramee/models"
	"cramee/myerror"
	"cramee/token"
	"cramee/util"
)

type SignTeacherService interface {
	CreateTeacher(params *models.Teacher) (*models.LimitedTeacherInfo, error)
	CreateTeacherWithZoom(params *models.Teacher) (*models.LimitedTeacherInfo, error)
	LoginTeacher(params *types.LoginTeacherRequest) (*types.LoginTeacherResponse, error)
}

type signTeacherServiceImpl struct {
	repo        TeacherRepository
	config      util.Config
	tokenMaker  token.Maker
	zoomService ZoomService
}

func NewSignTeacherService(repository TeacherRepository, config util.Config, tokenMaker token.Maker, zoomService ZoomService) SignTeacherService {
	res := &signTeacherServiceImpl{}
	res.repo = repository
	res.config = config
	res.tokenMaker = tokenMaker
	res.zoomService = zoomService
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

func (s *signTeacherServiceImpl) CreateTeacherWithZoom(params *models.Teacher) (*models.LimitedTeacherInfo, error) {
	hashedPassword, err := util.HashPassword(params.HashedPassword)
	if err != nil {
		return nil, myerror.NewPublic(myerror.ErrChangePasswordToHash, err)
	}
	rawPassword := params.HashedPassword
	params.HashedPassword = hashedPassword
	teacher, err := s.repo.Create(params)
	if err != nil {
		return nil, myerror.NewPublic(myerror.ErrCreate, err)
	}
	params.HashedPassword = rawPassword
	userInfo := zoom.CreateUserInfo{
		//TODO:validationをzoomAPIの仕様と合わせる
		Email:     params.Email,
		Type:      zoom.Basic,
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Password:  params.HashedPassword,
	}
	opt := zoom.CreateUserOptions{
		Action:   zoom.AutoCreate,
		UserInfo: userInfo,
	}
	_, err = s.zoomService.CreateUser(opt)
	if err != nil {
		return nil, err
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
