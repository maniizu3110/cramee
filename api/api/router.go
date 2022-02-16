package api

import (
	"cramee/api/handler"
	"cramee/api/middleware"
	"cramee/myerror"
	"cramee/util"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (server *Server) SetRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggingMiddleware, myerror.HandleErrorMiddleware())
	validator, err := util.NewValidator()
	if err != nil {
		logrus.Fatal("バリデージョンの設定に失敗しました")
	}
	if server.config.Env != "prod" {
		e.Debug = true
	}
	e.Validator = validator
	middleware.CORS(e)
	{
		// 認証不要
		g := e.Group("/v1",
			middleware.SetDB(server.db),
			middleware.SetConfig(server.config),
			middleware.SetTokenMaker(server.tokenMaker),
			middleware.SetStripeClient(server.config),
		)
		handler.AssignSignStudentHandler(g.Group("/sign-student"))
		handler.AssignSignTeacherHandler(g.Group("/sign-teacher"))
	}
	{
		// 要認証
		g := e.Group("/v1",
			middleware.SetDB(server.db),
			middleware.SetConfig(server.config),
			middleware.SetTokenMaker(server.tokenMaker),
			middleware.SetStripeClient(server.config),
			middleware.AuthMiddleware(server.tokenMaker),
		)
		handler.AssignStudentHandler(g.Group("/student"))
		handler.AssignTeacherHandler(g.Group("/teacher"))
		handler.AssignZoomHandler(g.Group("/zoom"))
		handler.AssignStripeHandler(g.Group("/stripe"))
		handler.AssignLectureHandler(g.Group("/lecture"))
	}
	return e
}
