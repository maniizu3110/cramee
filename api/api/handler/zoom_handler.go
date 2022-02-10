package handler

import (
	"cramee/api/services"
	"cramee/token"
	"cramee/util"
	"cramee/zoom"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AssignZoomHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			conf := c.Get("config").(util.Config)
			tk := c.Get("tk").(token.Maker)
			zc := zoom.NewClient(conf.ZoomApiKey, conf.ZoomApiSecret)
			s := services.NewZoomService(conf, tk, zc)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("/create-meeting", CreateZoomMeeting)
	g.GET("/users", GetUsers)
}

func GetUsers(c echo.Context) error {
	service := c.Get("Service").(services.ZoomService)
	params := &zoom.ListUsersOptions{}
	if err := c.Bind(params); err != nil {
		return err
	}
	users, err := service.ListUsers(*params)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func CreateZoomMeeting(c echo.Context) error {
	service := c.Get("Service").(services.ZoomService)
	params := &zoom.CreateMeetingOptions{}
	if err := c.Bind(params); err != nil {
		return err
	}
	url, err := service.CreateMeeting(*params)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, url)
}
