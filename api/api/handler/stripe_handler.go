package handler

import (
	"cramee/api/services"
	"cramee/token"
	"cramee/util"

	"github.com/labstack/echo/v4"
)

func AssignStripeHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			conf := c.Get("config").(util.Config)
			tk := c.Get("tk").(token.Maker)
			s := services.NewStripeService(conf, tk)
			c.Set("Service", s)
			return handler(c)
		}
	})
}
