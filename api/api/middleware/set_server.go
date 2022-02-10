package middleware

import (
	"cramee/token"
	"cramee/util"

	"github.com/stripe/stripe-go/v72/client"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetDB(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("Tx", db)
			return next(c)
		}
	}
}
func SetConfig(config util.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("config", config)
			return next(c)
		}
	}
}
func SetTokenMaker(tokenMaker token.Maker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("tk", tokenMaker)
			return next(c)
		}
	}
}
func SetStripeClient(config util.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sc := &client.API{}
			sc.Init(config.StripeSecretKey, nil)
			c.Set("sc", sc)
			return next(c)
		}
	}
}
