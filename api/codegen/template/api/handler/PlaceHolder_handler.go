package handler

import (
	"cramee/codegen/template/api/repository"
	"cramee/codegen/template/api/services"
	"cramee/codegen/template/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AssignPlaceHolderHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db := c.Get("Tx").(*gorm.DB)
			r := repository.NewPlaceHolderRepository(db)
			s := services.NewPlaceHolderService(r)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("", CreatePlaceHolder)
	g.PUT("/:id", UpdatePlaceHolder)
	g.DELETE("/:id", DeletePlaceHolder)
	g.PUT("/:id/restore", RestorePlaceHolder)
	g.GET("/:id", GetPlaceHolderByID)
	g.GET("", GetPlaceHolderList)
}

type CreatePlaceHolderBaseCallbackFunc func(services.PlaceHolderService, *models.PlaceHolder) (*models.PlaceHolder, error)

func CreatePlaceHolderBase(c echo.Context, params interface{}, callback CreatePlaceHolderBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.PlaceHolderService)

	data := &models.PlaceHolder{}
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

func CreatePlaceHolder(c echo.Context) (err error) {
	return CreatePlaceHolderBase(c, nil, func(service services.PlaceHolderService, data *models.PlaceHolder) (*models.PlaceHolder, error) {
		return service.Create(data)
	})
}

type UpdatePlaceHolderBaseCallbackFunc func(services.PlaceHolderService, uint, *models.PlaceHolder) (*models.PlaceHolder, error)

func UpdatePlaceHolderBase(c echo.Context, params interface{}, callback UpdatePlaceHolderBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.PlaceHolderService)

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

func UpdatePlaceHolder(c echo.Context) (err error) {
	return UpdatePlaceHolderBase(c, nil, func(service services.PlaceHolderService, id uint, data *models.PlaceHolder) (*models.PlaceHolder, error) {
		return service.Update(uint(id), data)
	})
}

type DeletePlaceHolderBaseCallbackFunc func(services.PlaceHolderService, uint) (*models.PlaceHolder, error)

func DeletePlaceHolderBase(c echo.Context, params interface{}, callback DeletePlaceHolderBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.PlaceHolderService)

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

func DeletePlaceHolder(c echo.Context) (err error) {
	var param struct {
		HardDelete bool
	}
	return DeletePlaceHolderBase(c, &param, func(service services.PlaceHolderService, id uint) (*models.PlaceHolder, error) {
		if param.HardDelete {
			return service.HardDelete(id)
		}
		return service.SoftDelete(id)
	})
}

type RestorePlaceHolderBaseCallbackFunc func(services.PlaceHolderService, uint) (*models.PlaceHolder, error)

func RestorePlaceHolderBase(c echo.Context, params interface{}, callback RestorePlaceHolderBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.PlaceHolderService)

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

func RestorePlaceHolder(c echo.Context) (err error) {
	return RestorePlaceHolderBase(c, nil, func(service services.PlaceHolderService, id uint) (*models.PlaceHolder, error) {
		return service.Restore(id)
	})
}

type GetPlaceHolderByIDBaseCallbackFunc func(services.PlaceHolderService, uint) (*models.PlaceHolder, error)

func GetPlaceHolderByIDBase(c echo.Context, params interface{}, callback GetPlaceHolderByIDBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.PlaceHolderService)

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

func GetPlaceHolderByID(c echo.Context) (err error) {
	var param struct {
		Expand []string
	}
	return GetPlaceHolderByIDBase(c, &param, func(service services.PlaceHolderService, id uint) (*models.PlaceHolder, error) {
		return service.GetByID(id, param.Expand...)
	})
}

type GetPlaceHolderListResponse struct {
	AllCount uint `json:",omitempty"`
	Offset   uint `json:",omitempty"`
	Limit    uint `json:",omitempty"`
	Data     []*models.PlaceHolder
}

type GetPlaceHolderListBaseCallbackFunc func(services.PlaceHolderService) (*GetPlaceHolderListResponse, error)

func GetPlaceHolderListBase(c echo.Context, params interface{}, callback GetPlaceHolderListBaseCallbackFunc) (err error) {
	service := c.Get("Service").(services.PlaceHolderService)
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

func GetPlaceHolderList(c echo.Context) (err error) {
	var param struct {
		services.GetAllConfig
	}

	return GetPlaceHolderListBase(c, &param, func(service services.PlaceHolderService) (*GetPlaceHolderListResponse, error) {
		m, count, err := service.GetAll(param.GetAllConfig)
		if err != nil {
			return nil, err
		}
		return &GetPlaceHolderListResponse{
			AllCount: count,
			Limit:    param.Limit,
			Offset:   param.Offset,
			Data:     m,
		}, nil
	})
}
