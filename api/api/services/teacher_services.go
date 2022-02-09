package services

import "cramee/models"

//go:generate $GOPATH/bin/mockgen -source=$GOFILE -destination=${GOPACKAGE}_mock/${GOFILE}.mock.go -package=${GOPACKAGE}_mock

type TeacherRepository interface {
	GetByID(id uint, expand ...string) (*models.Teacher, error)
	GetAll(config GetAllConfig) (data []*models.Teacher, count uint, err error)
	Create(data *models.Teacher) (*models.Teacher, error)
	Update(id uint, data *models.Teacher) (*models.Teacher, error)
	SoftDelete(id uint) (*models.Teacher, error)
	HardDelete(id uint) (*models.Teacher, error)
	Restore(id uint) (*models.Teacher, error)
	GetByEmail(email string) (*models.Teacher, error)
}

type TeacherService interface {
	GetByID(id uint, expand ...string) (*models.Teacher, error)
	GetAll(config GetAllConfig) (data []*models.Teacher, count uint, err error)
	Create(data *models.Teacher) (*models.Teacher, error)
	Update(id uint, data *models.Teacher) (*models.Teacher, error)
	SoftDelete(id uint) (*models.Teacher, error)
	HardDelete(id uint) (*models.Teacher, error)
	Restore(id uint) (*models.Teacher, error)
}

type teacherServiceImpl struct {
	repo TeacherRepository
	TeacherService
}

func NewTeacherService(repository TeacherRepository) TeacherService {
	res := &teacherServiceImpl{}
	res.repo = repository
	return res
}

func (c *teacherServiceImpl) GetByID(id uint, expand ...string) (*models.Teacher, error) {
	return c.repo.GetByID(id, expand...)
}

func (c *teacherServiceImpl) GetAll(config GetAllConfig) ([]*models.Teacher, uint, error) {
	return c.repo.GetAll(config)
}

func (c *teacherServiceImpl) Create(data *models.Teacher) (*models.Teacher, error) {
	return c.repo.Create(data)
}

func (c *teacherServiceImpl) Update(id uint, data *models.Teacher) (*models.Teacher, error) {
	return c.repo.Update(id, data)
}

func (c *teacherServiceImpl) SoftDelete(id uint) (*models.Teacher, error) {
	return c.repo.SoftDelete(id)
}

func (c *teacherServiceImpl) HardDelete(id uint) (*models.Teacher, error) {
	return c.repo.HardDelete(id)
}

func (c *teacherServiceImpl) Restore(id uint) (*models.Teacher, error) {
	return c.repo.Restore(id)
}
