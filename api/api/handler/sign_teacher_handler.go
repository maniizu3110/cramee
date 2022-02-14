package handler

import (
	"cramee/api/repository"
	"cramee/api/services"
	"cramee/api/services/types"
	"cramee/lib/zoom"
	"cramee/models"
	"cramee/myerror"
	"cramee/token"
	"cramee/util"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AssignSignTeacherHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			conf := c.Get("config").(util.Config)
			tk := c.Get("tk").(token.Maker)
			db := c.Get("Tx").(*gorm.DB)
			zc := zoom.NewClient(conf.ZoomApiKey, conf.ZoomApiSecret)
			r := repository.NewTeacherRepository(db)
			zs := services.NewZoomService(conf, tk, zc)
			s := services.NewSignTeacherService(r, conf, tk, zs)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("", CreateTeacherHandler)
	g.POST("/with-zoom", CreateTeacherWithZoomHandler)
	g.POST("/login", LoginTeacherHandler)
}

func CreateTeacherHandler(c echo.Context) error {
	services := c.Get("Service").(services.SignTeacherService)
	params := &models.Teacher{}
	if err := c.Bind(params); err != nil {
		return myerror.NewPublic(myerror.ErrBindData, err)
	}
	if err := c.Validate(params); err != nil {
		return myerror.New(myerror.ErrRequestData, err, err)
	}
	teacher, err := services.CreateTeacher(params)
	if err != nil {
		return myerror.NewPublic(myerror.ErrCreate, err)
	}
	return c.JSON(http.StatusOK, teacher)
}

func CreateTeacherWithZoomHandler(c echo.Context) error {
	services := c.Get("Service").(services.SignTeacherService)
	params := &models.Teacher{}
	if err := c.Bind(params); err != nil {
		return myerror.NewPublic(myerror.ErrBindData, err)
	}
	if err := c.Validate(params); err != nil {
		return myerror.New(myerror.ErrRequestData, err, err)
	}
	teacher, err := services.CreateTeacher(params)
	if err != nil {
		return myerror.NewPublic(myerror.ErrCreate, err)
	}
	return c.JSON(http.StatusOK, teacher)
}

func LoginTeacherHandler(c echo.Context) error {
	services := c.Get("Service").(services.SignTeacherService)
	params := &types.LoginTeacherRequest{}
	if err := c.Bind(params); err != nil {
		return myerror.NewPublic(myerror.ErrBindData, err)
	}
	if err := c.Validate(params); err != nil {
		return myerror.New(myerror.ErrRequestData, err, err)
	}
	res, err := services.LoginTeacher(params)
	if err != nil {
		return myerror.NewPublic(myerror.ErrLogin, err)
	}
	return c.JSON(http.StatusOK, res)
}
