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

func AssignStudentLectureScheduleHandlers(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db := c.Get("Tx").(*gorm.DB)
			r := repository.NewStudentLectureScheduleRepository(db)
			s := services.NewStudentLectureScheduleService(r)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("", CreateStudentLectureScheduleHandler)
	g.PUT("/:id", UpdateStudentLectureScheduleHandler)
	g.DELETE("/:id", DeleteStudentLectureScheduleHandler)
	g.PUT("/:id/restore", RestoreStudentLectureScheduleHandler)
	g.GET("/:id", GetStudentLectureScheduleByIDHandler)
	g.GET("", GetStudentLectureScheduleListHandler)
}

type CreateStudentLectureScheduleHandlerBaseCallbackFunc func(services.StudentLectureScheduleService, *models.StudentLectureSchedule) (*models.StudentLectureSchedule, error)

func CreateStudentLectureScheduleHandlerBase(c echo.Context, params interface{}, callback CreateStudentLectureScheduleHandlerBaseCallbackFunc) (err error) {
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

func CreateStudentLectureScheduleHandler(c echo.Context) (err error) {
	return CreateStudentLectureScheduleHandlerBase(c, nil, func(service services.StudentLectureScheduleService, data *models.StudentLectureSchedule) (*models.StudentLectureSchedule, error) {
		return service.Create(data)
	})
}

type UpdateStudentLectureScheduleHandlerBaseCallbackFunc func(services.StudentLectureScheduleService, uint, *models.StudentLectureSchedule) (*models.StudentLectureSchedule, error)

func UpdateStudentLectureScheduleHandlerBase(c echo.Context, params interface{}, callback UpdateStudentLectureScheduleHandlerBaseCallbackFunc) (err error) {
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

func UpdateStudentLectureScheduleHandler(c echo.Context) (err error) {
	return UpdateStudentLectureScheduleHandlerBase(c, nil, func(service services.StudentLectureScheduleService, id uint, data *models.StudentLectureSchedule) (*models.StudentLectureSchedule, error) {
		return service.Update(uint(id), data)
	})
}

type DeleteStudentLectureScheduleHandlerBaseCallbackFunc func(services.StudentLectureScheduleService, uint) (*models.StudentLectureSchedule, error)

func DeleteStudentLectureScheduleHandlerBase(c echo.Context, params interface{}, callback DeleteStudentLectureScheduleHandlerBaseCallbackFunc) (err error) {
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

func DeleteStudentLectureScheduleHandler(c echo.Context) (err error) {
	var param struct {
		HardDelete bool
	}
	return DeleteStudentLectureScheduleHandlerBase(c, &param, func(service services.StudentLectureScheduleService, id uint) (*models.StudentLectureSchedule, error) {
		if param.HardDelete {
			return service.HardDelete(id)
		}
		return service.SoftDelete(id)
	})
}

type RestoreStudentLectureScheduleHandlerBaseCallbackFunc func(services.StudentLectureScheduleService, uint) (*models.StudentLectureSchedule, error)

func RestoreStudentLectureScheduleHandlerBase(c echo.Context, params interface{}, callback RestoreStudentLectureScheduleHandlerBaseCallbackFunc) (err error) {
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

func RestoreStudentLectureScheduleHandler(c echo.Context) (err error) {
	return RestoreStudentLectureScheduleHandlerBase(c, nil, func(service services.StudentLectureScheduleService, id uint) (*models.StudentLectureSchedule, error) {
		return service.Restore(id)
	})
}

type GetStudentLectureScheduleByIDHandlerBaseCallbackFunc func(services.StudentLectureScheduleService, uint) (*models.StudentLectureSchedule, error)

func GetStudentLectureScheduleByIDHandlerBase(c echo.Context, params interface{}, callback GetStudentLectureScheduleByIDHandlerBaseCallbackFunc) (err error) {
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

func GetStudentLectureScheduleByIDHandler(c echo.Context) (err error) {
	var param struct {
		Expand []string
	}
	return GetStudentLectureScheduleByIDHandlerBase(c, &param, func(service services.StudentLectureScheduleService, id uint) (*models.StudentLectureSchedule, error) {
		return service.GetByID(id, param.Expand...)
	})
}

type GetStudentLectureScheduleListResponse struct {
	AllCount uint `json:",omitempty"`
	Offset   uint `json:",omitempty"`
	Limit    uint `json:",omitempty"`
	Data     []*models.StudentLectureSchedule
}

type GetStudentLectureScheduleListHandlerBaseCallbackFunc func(services.StudentLectureScheduleService) (*GetStudentLectureScheduleListResponse, error)

func GetStudentLectureScheduleListHandlerBase(c echo.Context, params interface{}, callback GetStudentLectureScheduleListHandlerBaseCallbackFunc) (err error) {
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

func GetStudentLectureScheduleListHandler(c echo.Context) (err error) {
	var param struct {
		services.GetAllConfig
	}

	return GetStudentLectureScheduleListHandlerBase(c, &param, func(service services.StudentLectureScheduleService) (*GetStudentLectureScheduleListResponse, error) {
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
