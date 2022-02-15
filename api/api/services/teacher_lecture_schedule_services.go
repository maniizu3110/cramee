package services

import (
	"cramee/models"
	"cramee/token"

	"github.com/sirupsen/logrus"
)

//go:generate $GOPATH/bin/mockgen -source=$GOFILE -destination=${GOPACKAGE}_mock/${GOFILE}.mock.go -package=${GOPACKAGE}_mock

type TeacherLectureScheduleRepository interface {
	GetByID(id uint, expand ...string) (*models.TeacherLectureSchedule, error)
	GetAll(config GetAllConfig) (data []*models.TeacherLectureSchedule, count uint, err error)
	Create(data *models.TeacherLectureSchedule) (*models.TeacherLectureSchedule, error)
	Update(id uint, data *models.TeacherLectureSchedule) (*models.TeacherLectureSchedule, error)
	SoftDelete(id uint) (*models.TeacherLectureSchedule, error)
	HardDelete(id uint) (*models.TeacherLectureSchedule, error)
	Restore(id uint) (*models.TeacherLectureSchedule, error)
}

type TeacherLectureScheduleService interface {
	GetByID(id uint, expand ...string) (*models.TeacherLectureSchedule, error)
	GetAll(config GetAllConfig) (data []*models.TeacherLectureSchedule, count uint, err error)
	Create(data *models.TeacherLectureSchedule) (*models.TeacherLectureSchedule, error)
	Update(id uint, data *models.TeacherLectureSchedule) (*models.TeacherLectureSchedule, error)
	SoftDelete(id uint) (*models.TeacherLectureSchedule, error)
	HardDelete(id uint) (*models.TeacherLectureSchedule, error)
	Restore(id uint) (*models.TeacherLectureSchedule, error)
	UpdateWithStudentLectureSchedule(id uint, data *models.TeacherLectureSchedule) (*models.TeacherLectureSchedule, error)
}

type teacherLectureScheduleServiceImpl struct {
	token                      *token.Payload
	repo                       TeacherLectureScheduleRepository
	studentLectureScheduleRepo StudentLectureScheduleRepository
	TeacherLectureScheduleService
}

func NewTeacherLectureScheduleService(repository TeacherLectureScheduleRepository, studentLectureScheduleRepo StudentLectureScheduleRepository, token *token.Payload) TeacherLectureScheduleService {
	res := &teacherLectureScheduleServiceImpl{}
	res.token = token
	res.repo = repository
	res.studentLectureScheduleRepo = studentLectureScheduleRepo
	return res
}

func (c *teacherLectureScheduleServiceImpl) GetByID(id uint, expand ...string) (*models.TeacherLectureSchedule, error) {
	return c.repo.GetByID(id, expand...)
}

func (c *teacherLectureScheduleServiceImpl) GetAll(config GetAllConfig) ([]*models.TeacherLectureSchedule, uint, error) {
	return c.repo.GetAll(config)
}

func (c *teacherLectureScheduleServiceImpl) Create(data *models.TeacherLectureSchedule) (*models.TeacherLectureSchedule, error) {
	return c.repo.Create(data)
}

func (c *teacherLectureScheduleServiceImpl) Update(id uint, data *models.TeacherLectureSchedule) (*models.TeacherLectureSchedule, error) {
	return c.repo.Update(id, data)
}

func (c *teacherLectureScheduleServiceImpl) SoftDelete(id uint) (*models.TeacherLectureSchedule, error) {
	return c.repo.SoftDelete(id)
}

func (c *teacherLectureScheduleServiceImpl) HardDelete(id uint) (*models.TeacherLectureSchedule, error) {
	return c.repo.HardDelete(id)
}

func (c *teacherLectureScheduleServiceImpl) Restore(id uint) (*models.TeacherLectureSchedule, error) {
	return c.repo.Restore(id)
}

func (c *teacherLectureScheduleServiceImpl) UpdateWithStudentLectureSchedule(id uint, data *models.TeacherLectureSchedule) (*models.TeacherLectureSchedule, error) {
	logrus.Info(c.token.ID)
	var err error
	//scheduleの予約をするのは生徒のみ
	if c.token.IsTeacher {
		//TODO:適切なエラーメッセージの設定
		return nil, err
	}
	teacherLectureSchedule, err := c.repo.Update(id, data)
	if err != nil {
		return nil, err
	}
	_, err = c.studentLectureScheduleRepo.Create(&models.StudentLectureSchedule{
		StudentID: c.token.ID,
		StartTime: teacherLectureSchedule.StartTime,
		EndTime:   teacherLectureSchedule.EndTime,
		Status:    "pending",
	})
	if err != nil {
		return nil, err
	}
	return teacherLectureSchedule, nil
}
