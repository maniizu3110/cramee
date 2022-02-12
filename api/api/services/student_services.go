package services

import (
	"cramee/api/services/types"
	"cramee/models"

	"github.com/stripe/stripe-go/v72"
)

//go:generate $GOPATH/bin/mockgen -source=$GOFILE -destination=${GOPACKAGE}_mock/${GOFILE}.mock.go -package=${GOPACKAGE}_mock

type StudentRepository interface {
	GetByID(id uint, expand ...string) (*models.Student, error)
	GetAll(config GetAllConfig) (data []*models.Student, count uint, err error)
	Create(data *models.Student) (*models.Student, error)
	Update(id uint, data *models.Student) (*models.Student, error)
	SoftDelete(id uint) (*models.Student, error)
	HardDelete(id uint) (*models.Student, error)
	Restore(id uint) (*models.Student, error)
	GetByEmail(email string) (*models.Student, error)
}

type StudentService interface {
	GetByID(id uint, expand ...string) (*models.Student, error)
	GetAll(config GetAllConfig) (data []*models.Student, count uint, err error)
	Create(data *models.Student) (*models.Student, error)
	Update(id uint, data *models.Student) (*models.Student, error)
	SoftDelete(id uint) (*models.Student, error)
	HardDelete(id uint) (*models.Student, error)
	Restore(id uint) (*models.Student, error)
	ChargeWithID(params *types.ChargeWithIDParams) (*stripe.Charge, error)
}

type studentServiceImpl struct {
	repo          StudentRepository
	stripeService StripeService
}

func NewStudentService(repository StudentRepository, stripeService StripeService) StudentService {
	res := &studentServiceImpl{}
	res.repo = repository
	res.stripeService = stripeService
	return res
}

func (c *studentServiceImpl) GetByID(id uint, expand ...string) (*models.Student, error) {
	return c.repo.GetByID(id, expand...)
}

func (c *studentServiceImpl) GetAll(config GetAllConfig) ([]*models.Student, uint, error) {
	return c.repo.GetAll(config)
}

func (c *studentServiceImpl) Create(data *models.Student) (*models.Student, error) {
	return c.repo.Create(data)
}

func (c *studentServiceImpl) Update(id uint, data *models.Student) (*models.Student, error) {
	return c.repo.Update(id, data)
}

func (c *studentServiceImpl) SoftDelete(id uint) (*models.Student, error) {
	return c.repo.SoftDelete(id)
}

func (c *studentServiceImpl) HardDelete(id uint) (*models.Student, error) {
	return c.repo.HardDelete(id)
}

func (c *studentServiceImpl) Restore(id uint) (*models.Student, error) {
	return c.repo.Restore(id)
}
func (c *studentServiceImpl) ChargeWithID(params *types.ChargeWithIDParams) (*stripe.Charge, error) {
	student, err := c.GetByID(params.StudentID)
	if err != nil {
		return nil, err
	}
	chargeParams := &stripe.ChargeParams{
		Amount:   &params.Amount,
		Customer: &student.StripeID,
	}
	return c.stripeService.Charge(chargeParams)
}
