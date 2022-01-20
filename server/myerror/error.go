package myerror

import (
	"fmt"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type MyError interface {
	Error() string
	GetHash() string
	GetHTTPCode() int
	GetUnderLyingError() error
	GetPublicMessage() (message string, detail string)
}

type myError struct {
	hash            string
	HTTPCode        int
	PublicMessage   string
	InternalMessage string
	Err             error //ラップされる前のエラー
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error(%v)[%v] %v | %v | (%v)", e.hash, e.HTTPCode, e.PublicMessage, e.InternalMessage, e.Err)
}

func (e *myError) GetHash() string {
	return e.hash
}

func (e *myError) GetMergedPublicMessage() string {
	if e == nil {
		return ""
	}
	res := e.PublicMessage
	if v, ok := ToMyError(e.Err); ok {
		if len(res) == 0 {
			return v.GetMergedPublicMessage()
		}
		return res + " (" + v.GetMergedPublicMessage() + ")"
	}
	return res
}

//publicMessageの最新おのものをmessage,それ以外をdetailとして出力する
func (e *myError) GetMergedPublicMessage2() (message string, detail string) {
	if e == nil {
		return "", ""
	}
	split := strings.SplitN(e.PublicMessage, "\n", 2)
	message = split[0]
	if len(split) > 1 {
		detail = split[1]
	}
	if v, ok := ToMyError(e.Err); ok {
		if len(split) > 1 {
			detail = split[1] + "\n" + v.GetMergedPublicMessage()
		} else {
			detail = v.GetMergedPublicMessage()
		}
	}
	return
}

func (e *myError) GetMergedInternalMessage() string {
	if e == nil {
		return ""
	}
	res := e.InternalMessage
	if v, ok := ToMyError(e.Err); ok {
		return res + " (" + v.GetMergedInternalMessage() + ")"
	} else if e != nil {
		return res + " (" + e.Error() + ")"
	}
	return res
}

func (e *myError) GetHTTPCode() int {
	if e == nil {
		return 0
	}
	if e.HTTPCode != 0 {
		return e.HTTPCode
	}
	if v, ok := ToMyError(e.Err); ok {
		return v.GetHTTPCode()
	}
	return http.StatusInternalServerError
}

func (e *myError) GetUnderLyingError() error {
	return e.Err
}

func (e *myError) GetPublicMessage() (string, string) {
	return e.GetMergedPublicMessage2()
}

func ToMyError(e error) (*myError, bool) {
	v, ok := e.(*myError)
	return v, ok
}

func New(publicErr error, internalErr error, err error) *myError {
	myerror := &myError{}
	//publicErrのエラーコードがinternalErrorよりも優先されることに注意
	myerror.HTTPCode = GetCodeFromError(publicErr)
	h := uuid.NewV1()
	myerror.hash = h.String()
	myerror.PublicMessage = publicErr.Error()
	myerror.InternalMessage = internalErr.Error()
	myerror.Err = err
	return myerror
}

func NewPublic(publicErr error, err error) *myError {
	myerror := &myError{}
	//publicErrのエラーコードがinternalErrorよりも優先されることに注意
	myerror.HTTPCode = GetCodeFromError(publicErr)
	h := uuid.NewV1()
	myerror.hash = h.String()
	myerror.PublicMessage = publicErr.Error()
	myerror.Err = err
	return myerror
}

func NewInternal(internalErr error, err error) *myError {
	myerror := &myError{}
	myerror.HTTPCode = GetCodeFromError(internalErr)
	h := uuid.NewV1()
	myerror.hash = h.String()
	myerror.InternalMessage = internalErr.Error()
	myerror.Err = err
	return myerror
}

func NewReturn(err error) *myError {
	myerror := &myError{}
	myerror.Err = err
	return myerror
}
