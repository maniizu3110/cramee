package middleware

import (
	"cramee/myerror"
	"cramee/token"
	"errors"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// AuthMiddleware creates a middleware for authorization
func AuthMiddleware(tokenMaker token.Maker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorizationHeader := c.Request().Header.Get("Authorization")
			if len(authorizationHeader) == 0 {
				err := errors.New("authorization header is not provided")
				if err != nil {
					return myerror.NewPublic(myerror.ErrEmptyAuthorization, err)
				}
			}
			fields := strings.Fields(authorizationHeader)
			if len(fields) < 2 {
				err := errors.New("invalid authorization header format")
				if err != nil {
					return myerror.NewPublic(myerror.ErrInvalidAuthorization, err)
				}
			}

			authorizationType := strings.ToLower(fields[0])
			if authorizationType != authorizationTypeBearer {
				err := fmt.Errorf("unsupported authorization type %s", authorizationType)
				if err != nil {
					return myerror.NewPublic(myerror.ErrInvalidTypeAuthorization, err)
				}
			}

			accessToken := fields[1]
			payload, err := tokenMaker.VerifyToken(accessToken)
			if err != nil {
				return myerror.NewPublic(myerror.ErrVerifyToken, err)
			}
			c.Set("token", payload)
			return next(c)
		}
	}
}
