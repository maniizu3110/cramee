package services

import (
	"cramee/lib/zoom"
	"cramee/token"
	"cramee/util"
)

//go:generate mockgen -source=$GOFILE -destination=${GOPACKAGE}_mock/${GOFILE}.mock.go -package=${GOPACKAGE}_mock
type ZoomService interface {
	ListUsers(opts zoom.ListUsersOptions) (zoom.ListUsersResponse, error)
	CreateMeeting(opts zoom.CreateMeetingOptions) (zoom.Meeting, error)
	CreateUser(opts zoom.CreateUserOptions) (zoom.User, error)
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
func (z *zoomServiceImpl) CreateMeeting(opts zoom.CreateMeetingOptions) (zoom.Meeting, error) {
	res, err := z.client.CreateMeeting(opts)
	if err != nil {
		return zoom.Meeting{}, err
	}
	return res, nil
}
func (z *zoomServiceImpl) CreateUser(opts zoom.CreateUserOptions) (zoom.User, error) {
	res, err := z.client.CreateUser(opts)
	if err != nil {
		return zoom.User{}, err
	}
	return res, nil
}
