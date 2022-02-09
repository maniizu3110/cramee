package handler

import (
	"cramee/api/services"
	"cramee/myerror"
	"cramee/token"
	"cramee/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AssignSignStudentHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			conf := c.Get("config").(util.Config)
			tk := c.Get("tk").(token.Maker)
			s := services.NewSignStudentService(conf, tk)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("", CreateStudentHandler)
	g.POST("/login", LoginStudentHandler)
}

func CreateStudentHandler(c echo.Context) error {
	services := c.Get("Service").(services.SignStudentService)
	params := &services.CreateStudentParams{}
	if err := c.Bind(params); err != nil {
		return myerror.NewPublic(myerror.ErrBindData, err)
	}
	if err := c.Validate(params); err != nil {
		return myerror.New(myerror.ErrRequestData, err, err)
	}
	student, err := services.CreateStudent(params)
	if err != nil {
		return myerror.NewPublic(myerror.ErrCreate, err)
	}
	return c.JSON(http.StatusOK, student)
}

func LoginStudentHandler(c echo.Context) error {
	services := c.Get("Service").(services.SignStudentService)
	params := &services.LoginStudentParams{}
	if err := c.Bind(params); err != nil {
		return myerror.NewPublic(myerror.ErrBindData, err)
	}
	if err := c.Validate(params); err != nil {
		return myerror.New(myerror.ErrRequestData, err, err)
	}
	res, err := services.LoginStudent(params)
	if err != nil {
		return myerror.NewPublic(myerror.ErrLogin, err)
	}
	return c.JSON(http.StatusOK, res)
}