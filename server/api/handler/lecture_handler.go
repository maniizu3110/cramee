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

func AssignLectureHandlers(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db := c.Get("Tx").(*gorm.DB)
			r := repository.NewLectureRepository(db)
			s := services.NewLectureService(r)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("", CreateLectureHandler)
	g.PUT("/:id", UpdateLectureHandler)
	g.DELETE("/:id", DeleteLectureHandler)
	g.PUT("/:id/restore", RestoreLectureHandler)
	g.GET("/:id", GetLectureByIDHandler)
	g.GET("", GetLectureListHandler)
}

type CreateLectureHandlerBaseCallbackFunc func(services.LectureService, *models.Lecture) (*models.Lecture, error)

func CreateLectureHandlerBase(c echo.Context, params interface{}, callback CreateLectureHandlerBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.LectureService)

	data := &models.Lecture{}
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

func CreateLectureHandler(c echo.Context) (err error) {
	return CreateLectureHandlerBase(c, nil, func(service services.LectureService, data *models.Lecture) (*models.Lecture, error) {
		return service.Create(data)
	})
}

type UpdateLectureHandlerBaseCallbackFunc func(services.LectureService, uint, *models.Lecture) (*models.Lecture, error)

func UpdateLectureHandlerBase(c echo.Context, params interface{}, callback UpdateLectureHandlerBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.LectureService)

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

func UpdateLectureHandler(c echo.Context) (err error) {
	return UpdateLectureHandlerBase(c, nil, func(service services.LectureService, id uint, data *models.Lecture) (*models.Lecture, error) {
		return service.Update(uint(id), data)
	})
}

type DeleteLectureHandlerBaseCallbackFunc func(services.LectureService, uint) (*models.Lecture, error)

func DeleteLectureHandlerBase(c echo.Context, params interface{}, callback DeleteLectureHandlerBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.LectureService)

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

func DeleteLectureHandler(c echo.Context) (err error) {
	var param struct {
		HardDelete bool
	}
	return DeleteLectureHandlerBase(c, &param, func(service services.LectureService, id uint) (*models.Lecture, error) {
		if param.HardDelete {
			return service.HardDelete(id)
		}
		return service.SoftDelete(id)
	})
}

type RestoreLectureHandlerBaseCallbackFunc func(services.LectureService, uint) (*models.Lecture, error)

func RestoreLectureHandlerBase(c echo.Context, params interface{}, callback RestoreLectureHandlerBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.LectureService)

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

func RestoreLectureHandler(c echo.Context) (err error) {
	return RestoreLectureHandlerBase(c, nil, func(service services.LectureService, id uint) (*models.Lecture, error) {
		return service.Restore(id)
	})
}

type GetLectureByIDHandlerBaseCallbackFunc func(services.LectureService, uint) (*models.Lecture, error)

func GetLectureByIDHandlerBase(c echo.Context, params interface{}, callback GetLectureByIDHandlerBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.LectureService)

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

func GetLectureByIDHandler(c echo.Context) (err error) {
	var param struct {
		Expand []string
	}
	return GetLectureByIDHandlerBase(c, &param, func(service services.LectureService, id uint) (*models.Lecture, error) {
		return service.GetByID(id, param.Expand...)
	})
}

type GetLectureListResponse struct {
	AllCount uint `json:",omitempty"`
	Offset   uint `json:",omitempty"`
	Limit    uint `json:",omitempty"`
	Data     []*models.Lecture
}

type GetLectureListHandlerBaseCallbackFunc func(services.LectureService) (*GetLectureListResponse, error)

func GetLectureListHandlerBase(c echo.Context, params interface{}, callback GetLectureListHandlerBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.LectureService)
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

func GetLectureListHandler(c echo.Context) (err error) {
	var param struct {
		services.GetAllConfig
	}

	return GetLectureListHandlerBase(c, &param, func(service services.LectureService) (*GetLectureListResponse, error) {
		m, count, err := service.GetAll(param.GetAllConfig)
		if err != nil {
			return nil, err
		}
		return &GetLectureListResponse{
			AllCount: count,
			Limit:    param.Limit,
			Offset:   param.Offset,
			Data:     m,
		}, nil
	})
}
