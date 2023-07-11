package models

import "errors"

var (
	ErrNoHaveAccessCookie   = errors.New("access token is empty")
	ErrBadParseAccessCookie = errors.New("access token is not valid")
	ErrorTokenTime          = errors.New("token is expired")
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenValues struct {
	UserId int ` json:"user_id"`
}
