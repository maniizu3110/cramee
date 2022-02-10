package handler

import (
	"cramee/api/repository"
	"cramee/api/services"
	"cramee/api/services/types"
	"cramee/models"
	"cramee/token"
	"cramee/util"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72/client"
	"gorm.io/gorm"
)

func AssignStudentHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db := c.Get("Tx").(*gorm.DB)
			conf := c.Get("config").(util.Config)
			tk := c.Get("tk").(token.Maker)
			sc := c.Get("sc").(*client.API)

			r := repository.NewStudentRepository(db)
			ss := services.NewStripeService(conf, tk, sc)
			s := services.NewStudentService(r, ss)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("", CreateStudent)
	g.PUT("/:id", UpdateStudent)
	g.DELETE("/:id", DeleteStudent)
	g.PUT("/:id/restore", RestoreStudent)
	g.GET("/:id", GetStudentByID)
	g.GET("", GetStudentList)
	g.POST("/charge", ChargeWithID)
}

type CreateStudentBaseCallbackFunc func(services.StudentService, *models.Student) (*models.Student, error)

func CreateStudentBase(c echo.Context, params interface{}, callback CreateStudentBaseCallbackFunc) (err error) {
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

func CreateStudent(c echo.Context) (err error) {
	return CreateStudentBase(c, nil, func(service services.StudentService, data *models.Student) (*models.Student, error) {
		return service.Create(data)
	})
}

type UpdateStudentBaseCallbackFunc func(services.StudentService, uint, *models.Student) (*models.Student, error)

func UpdateStudentBase(c echo.Context, params interface{}, callback UpdateStudentBaseCallbackFunc) (err error) {
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

func UpdateStudent(c echo.Context) (err error) {
	return UpdateStudentBase(c, nil, func(service services.StudentService, id uint, data *models.Student) (*models.Student, error) {
		return service.Update(uint(id), data)
	})
}

type DeleteStudentBaseCallbackFunc func(services.StudentService, uint) (*models.Student, error)

func DeleteStudentBase(c echo.Context, params interface{}, callback DeleteStudentBaseCallbackFunc) (err error) {
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

func DeleteStudent(c echo.Context) (err error) {
	var param struct {
		HardDelete bool
	}
	return DeleteStudentBase(c, &param, func(service services.StudentService, id uint) (*models.Student, error) {
		if param.HardDelete {
			return service.HardDelete(id)
		}
		return service.SoftDelete(id)
	})
}

type RestoreStudentBaseCallbackFunc func(services.StudentService, uint) (*models.Student, error)

func RestoreStudentBase(c echo.Context, params interface{}, callback RestoreStudentBaseCallbackFunc) (err error) {
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

func RestoreStudent(c echo.Context) (err error) {
	return RestoreStudentBase(c, nil, func(service services.StudentService, id uint) (*models.Student, error) {
		return service.Restore(id)
	})
}

type GetStudentByIDBaseCallbackFunc func(services.StudentService, uint) (*models.Student, error)

func GetStudentByIDBase(c echo.Context, params interface{}, callback GetStudentByIDBaseCallbackFunc) (err error) {
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

func GetStudentByID(c echo.Context) (err error) {
	var param struct {
		Expand []string
	}
	return GetStudentByIDBase(c, &param, func(service services.StudentService, id uint) (*models.Student, error) {
		return service.GetByID(id, param.Expand...)
	})
}

type GetStudentListResponse struct {
	AllCount uint `json:",omitempty"`
	Offset   uint `json:",omitempty"`
	Limit    uint `json:",omitempty"`
	Data     []*models.Student
}

type GetStudentListBaseCallbackFunc func(services.StudentService) (*GetStudentListResponse, error)

func GetStudentListBase(c echo.Context, params interface{}, callback GetStudentListBaseCallbackFunc) (err error) {
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

func GetStudentList(c echo.Context) (err error) {
	var param struct {
		services.GetAllConfig
	}

	return GetStudentListBase(c, &param, func(service services.StudentService) (*GetStudentListResponse, error) {
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

func ChargeWithID(c echo.Context) error {
	services := c.Get("Service").(services.StudentService)
	params := &types.ChargeWithIDParams{}
	if err := c.Bind(params); err != nil {
		return errors.New(err.Error())
	}
	if err := c.Validate(params); err != nil {
		return err
	}
	charge, err := services.ChargeWithID(params)
	if err != nil {
		return nil
	}
	return c.JSON(http.StatusOK, charge)
}
