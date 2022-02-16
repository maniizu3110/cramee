package services

import (
	"cramee/models"

)

//go:generate $GOPATH/bin/mockgen -source=$GOFILE -destination=${GOPACKAGE}_mock/${GOFILE}.mock.go -package=${GOPACKAGE}_mock

type LectureRepository interface {
	GetByID(id uint, expand ...string) (*models.Lecture, error)
	GetAll(config GetAllConfig) (data []*models.Lecture, count uint, err error)
	Create(data *models.Lecture) (*models.Lecture, error)
	Update(id uint, data *models.Lecture) (*models.Lecture, error)
	SoftDelete(id uint) (*models.Lecture, error)
	HardDelete(id uint) (*models.Lecture, error)
	Restore(id uint) (*models.Lecture, error)
}

type LectureService interface {
	GetByID(id uint, expand ...string) (*models.Lecture, error)
	GetAll(config GetAllConfig) (data []*models.Lecture, count uint, err error)
	Create(data *models.Lecture) (*models.Lecture, error)
	Update(id uint, data *models.Lecture) (*models.Lecture, error)
	SoftDelete(id uint) (*models.Lecture, error)
	HardDelete(id uint) (*models.Lecture, error)
	Restore(id uint) (*models.Lecture, error)
}

type lectureServiceImpl struct {
	repo LectureRepository
	LectureService
}

func NewLectureService(repository LectureRepository) LectureService {
	res := &lectureServiceImpl{}
	res.repo = repository
	return res
}

func (c *lectureServiceImpl) GetByID(id uint, expand ...string) (*models.Lecture, error) {
	return c.repo.GetByID(id, expand...)
}

func (c *lectureServiceImpl) GetAll(config GetAllConfig) ([]*models.Lecture, uint, error) {
	return c.repo.GetAll(config)
}

func (c *lectureServiceImpl) Create(data *models.Lecture) (*models.Lecture, error) {
	return c.repo.Create(data)
}

func (c *lectureServiceImpl) Update(id uint, data *models.Lecture) (*models.Lecture, error) {
	return c.repo.Update(id, data)
}

func (c *lectureServiceImpl) SoftDelete(id uint) (*models.Lecture, error) {
	return c.repo.SoftDelete(id)
}

func (c *lectureServiceImpl) HardDelete(id uint) (*models.Lecture, error) {
	return c.repo.HardDelete(id)
}

func (c *lectureServiceImpl) Restore(id uint) (*models.Lecture, error) {
	return c.repo.Restore(id)
}
