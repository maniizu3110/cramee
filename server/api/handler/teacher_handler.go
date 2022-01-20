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

func AssignTeacherHandlers(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db := c.Get("Tx").(*gorm.DB)
			r := repository.NewTeacherRepository(db)
			s := services.NewTeacherService(r)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("", CreateTeacherHandler)
	g.PUT("/:id", UpdateTeacherHandler)
	g.DELETE("/:id", DeleteTeacherHandler)
	g.PUT("/:id/restore", RestoreTeacherHandler)
	g.GET("/:id", GetTeacherByIDHandler)
	g.GET("", GetTeacherListHandler)
}

type CreateTeacherHandlerBaseCallbackFunc func(services.TeacherService, *models.Teacher) (*models.Teacher, error)

func CreateTeacherHandlerBase(c echo.Context, params interface{}, callback CreateTeacherHandlerBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.TeacherService)

	data := &models.Teacher{}
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

func CreateTeacherHandler(c echo.Context) (err error) {
	return CreateTeacherHandlerBase(c, nil, func(service services.TeacherService, data *models.Teacher) (*models.Teacher, error) {
		return service.Create(data)
	})
}

type UpdateTeacherHandlerBaseCallbackFunc func(services.TeacherService, uint, *models.Teacher) (*models.Teacher, error)

func UpdateTeacherHandlerBase(c echo.Context, params interface{}, callback UpdateTeacherHandlerBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.TeacherService)

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

func UpdateTeacherHandler(c echo.Context) (err error) {
	return UpdateTeacherHandlerBase(c, nil, func(service services.TeacherService, id uint, data *models.Teacher) (*models.Teacher, error) {
		return service.Update(uint(id), data)
	})
}

type DeleteTeacherHandlerBaseCallbackFunc func(services.TeacherService, uint) (*models.Teacher, error)

func DeleteTeacherHandlerBase(c echo.Context, params interface{}, callback DeleteTeacherHandlerBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.TeacherService)

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

func DeleteTeacherHandler(c echo.Context) (err error) {
	var param struct {
		HardDelete bool
	}
	return DeleteTeacherHandlerBase(c, &param, func(service services.TeacherService, id uint) (*models.Teacher, error) {
		if param.HardDelete {
			return service.HardDelete(id)
		}
		return service.SoftDelete(id)
	})
}

type RestoreTeacherHandlerBaseCallbackFunc func(services.TeacherService, uint) (*models.Teacher, error)

func RestoreTeacherHandlerBase(c echo.Context, params interface{}, callback RestoreTeacherHandlerBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.TeacherService)

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

func RestoreTeacherHandler(c echo.Context) (err error) {
	return RestoreTeacherHandlerBase(c, nil, func(service services.TeacherService, id uint) (*models.Teacher, error) {
		return service.Restore(id)
	})
}

type GetTeacherByIDHandlerBaseCallbackFunc func(services.TeacherService, uint) (*models.Teacher, error)

func GetTeacherByIDHandlerBase(c echo.Context, params interface{}, callback GetTeacherByIDHandlerBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.TeacherService)

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

func GetTeacherByIDHandler(c echo.Context) (err error) {
	var param struct {
		Expand []string
	}
	return GetTeacherByIDHandlerBase(c, &param, func(service services.TeacherService, id uint) (*models.Teacher, error) {
		return service.GetByID(id, param.Expand...)
	})
}

type GetTeacherListResponse struct {
	AllCount uint `json:",omitempty"`
	Offset   uint `json:",omitempty"`
	Limit    uint `json:",omitempty"`
	Data     []*models.Teacher
}

type GetTeacherListHandlerBaseCallbackFunc func(services.TeacherService) (*GetTeacherListResponse, error)

func GetTeacherListHandlerBase(c echo.Context, params interface{}, callback GetTeacherListHandlerBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.TeacherService)
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

func GetTeacherListHandler(c echo.Context) (err error) {
	var param struct {
		services.GetAllConfig
	}

	return GetTeacherListHandlerBase(c, &param, func(service services.TeacherService) (*GetTeacherListResponse, error) {
		m, count, err := service.GetAll(param.GetAllConfig)
		if err != nil {
			return nil, err
		}
		return &GetTeacherListResponse{
			AllCount: count,
			Limit:    param.Limit,
			Offset:   param.Offset,
			Data:     m,
		}, nil
	})
}
