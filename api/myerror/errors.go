package myerror

import (
	"errors"
	"net/http"
)

var ErrRequestData = errors.New("送信したデータに誤りがあります")
var ErrBindData = errors.New("データの読み取りに失敗しました")
var ErrCreate = errors.New("作成に失敗しました")
var ErrGet = errors.New("検索に失敗しました")
var ErrGetByID = errors.New("IDを用いた検索に失敗しました")
var ErrLogin = errors.New("ログインに失敗しました")
var ErrChangePasswordToHash = errors.New("パスワードのハッシュ化に失敗しました")
var ErrCheckPassword = errors.New("パスワードが違います")
var ErrEmptyAuthorization = errors.New("認証情報が空です")
var ErrInvalidAuthorization = errors.New("無効な認証情報です")
var ErrInvalidTypeAuthorization = errors.New("トークンタイプが無効です")
var ErrVerifyToken = errors.New("トークンの検証に失敗しました")

func NewErrorStatusCodeMaps() map[error]int {
	errorStatusCodeMaps := make(map[error]int)
	errorStatusCodeMaps[ErrRequestData] = http.StatusBadRequest
	errorStatusCodeMaps[ErrBindData] = http.StatusBadRequest
	errorStatusCodeMaps[ErrCreate] = http.StatusInternalServerError
	errorStatusCodeMaps[ErrGet] = http.StatusInternalServerError
	errorStatusCodeMaps[ErrGetByID] = http.StatusInternalServerError
	errorStatusCodeMaps[ErrLogin] = http.StatusInternalServerError
	errorStatusCodeMaps[ErrChangePasswordToHash] = http.StatusInternalServerError
	errorStatusCodeMaps[ErrCheckPassword] = http.StatusInternalServerError
	errorStatusCodeMaps[ErrEmptyAuthorization] = http.StatusUnauthorized
	errorStatusCodeMaps[ErrInvalidAuthorization] = http.StatusUnauthorized
	errorStatusCodeMaps[ErrInvalidTypeAuthorization] = http.StatusUnauthorized
	errorStatusCodeMaps[ErrVerifyToken] = http.StatusUnauthorized
	return errorStatusCodeMaps
}

func GetCodeFromError(err error) int {
	statusCodes := NewErrorStatusCodeMaps()
	for key, value := range statusCodes {
		if errors.Is(err, key) {
			return value
		}
	}
	return http.StatusInternalServerError
}
