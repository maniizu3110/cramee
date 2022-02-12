package handler

import (
	"cramee/api/services"
	"cramee/token"
	"cramee/util"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
)

func AssignStripeHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			conf := c.Get("config").(util.Config)
			tk := c.Get("tk").(token.Maker)
			sc := c.Get("sc").(*client.API)

			s := services.NewStripeService(conf, tk, sc)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("/customer", CreateCustomer)
	g.POST("/card", CreateCard)
	//TODO:セキュリティのためHTTPで実行できないようにする
	g.POST("/charge", Charge)

}
func CreateCustomer(c echo.Context) error {
	services := c.Get("Service").(services.StripeService)
	params := &stripe.CustomerParams{}
	if err := c.Bind(params); err != nil {
		return errors.New(err.Error())
	}
	if err := c.Validate(params); err != nil {
		return err
	}
	client, err := services.CreateCustomer(params)
	if err != nil {
		return nil
	}
	return c.JSON(http.StatusOK, client)
}
func CreateCard(c echo.Context) error {
	services := c.Get("Service").(services.StripeService)
	params := &stripe.CardParams{}
	if err := c.Bind(params); err != nil {
		return errors.New(err.Error())
	}
	if err := c.Validate(params); err != nil {
		return err
	}
	client, err := services.CreateCard(params)
	if err != nil {
		return nil
	}
	return c.JSON(http.StatusOK, client)
}
func Charge(c echo.Context) error {
	services := c.Get("Service").(services.StripeService)
	params := &stripe.ChargeParams{}
	if err := c.Bind(params); err != nil {
		return errors.New(err.Error())
	}
	if err := c.Validate(params); err != nil {
		return err
	}
	client, err := services.Charge(params)
	if err != nil {
		return nil
	}
	return c.JSON(http.StatusOK, client)
}
