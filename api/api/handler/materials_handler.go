package handler

import (
	"cramee/api/services"
	"cramee/myerror"
	"cramee/token"
	"cramee/util"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AssignMarterialsHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			conf := c.Get("config").(util.Config)
			tk := c.Get("tk").(token.Maker)
			s := services.NewMaterialsService(conf, tk)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("/upload", UploadMaterials)
	g.GET("/get-url", GetUrlOfMaterials)
}

func UploadMaterials(c echo.Context) error {
	services := c.Get("Service").(services.MaterialsService)
	var params struct {
		status      string
		studentId   string
		teacherId   string
		schoolHours string
		file        multipart.FileHeader
	}
	if err := c.Bind(&params); err != nil {
		return myerror.NewPublic(myerror.ErrBindData, err)
	}
	err := services.UploadMaterials(params.file, params.status, params.teacherId, params.studentId, params.schoolHours)
	if err != nil {
		return myerror.NewPublic(myerror.ErrUpload, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func GetUrlOfMaterials(c echo.Context) error {
	services := c.Get("Service").(services.MaterialsService)
	var params struct {
		status      string
		studentId   string
		teacherId   string
		schoolHours string
	}
	if err := c.Bind(&params); err != nil {
		return myerror.NewPublic(myerror.ErrBindData, err)
	}
	urlStr, err := services.GetUrlOfMarterials(params.status, params.teacherId, params.studentId, params.schoolHours)
	if err != nil {
		return myerror.NewPublic(myerror.ErrGetData, err)
	}
	return c.JSON(http.StatusOK, urlStr)
}
