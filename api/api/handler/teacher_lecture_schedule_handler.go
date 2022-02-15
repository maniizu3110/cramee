package handler

import (
	"cramee/api/repository"
	"cramee/api/services"
	"cramee/models"
	"cramee/token"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AssignTeacherLectureScheduleHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db := c.Get("Tx").(*gorm.DB)
			token := c.Get("token").(*token.Payload)
			r := repository.NewTeacherLectureScheduleRepository(db)
			sl := repository.NewStudentLectureScheduleRepository(db)
			s := services.NewTeacherLectureScheduleService(r, sl, token)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("", CreateTeacherLectureSchedule)
	g.PUT("/:id", UpdateTeacherLectureSchedule)
	g.DELETE("/:id", DeleteTeacherLectureSchedule)
	g.PUT("/:id/restore", RestoreTeacherLectureSchedule)
	g.GET("/:id", GetTeacherLectureScheduleByID)
	g.GET("", GetTeacherLectureScheduleList)
	g.PUT("/with-student-schedule/:id", UpdateTeacherLectureScheduleWithStudentLectureSchedule)
}

type CreateTeacherLectureScheduleBaseCallbackFunc func(services.TeacherLectureScheduleService, *models.TeacherLectureSchedule) (*models.TeacherLectureSchedule, error)

func CreateTeacherLectureScheduleBase(c echo.Context, params interface{}, callback CreateTeacherLectureScheduleBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.TeacherLectureScheduleService)

	data := &models.TeacherLectureSchedule{}
	if err != nil {
		return err
	}
	if err = c.Bind(data); err != nil {
		return err
	}
	if params != nil {
		if err = c.Bind(params); err != nil {
			return err
		}
	}
	if err = c.Validate(data); err != nil {
		return err
	}
	m, err := callback(service, data)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, m)
}

func CreateTeacherLectureSchedule(c echo.Context) (err error) {
	return CreateTeacherLectureScheduleBase(c, nil, func(service services.TeacherLectureScheduleService, data *models.TeacherLectureSchedule) (*models.TeacherLectureSchedule, error) {
		return service.Create(data)
	})
}

type UpdateTeacherLectureScheduleBaseCallbackFunc func(services.TeacherLectureScheduleService, uint, *models.TeacherLectureSchedule) (*models.TeacherLectureSchedule, error)

func UpdateTeacherLectureScheduleBase(c echo.Context, params interface{}, callback UpdateTeacherLectureScheduleBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.TeacherLectureScheduleService)

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errors.New(err.Error())
	}
	data, err := service.GetByID(uint(id))
	if err != nil {
		return err
	}
	if err = c.Bind(data); err != nil {
		return errors.New(err.Error())
	}
	if params != nil {
		if err = c.Bind(params); err != nil {
			return errors.New(err.Error())
		}
	}
	if err = c.Validate(data); err != nil {
		return err
	}
	m, err := callback(service, uint(id), data)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, m)
}

func UpdateTeacherLectureSchedule(c echo.Context) (err error) {
	return UpdateTeacherLectureScheduleBase(c, nil, func(service services.TeacherLectureScheduleService, id uint, data *models.TeacherLectureSchedule) (*models.TeacherLectureSchedule, error) {
		return service.Update(uint(id), data)
	})
}
func UpdateTeacherLectureScheduleWithStudentLectureSchedule(c echo.Context) (err error) {
	return UpdateTeacherLectureScheduleBase(c, nil, func(service services.TeacherLectureScheduleService, id uint, data *models.TeacherLectureSchedule) (*models.TeacherLectureSchedule, error) {
		return service.UpdateWithStudentLectureSchedule(uint(id), data)
	})
}

type DeleteTeacherLectureScheduleBaseCallbackFunc func(services.TeacherLectureScheduleService, uint) (*models.TeacherLectureSchedule, error)

func DeleteTeacherLectureScheduleBase(c echo.Context, params interface{}, callback DeleteTeacherLectureScheduleBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.TeacherLectureScheduleService)

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errors.New(err.Error())
	}
	if params != nil {
		if err = c.Bind(params); err != nil {
			return errors.New(err.Error())
		}
	}
	data, err := callback(service, uint(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, data)
}

func DeleteTeacherLectureSchedule(c echo.Context) (err error) {
	var param struct {
		HardDelete bool
	}
	return DeleteTeacherLectureScheduleBase(c, &param, func(service services.TeacherLectureScheduleService, id uint) (*models.TeacherLectureSchedule, error) {
		if param.HardDelete {
			return service.HardDelete(id)
		}
		return service.SoftDelete(id)
	})
}

type RestoreTeacherLectureScheduleBaseCallbackFunc func(services.TeacherLectureScheduleService, uint) (*models.TeacherLectureSchedule, error)

func RestoreTeacherLectureScheduleBase(c echo.Context, params interface{}, callback RestoreTeacherLectureScheduleBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.TeacherLectureScheduleService)

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errors.New(err.Error())
	}
	if params != nil {
		if err = c.Bind(params); err != nil {
			return errors.New(err.Error())
		}
	}
	m, err := callback(service, uint(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, m)
}

func RestoreTeacherLectureSchedule(c echo.Context) (err error) {
	return RestoreTeacherLectureScheduleBase(c, nil, func(service services.TeacherLectureScheduleService, id uint) (*models.TeacherLectureSchedule, error) {
		return service.Restore(id)
	})
}

type GetTeacherLectureScheduleByIDBaseCallbackFunc func(services.TeacherLectureScheduleService, uint) (*models.TeacherLectureSchedule, error)

func GetTeacherLectureScheduleByIDBase(c echo.Context, params interface{}, callback GetTeacherLectureScheduleByIDBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.TeacherLectureScheduleService)

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errors.New(err.Error())
	}
	if params != nil {
		if err = c.Bind(params); err != nil {
			return errors.New(err.Error())
		}
	}
	m, err := callback(service, uint(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, m)
}

func GetTeacherLectureScheduleByID(c echo.Context) (err error) {
	var param struct {
		Expand []string
	}
	return GetTeacherLectureScheduleByIDBase(c, &param, func(service services.TeacherLectureScheduleService, id uint) (*models.TeacherLectureSchedule, error) {
		return service.GetByID(id, param.Expand...)
	})
}

type GetTeacherLectureScheduleListResponse struct {
	AllCount uint `json:",omitempty"`
	Offset   uint `json:",omitempty"`
	Limit    uint `json:",omitempty"`
	Data     []*models.TeacherLectureSchedule
}

type GetTeacherLectureScheduleListBaseCallbackFunc func(services.TeacherLectureScheduleService) (*GetTeacherLectureScheduleListResponse, error)

func GetTeacherLectureScheduleListBase(c echo.Context, params interface{}, callback GetTeacherLectureScheduleListBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.TeacherLectureScheduleService)
	if params != nil {
		if err = c.Bind(params); err != nil {
			return errors.New(err.Error())
		}
	}
	response, err := callback(service)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response)
}

func GetTeacherLectureScheduleList(c echo.Context) (err error) {
	var param struct {
		services.GetAllConfig
	}

	return GetTeacherLectureScheduleListBase(c, &param, func(service services.TeacherLectureScheduleService) (*GetTeacherLectureScheduleListResponse, error) {
		m, count, err := service.GetAll(param.GetAllConfig)
		if err != nil {
			return nil, err
		}
		return &GetTeacherLectureScheduleListResponse{
			AllCount: count,
			Limit:    param.Limit,
			Offset:   param.Offset,
			Data:     m,
		}, nil
	})
}
