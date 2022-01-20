package myerror

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	IsError bool
	Hash    string
	Message interface{}
	Detail  interface{}
}

func HandleErrorMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				if c.Response().Committed {
					logrus.Errorln("error returned, but response already commited", err)
					return nil
				}
				switch err := err.(type) {
				case MyError:
					message, detail := err.GetPublicMessage()
					c.Response().Header().Set("error-hash", err.GetHash())
					logrus.Errorln(err)
					return c.JSON(err.GetHTTPCode(), ErrorResponse{
						IsError: true,
						Hash:    err.GetHash(),
						Message: message,
						Detail:  detail,
					})
				case *echo.HTTPError:
					c.Response().Header().Set("error-hash", "none")
					logrus.Errorln(err)
					return c.JSON(err.Code, ErrorResponse{
						IsError: true,
						Hash:    "",
						Message: err.Message,
						Detail:  "",
					})
				default:
					message := "エラーが発生しました"
					c.Response().Header().Set("error-hash", "none")
					logrus.Errorln("unknown error", err)
					return c.JSON(http.StatusInternalServerError, ErrorResponse{
						IsError: true,
						Hash:    "",
						Message: message,
						Detail:  "",
					})
				}
			}
			return nil
		}
	}
}
