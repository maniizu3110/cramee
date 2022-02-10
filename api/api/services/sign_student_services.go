package services

import (
	"cramee/api/services/types"
	"cramee/models"
	"cramee/myerror"
	"cramee/token"
	"cramee/util"

	"github.com/stripe/stripe-go/v72"
)

type SignStudentService interface {
	CreateStudent(params *models.Student) (*models.LimitedStudentInfo, error)
	LoginStudent(params *types.LoginStudentRequest) (*types.LoginStudentResponse, error)
}

type signStudentServiceImpl struct {
	repo          StudentRepository
	config        util.Config
	tokenMaker    token.Maker
	stripeService StripeService
}

func NewSignStudentService(repository StudentRepository, config util.Config, tokenMaker token.Maker, stripeService StripeService) SignStudentService {
	res := &signStudentServiceImpl{}
	res.repo = repository
	res.config = config
	res.tokenMaker = tokenMaker
	res.stripeService = stripeService
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
	client, err := s.stripeService.CreateCustomer(&stripe.CustomerParams{
		//TODO:登録する情報を増やす
		Email: &student.Email,
		Phone: &student.PhoneNumber,
	})
	if err != nil {
		return nil, err
	}
	student.StripeID = client.ID
	student, err = s.repo.Update(student.ID, student)
	if err != nil {
		return nil, err
	}
	return student.GetLimitedInfo(), nil
}

func (s *signStudentServiceImpl) LoginStudent(params *types.LoginStudentRequest) (*types.LoginStudentResponse, error) {
	student, err := s.repo.GetByEmail(params.Email)
	if err != nil {
		return nil, myerror.NewPublic(myerror.ErrGet, err)
	}
	err = util.CheckPassword(params.Password, student.HashedPassword)
	if err != nil {
		return nil, myerror.NewPublic(myerror.ErrCheckPassword, err)
	}
	accessToken, err := s.tokenMaker.CreateToken(int64(student.ID), false, s.config.AccessTokenDuration)
	if err != nil {
		return nil, myerror.NewPublic(myerror.ErrCheckPassword, err)
	}
	res := &types.LoginStudentResponse{
		AccessToken: accessToken,
		Student:     student.GetLimitedInfo(),
	}
	return res, nil
}
