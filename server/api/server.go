package api

import (
	"cramee/token"
	"cramee/util"
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	config     util.Config
	tokenMaker token.Maker
	db         *gorm.DB
	router     *echo.Echo
}

func NewServer(db *gorm.DB, config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
		db:         db,
	}
	server.router = server.SetRouter()
	return server, nil

}
func (server *Server) Start(address string) {
	server.router.Logger.Fatal(server.router.Start(address))
}
