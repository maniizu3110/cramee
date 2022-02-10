package services

import (
	"cramee/token"
	"cramee/util"
	"cramee/zoom"
)

//go:generate mockgen -source=$GOFILE -destination=${GOPACKAGE}_mock/${GOFILE}.mock.go -package=${GOPACKAGE}_mock
type ZoomService interface {
	ListUsers(opts zoom.ListUsersOptions) (zoom.ListUsersResponse, error)
}

type zoomServiceImpl struct {
	config     util.Config
	tokenMaker token.Maker
	client     *zoom.Client
}

func NewZoomService(config util.Config, tokenMaker token.Maker, client *zoom.Client) ZoomService {
	res := &zoomServiceImpl{}
	res.config = config
	res.tokenMaker = tokenMaker
	res.client = client
	return res
}

func (z *zoomServiceImpl) ListUsers(opts zoom.ListUsersOptions) (zoom.ListUsersResponse, error) {
	res, err := z.client.ListUsers(opts)
	if err != nil {
		return zoom.ListUsersResponse{}, err
	}
	return res, nil
}
