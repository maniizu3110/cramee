package services

import "cramee/models"

//go:generate $GOPATH/bin/mockgen -source=$GOFILE -destination=${GOPACKAGE}_mock/${GOFILE}.mock.go -package=${GOPACKAGE}_mock

type StudentLectureScheduleRepository interface {
	GetByID(id uint, expand ...string) (*models.StudentLectureSchedule, error)
	GetAll(config GetAllConfig) (data []*models.StudentLectureSchedule, count uint, err error)
	Create(data *models.StudentLectureSchedule) (*models.StudentLectureSchedule, error)
	Update(id uint, data *models.StudentLectureSchedule) (*models.StudentLectureSchedule, error)
	SoftDelete(id uint) (*models.StudentLectureSchedule, error)
	HardDelete(id uint) (*models.StudentLectureSchedule, error)
	Restore(id uint) (*models.StudentLectureSchedule, error)
}

type StudentLectureScheduleService interface {
	GetByID(id uint, expand ...string) (*models.StudentLectureSchedule, error)
	GetAll(config GetAllConfig) (data []*models.StudentLectureSchedule, count uint, err error)
	Create(data *models.StudentLectureSchedule) (*models.StudentLectureSchedule, error)
	Update(id uint, data *models.StudentLectureSchedule) (*models.StudentLectureSchedule, error)
	SoftDelete(id uint) (*models.StudentLectureSchedule, error)
	HardDelete(id uint) (*models.StudentLectureSchedule, error)
	Restore(id uint) (*models.StudentLectureSchedule, error)
}

type studentLectureScheduleServiceImpl struct {
	repo StudentLectureScheduleRepository
	StudentLectureScheduleService
}

func NewStudentLectureScheduleService(repository StudentLectureScheduleRepository) StudentLectureScheduleService {
	res := &studentLectureScheduleServiceImpl{}
	res.repo = repository
	return res
}

func (c *studentLectureScheduleServiceImpl) GetByID(id uint, expand ...string) (*models.StudentLectureSchedule, error) {
	return c.repo.GetByID(id, expand...)
}

func (c *studentLectureScheduleServiceImpl) GetAll(config GetAllConfig) ([]*models.StudentLectureSchedule, uint, error) {
	return c.repo.GetAll(config)
}

func (c *studentLectureScheduleServiceImpl) Create(data *models.StudentLectureSchedule) (*models.StudentLectureSchedule, error) {
	return c.repo.Create(data)
}

func (c *studentLectureScheduleServiceImpl) Update(id uint, data *models.StudentLectureSchedule) (*models.StudentLectureSchedule, error) {
	return c.repo.Update(id, data)
}

func (c *studentLectureScheduleServiceImpl) SoftDelete(id uint) (*models.StudentLectureSchedule, error) {
	return c.repo.SoftDelete(id)
}

func (c *studentLectureScheduleServiceImpl) HardDelete(id uint) (*models.StudentLectureSchedule, error) {
	return c.repo.HardDelete(id)
}

func (c *studentLectureScheduleServiceImpl) Restore(id uint) (*models.StudentLectureSchedule, error) {
	return c.repo.Restore(id)
}
