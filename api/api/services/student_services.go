package services

import "cramee/models"

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
}

type studentServiceImpl struct {
	repo StudentRepository
	StudentService
}

func NewStudentService(repository StudentRepository) StudentService {
	res := &studentServiceImpl{}
	res.repo = repository
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
