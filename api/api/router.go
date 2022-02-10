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
		)
		handler.AssignSignStudentHandler(g.Group("/sign-student"))
		handler.AssignSignTeacherHandler(g.Group("/sign-teacher"))
	}
	{
		// 認証必要
		g := e.Group("/v1",
			middleware.SetDB(server.db),
			middleware.SetConfig(server.config),
			middleware.AuthMiddleware(server.tokenMaker),
		)
		handler.AssignStudentHandler(g.Group("/student"))
		handler.AssignTeacherHandler(g.Group("/student"))
		handler.AssignZoomHandler(g.Group("/zoom"))
		handler.AssignTeacherLectureScheduleHandler(g.Group("/teacher-lecture-schedule"))
	}
	return e
}
