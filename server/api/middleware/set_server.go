package middleware

import (
	"cramee/token"
	"cramee/util"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func SetDB(db *sql.DB) echo.MiddlewareFunc {
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
