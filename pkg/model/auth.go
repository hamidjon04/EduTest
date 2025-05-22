package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RegisterReq struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Password string `json:"password"`
}

type Tokens struct {
	UserId       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
}

type RefreshTokenReq struct {
	UserId       string `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}

type ClaimItems struct {
	UserId    string        `json:"user_id"`
	CreatedAt time.Duration `json:"created_at"`
}

type Claim struct {
	Items map[string]interface{} `json:"items"`
	jwt.RegisteredClaims
}