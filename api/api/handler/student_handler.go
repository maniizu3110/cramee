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

func AssignStudentHandlers(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db := c.Get("Tx").(*gorm.DB)
			r := repository.NewStudentRepository(db)
			s := services.NewStudentService(r)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("", CreateStudentHandler)
	g.PUT("/:id", UpdateStudentHandler)
	g.DELETE("/:id", DeleteStudentHandler)
	g.PUT("/:id/restore", RestoreStudentHandler)
	g.GET("/:id", GetStudentByIDHandler)
	g.GET("", GetStudentListHandler)
}

type CreateStudentHandlerBaseCallbackFunc func(services.StudentService, *models.Student) (*models.Student, error)

func CreateStudentHandlerBase(c echo.Context, params interface{}, callback CreateStudentHandlerBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.StudentService)

	data := &models.Student{}
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

func CreateStudentHandler(c echo.Context) (err error) {
	return CreateStudentHandlerBase(c, nil, func(service services.StudentService, data *models.Student) (*models.Student, error) {
		return service.Create(data)
	})
}

type UpdateStudentHandlerBaseCallbackFunc func(services.StudentService, uint, *models.Student) (*models.Student, error)

func UpdateStudentHandlerBase(c echo.Context, params interface{}, callback UpdateStudentHandlerBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.StudentService)

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

func UpdateStudentHandler(c echo.Context) (err error) {
	return UpdateStudentHandlerBase(c, nil, func(service services.StudentService, id uint, data *models.Student) (*models.Student, error) {
		return service.Update(uint(id), data)
	})
}

type DeleteStudentHandlerBaseCallbackFunc func(services.StudentService, uint) (*models.Student, error)

func DeleteStudentHandlerBase(c echo.Context, params interface{}, callback DeleteStudentHandlerBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.StudentService)

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

func DeleteStudentHandler(c echo.Context) (err error) {
	var param struct {
		HardDelete bool
	}
	return DeleteStudentHandlerBase(c, &param, func(service services.StudentService, id uint) (*models.Student, error) {
		if param.HardDelete {
			return service.HardDelete(id)
		}
		return service.SoftDelete(id)
	})
}

type RestoreStudentHandlerBaseCallbackFunc func(services.StudentService, uint) (*models.Student, error)

func RestoreStudentHandlerBase(c echo.Context, params interface{}, callback RestoreStudentHandlerBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.StudentService)

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

func RestoreStudentHandler(c echo.Context) (err error) {
	return RestoreStudentHandlerBase(c, nil, func(service services.StudentService, id uint) (*models.Student, error) {
		return service.Restore(id)
	})
}

type GetStudentByIDHandlerBaseCallbackFunc func(services.StudentService, uint) (*models.Student, error)

func GetStudentByIDHandlerBase(c echo.Context, params interface{}, callback GetStudentByIDHandlerBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.StudentService)

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

func GetStudentByIDHandler(c echo.Context) (err error) {
	var param struct {
		Expand []string
	}
	return GetStudentByIDHandlerBase(c, &param, func(service services.StudentService, id uint) (*models.Student, error) {
		return service.GetByID(id, param.Expand...)
	})
}

type GetStudentListResponse struct {
	AllCount uint `json:",omitempty"`
	Offset   uint `json:",omitempty"`
	Limit    uint `json:",omitempty"`
	Data     []*models.Student
}

type GetStudentListHandlerBaseCallbackFunc func(services.StudentService) (*GetStudentListResponse, error)

func GetStudentListHandlerBase(c echo.Context, params interface{}, callback GetStudentListHandlerBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.StudentService)
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

func GetStudentListHandler(c echo.Context) (err error) {
	var param struct {
		services.GetAllConfig
	}

	return GetStudentListHandlerBase(c, &param, func(service services.StudentService) (*GetStudentListResponse, error) {
		m, count, err := service.GetAll(param.GetAllConfig)
		if err != nil {
			return nil, err
		}
		return &GetStudentListResponse{
			AllCount: count,
			Limit:    param.Limit,
			Offset:   param.Offset,
			Data:     m,
		}, nil
	})
}
