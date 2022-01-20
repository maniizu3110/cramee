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

func AssignTeacherLectureScheduleHandlers(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db := c.Get("Tx").(*gorm.DB)
			r := repository.NewTeacherLectureScheduleRepository(db)
			s := services.NewTeacherLectureScheduleService(r)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("", CreateTeacherLectureScheduleHandler)
	g.PUT("/:id", UpdateTeacherLectureScheduleHandler)
	g.DELETE("/:id", DeleteTeacherLectureScheduleHandler)
	g.PUT("/:id/restore", RestoreTeacherLectureScheduleHandler)
	g.GET("/:id", GetTeacherLectureScheduleByIDHandler)
	g.GET("", GetTeacherLectureScheduleListHandler)
}

type CreateTeacherLectureScheduleHandlerBaseCallbackFunc func(services.TeacherLectureScheduleService, *models.TeacherLectureSchedule) (*models.TeacherLectureSchedule, error)

func CreateTeacherLectureScheduleHandlerBase(c echo.Context, params interface{}, callback CreateTeacherLectureScheduleHandlerBaseCallbackFunc) (err error) {
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

func CreateTeacherLectureScheduleHandler(c echo.Context) (err error) {
	return CreateTeacherLectureScheduleHandlerBase(c, nil, func(service services.TeacherLectureScheduleService, data *models.TeacherLectureSchedule) (*models.TeacherLectureSchedule, error) {
		return service.Create(data)
	})
}

type UpdateTeacherLectureScheduleHandlerBaseCallbackFunc func(services.TeacherLectureScheduleService, uint, *models.TeacherLectureSchedule) (*models.TeacherLectureSchedule, error)

func UpdateTeacherLectureScheduleHandlerBase(c echo.Context, params interface{}, callback UpdateTeacherLectureScheduleHandlerBaseCallbackFunc) (err error) {
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

func UpdateTeacherLectureScheduleHandler(c echo.Context) (err error) {
	return UpdateTeacherLectureScheduleHandlerBase(c, nil, func(service services.TeacherLectureScheduleService, id uint, data *models.TeacherLectureSchedule) (*models.TeacherLectureSchedule, error) {
		return service.Update(uint(id), data)
	})
}

type DeleteTeacherLectureScheduleHandlerBaseCallbackFunc func(services.TeacherLectureScheduleService, uint) (*models.TeacherLectureSchedule, error)

func DeleteTeacherLectureScheduleHandlerBase(c echo.Context, params interface{}, callback DeleteTeacherLectureScheduleHandlerBaseCallbackFunc) (err error) {
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

func DeleteTeacherLectureScheduleHandler(c echo.Context) (err error) {
	var param struct {
		HardDelete bool
	}
	return DeleteTeacherLectureScheduleHandlerBase(c, &param, func(service services.TeacherLectureScheduleService, id uint) (*models.TeacherLectureSchedule, error) {
		if param.HardDelete {
			return service.HardDelete(id)
		}
		return service.SoftDelete(id)
	})
}

type RestoreTeacherLectureScheduleHandlerBaseCallbackFunc func(services.TeacherLectureScheduleService, uint) (*models.TeacherLectureSchedule, error)

func RestoreTeacherLectureScheduleHandlerBase(c echo.Context, params interface{}, callback RestoreTeacherLectureScheduleHandlerBaseCallbackFunc) (err error) {
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

func RestoreTeacherLectureScheduleHandler(c echo.Context) (err error) {
	return RestoreTeacherLectureScheduleHandlerBase(c, nil, func(service services.TeacherLectureScheduleService, id uint) (*models.TeacherLectureSchedule, error) {
		return service.Restore(id)
	})
}

type GetTeacherLectureScheduleByIDHandlerBaseCallbackFunc func(services.TeacherLectureScheduleService, uint) (*models.TeacherLectureSchedule, error)

func GetTeacherLectureScheduleByIDHandlerBase(c echo.Context, params interface{}, callback GetTeacherLectureScheduleByIDHandlerBaseCallbackFunc) (err error) {
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

func GetTeacherLectureScheduleByIDHandler(c echo.Context) (err error) {
	var param struct {
		Expand []string
	}
	return GetTeacherLectureScheduleByIDHandlerBase(c, &param, func(service services.TeacherLectureScheduleService, id uint) (*models.TeacherLectureSchedule, error) {
		return service.GetByID(id, param.Expand...)
	})
}

type GetTeacherLectureScheduleListResponse struct {
	AllCount uint `json:",omitempty"`
	Offset   uint `json:",omitempty"`
	Limit    uint `json:",omitempty"`
	Data     []*models.TeacherLectureSchedule
}

type GetTeacherLectureScheduleListHandlerBaseCallbackFunc func(services.TeacherLectureScheduleService) (*GetTeacherLectureScheduleListResponse, error)

func GetTeacherLectureScheduleListHandlerBase(c echo.Context, params interface{}, callback GetTeacherLectureScheduleListHandlerBaseCallbackFunc) (err error) {
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

func GetTeacherLectureScheduleListHandler(c echo.Context) (err error) {
	var param struct {
		services.GetAllConfig
	}

	return GetTeacherLectureScheduleListHandlerBase(c, &param, func(service services.TeacherLectureScheduleService) (*GetTeacherLectureScheduleListResponse, error) {
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
