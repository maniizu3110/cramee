package services

import (
	"cramee/myerror"
	"cramee/token"
	"cramee/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

//go:generate mockgen -source=$GOFILE -destination=${GOPACKAGE}_mock/${GOFILE}.mock.go -package=${GOPACKAGE}_mock
type ZoomService interface {
	CreateMeeting(sh int, zoomEmail string) (string, error)
}

type ZoomServiceImpl struct {
	config     util.Config
	tokenMaker token.Maker
}

func NewZoomService(config util.Config, tokenMaker token.Maker) ZoomService {
	res := &ZoomServiceImpl{}
	res.config = config
	res.tokenMaker = tokenMaker
	return res
}

func (z *ZoomServiceImpl) CreateMeeting(sh int, zoomEmail string) (string, error) {
	ZOOM_API_KEY := z.config.ZoomApiKey
	ZOOM_API_SECRET := z.config.ZoomApiSecret

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": ZOOM_API_KEY,
		"exp": fmt.Sprintf("%d", time.Now().Add(time.Duration(2*sh)*time.Minute).Unix()),
	})

	tokenString, err := token.SignedString([]byte(ZOOM_API_SECRET))
	if err != nil {
		return "", err
	}

	USER_ID := zoomEmail
	CREATE_MEETING_URL := "https://api.zoom.us/v2/users/" + USER_ID + "/meetings"
	JWT := tokenString

	type Meeting struct {
		UUID      string    `json:"uuid"`
		ID        int       `json:"id"`
		HostID    string    `json:"host_id"`
		Topic     string    `json:"topic"`
		Type      int       `json:"type"`
		Duration  int       `json:"duration"`
		Timezone  string    `json:"timezone"`
		CreatedAt time.Time `json:"created_at"`
		StartURL  string    `json:"start_url"`
		JoinURL   string    `json:"join_url"`
		Settings  struct {
			HostVideo           bool   `json:"host_video"`
			ParticipantVideo    bool   `json:"participant_video"`
			CnMeeting           bool   `json:"cn_meeting"`
			InMeeting           bool   `json:"in_meeting"`
			JoinBeforeHost      bool   `json:"join_before_host"`
			MuteUponEntry       bool   `json:"mute_upon_entry"`
			Watermark           bool   `json:"watermark"`
			UsePmi              bool   `json:"use_pmi"`
			ApprovalType        int    `json:"approval_type"`
			Audio               string `json:"audio"`
			AutoRecording       string `json:"auto_recording"`
			EnforceLogin        bool   `json:"enforce_login"`
			EnforceLoginDomains string `json:"enforce_login_domains"`
			AlternativeHosts    string `json:"alternative_hosts"`
		} `json:"settings"`
	}

	payload := strings.NewReader(`{"type":1}`)

	req, err := http.NewRequest("POST", CREATE_MEETING_URL, payload)
	if err != nil {
		return "", myerror.NewPublic(myerror.ErrCreate, err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+JWT)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", myerror.NewPublic(myerror.ErrCreate, err)
	}

	if res.StatusCode != 201 {
		return "", myerror.NewPublic(myerror.ErrInvalidAuthorization, err)
	}
	defer res.Body.Close()

	var meeting Meeting
	err = json.NewDecoder(res.Body).Decode(&meeting)
	if err != nil {
		return "", myerror.NewPublic(myerror.ErrCreate, err)
	}
	return meeting.JoinURL, nil
}