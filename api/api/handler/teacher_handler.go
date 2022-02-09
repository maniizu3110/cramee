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

func AssignTeacherHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db := c.Get("Tx").(*gorm.DB)
			r := repository.NewTeacherRepository(db)
			s := services.NewTeacherService(r)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("", CreateTeacher)
	g.PUT("/:id", UpdateTeacher)
	g.DELETE("/:id", DeleteTeacher)
	g.PUT("/:id/restore", RestoreTeacher)
	g.GET("/:id", GetTeacherByID)
	g.GET("", GetTeacherList)
}

type CreateTeacherBaseCallbackFunc func(services.TeacherService, *models.Teacher) (*models.Teacher, error)

func CreateTeacherBase(c echo.Context, params interface{}, callback CreateTeacherBaseCallbackFunc) (err error) {
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

func CreateTeacher(c echo.Context) (err error) {
	return CreateTeacherBase(c, nil, func(service services.TeacherService, data *models.Teacher) (*models.Teacher, error) {
		return service.Create(data)
	})
}

type UpdateTeacherBaseCallbackFunc func(services.TeacherService, uint, *models.Teacher) (*models.Teacher, error)

func UpdateTeacherBase(c echo.Context, params interface{}, callback UpdateTeacherBaseCallbackFunc) (err error) {
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

func UpdateTeacher(c echo.Context) (err error) {
	return UpdateTeacherBase(c, nil, func(service services.TeacherService, id uint, data *models.Teacher) (*models.Teacher, error) {
		return service.Update(uint(id), data)
	})
}

type DeleteTeacherBaseCallbackFunc func(services.TeacherService, uint) (*models.Teacher, error)

func DeleteTeacherBase(c echo.Context, params interface{}, callback DeleteTeacherBaseCallbackFunc) (err error) {
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

func DeleteTeacher(c echo.Context) (err error) {
	var param struct {
		HardDelete bool
	}
	return DeleteTeacherBase(c, &param, func(service services.TeacherService, id uint) (*models.Teacher, error) {
		if param.HardDelete {
			return service.HardDelete(id)
		}
		return service.SoftDelete(id)
	})
}

type RestoreTeacherBaseCallbackFunc func(services.TeacherService, uint) (*models.Teacher, error)

func RestoreTeacherBase(c echo.Context, params interface{}, callback RestoreTeacherBaseCallbackFunc) (err error) {
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

func RestoreTeacher(c echo.Context) (err error) {
	return RestoreTeacherBase(c, nil, func(service services.TeacherService, id uint) (*models.Teacher, error) {
		return service.Restore(id)
	})
}

type GetTeacherByIDBaseCallbackFunc func(services.TeacherService, uint) (*models.Teacher, error)

func GetTeacherByIDBase(c echo.Context, params interface{}, callback GetTeacherByIDBaseCallbackFunc) (err error) {
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

func GetTeacherByID(c echo.Context) (err error) {
	var param struct {
		Expand []string
	}
	return GetTeacherByIDBase(c, &param, func(service services.TeacherService, id uint) (*models.Teacher, error) {
		return service.GetByID(id, param.Expand...)
	})
}

type GetTeacherListResponse struct {
	AllCount uint `json:",omitempty"`
	Offset   uint `json:",omitempty"`
	Limit    uint `json:",omitempty"`
	Data     []*models.Teacher
}

type GetTeacherListBaseCallbackFunc func(services.TeacherService) (*GetTeacherListResponse, error)

func GetTeacherListBase(c echo.Context, params interface{}, callback GetTeacherListBaseCallbackFunc) (err error) {
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

func GetTeacherList(c echo.Context) (err error) {
	var param struct {
		services.GetAllConfig
	}

	return GetTeacherListBase(c, &param, func(service services.TeacherService) (*GetTeacherListResponse, error) {
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
