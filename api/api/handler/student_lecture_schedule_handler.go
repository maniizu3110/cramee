package handler

import (
	"cramee/api/repository"
	"cramee/api/services"
	"cramee/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AssignStudentLectureScheduleHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db := c.Get("Tx").(*gorm.DB)
			r := repository.NewStudentLectureScheduleRepository(db)
			s := services.NewStudentLectureScheduleService(r)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("", CreateStudentLectureSchedule)
	g.PUT("/:id", UpdateStudentLectureSchedule)
	g.DELETE("/:id", DeleteStudentLectureSchedule)
	g.PUT("/:id/restore", RestoreStudentLectureSchedule)
	g.GET("/:id", GetStudentLectureScheduleByID)
	g.GET("", GetStudentLectureScheduleList)
}

type CreateStudentLectureScheduleBaseCallbackFunc func(services.StudentLectureScheduleService, *models.StudentLectureSchedule) (*models.StudentLectureSchedule, error)

func CreateStudentLectureScheduleBase(c echo.Context, params interface{}, callback CreateStudentLectureScheduleBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.StudentLectureScheduleService)

	data := &models.StudentLectureSchedule{}
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

func CreateStudentLectureSchedule(c echo.Context) (err error) {
	return CreateStudentLectureScheduleBase(c, nil, func(service services.StudentLectureScheduleService, data *models.StudentLectureSchedule) (*models.StudentLectureSchedule, error) {
		return service.Create(data)
	})
}

type UpdateStudentLectureScheduleBaseCallbackFunc func(services.StudentLectureScheduleService, uint, *models.StudentLectureSchedule) (*models.StudentLectureSchedule, error)

func UpdateStudentLectureScheduleBase(c echo.Context, params interface{}, callback UpdateStudentLectureScheduleBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.StudentLectureScheduleService)

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

func UpdateStudentLectureSchedule(c echo.Context) (err error) {
	return UpdateStudentLectureScheduleBase(c, nil, func(service services.StudentLectureScheduleService, id uint, data *models.StudentLectureSchedule) (*models.StudentLectureSchedule, error) {
		return service.Update(uint(id), data)
	})
}

type DeleteStudentLectureScheduleBaseCallbackFunc func(services.StudentLectureScheduleService, uint) (*models.StudentLectureSchedule, error)

func DeleteStudentLectureScheduleBase(c echo.Context, params interface{}, callback DeleteStudentLectureScheduleBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.StudentLectureScheduleService)

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

func DeleteStudentLectureSchedule(c echo.Context) (err error) {
	var param struct {
		HardDelete bool
	}
	return DeleteStudentLectureScheduleBase(c, &param, func(service services.StudentLectureScheduleService, id uint) (*models.StudentLectureSchedule, error) {
		if param.HardDelete {
			return service.HardDelete(id)
		}
		return service.SoftDelete(id)
	})
}

type RestoreStudentLectureScheduleBaseCallbackFunc func(services.StudentLectureScheduleService, uint) (*models.StudentLectureSchedule, error)

func RestoreStudentLectureScheduleBase(c echo.Context, params interface{}, callback RestoreStudentLectureScheduleBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.StudentLectureScheduleService)

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

func RestoreStudentLectureSchedule(c echo.Context) (err error) {
	return RestoreStudentLectureScheduleBase(c, nil, func(service services.StudentLectureScheduleService, id uint) (*models.StudentLectureSchedule, error) {
		return service.Restore(id)
	})
}

type GetStudentLectureScheduleByIDBaseCallbackFunc func(services.StudentLectureScheduleService, uint) (*models.StudentLectureSchedule, error)

func GetStudentLectureScheduleByIDBase(c echo.Context, params interface{}, callback GetStudentLectureScheduleByIDBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.StudentLectureScheduleService)

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

func GetStudentLectureScheduleByID(c echo.Context) (err error) {
	var param struct {
		Expand []string
	}
	return GetStudentLectureScheduleByIDBase(c, &param, func(service services.StudentLectureScheduleService, id uint) (*models.StudentLectureSchedule, error) {
		return service.GetByID(id, param.Expand...)
	})
}

type GetStudentLectureScheduleListResponse struct {
	AllCount uint `json:",omitempty"`
	Offset   uint `json:",omitempty"`
	Limit    uint `json:",omitempty"`
	Data     []*models.StudentLectureSchedule
}

type GetStudentLectureScheduleListBaseCallbackFunc func(services.StudentLectureScheduleService) (*GetStudentLectureScheduleListResponse, error)

func GetStudentLectureScheduleListBase(c echo.Context, params interface{}, callback GetStudentLectureScheduleListBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.StudentLectureScheduleService)
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

func GetStudentLectureScheduleList(c echo.Context) (err error) {
	var param struct {
		services.GetAllConfig
	}

	return GetStudentLectureScheduleListBase(c, &param, func(service services.StudentLectureScheduleService) (*GetStudentLectureScheduleListResponse, error) {
		m, count, err := service.GetAll(param.GetAllConfig)
		if err != nil {
			return nil, err
		}
		return &GetStudentLectureScheduleListResponse{
			AllCount: count,
			Limit:    param.Limit,
			Offset:   param.Offset,
			Data:     m,
		}, nil
	})
}
