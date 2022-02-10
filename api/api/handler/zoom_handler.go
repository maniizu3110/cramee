package handler

import (
	"cramee/api/services"
	"cramee/token"
	"cramee/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AssignZoomHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			conf := c.Get("config").(util.Config)
			tk := c.Get("tk").(token.Maker)
			s := services.NewZoomService(conf, tk)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.GET("/create-zoom-meeting", CreateZoomMeeting)
}

func CreateZoomMeeting(c echo.Context) error {

	return c.JSON(http.StatusOK, nil) // TODO:ここ変更
}