package handler

import (
	"cramee/api/repository"
	"cramee/api/services"
	"cramee/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AssignLectureHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db := c.Get("Tx").(*gorm.DB)
			r := repository.NewLectureRepository(db)
			s := services.NewLectureService(r)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("", CreateLecture)
	g.PUT("/:id", UpdateLecture)
	g.DELETE("/:id", DeleteLecture)
	g.PUT("/:id/restore", RestoreLecture)
	g.GET("/:id", GetLectureByID)
	g.GET("", GetLectureList)
}

type CreateLectureBaseCallbackFunc func(services.LectureService, *models.Lecture) (*models.Lecture, error)

func CreateLectureBase(c echo.Context, params interface{}, callback CreateLectureBaseCallbackFunc) (err error) {
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

func CreateLecture(c echo.Context) (err error) {
	return CreateLectureBase(c, nil, func(service services.LectureService, data *models.Lecture) (*models.Lecture, error) {
		return service.Create(data)
	})
}

type UpdateLectureBaseCallbackFunc func(services.LectureService, uint, *models.Lecture) (*models.Lecture, error)

func UpdateLectureBase(c echo.Context, params interface{}, callback UpdateLectureBaseCallbackFunc) (err error) {
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

func UpdateLecture(c echo.Context) (err error) {
	return UpdateLectureBase(c, nil, func(service services.LectureService, id uint, data *models.Lecture) (*models.Lecture, error) {
		return service.Update(uint(id), data)
	})
}

type DeleteLectureBaseCallbackFunc func(services.LectureService, uint) (*models.Lecture, error)

func DeleteLectureBase(c echo.Context, params interface{}, callback DeleteLectureBaseCallbackFunc) (err error) {
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

func DeleteLecture(c echo.Context) (err error) {
	var param struct {
		HardDelete bool
	}
	return DeleteLectureBase(c, &param, func(service services.LectureService, id uint) (*models.Lecture, error) {
		if param.HardDelete {
			return service.HardDelete(id)
		}
		return service.SoftDelete(id)
	})
}

type RestoreLectureBaseCallbackFunc func(services.LectureService, uint) (*models.Lecture, error)

func RestoreLectureBase(c echo.Context, params interface{}, callback RestoreLectureBaseCallbackFunc) (err error) {
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

func RestoreLecture(c echo.Context) (err error) {
	return RestoreLectureBase(c, nil, func(service services.LectureService, id uint) (*models.Lecture, error) {
		return service.Restore(id)
	})
}

type GetLectureByIDBaseCallbackFunc func(services.LectureService, uint) (*models.Lecture, error)

func GetLectureByIDBase(c echo.Context, params interface{}, callback GetLectureByIDBaseCallbackFunc) (err error) {
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

func GetLectureByID(c echo.Context) (err error) {
	var param struct {
		Expand []string
	}
	return GetLectureByIDBase(c, &param, func(service services.LectureService, id uint) (*models.Lecture, error) {
		return service.GetByID(id, param.Expand...)
	})
}

type GetLectureListResponse struct {
	AllCount uint `json:",omitempty"`
	Offset   uint `json:",omitempty"`
	Limit    uint `json:",omitempty"`
	Data     []*models.Lecture
}

type GetLectureListBaseCallbackFunc func(services.LectureService) (*GetLectureListResponse, error)

func GetLectureListBase(c echo.Context, params interface{}, callback GetLectureListBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.LectureService)
	if params != nil {
		if err = c.Bind(params); err != nil {
			return errors.New(err.Error())
		}
		logrus.Info(params)
	}
	response, err := callback(service)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response)
}

func GetLectureList(c echo.Context) (err error) {
	var param struct {
		services.GetAllConfig
	}
	return GetLectureListBase(c, &param, func(service services.LectureService) (*GetLectureListResponse, error) {
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
