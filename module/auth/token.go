package auth

import (
	"fiberun/module/conf"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Guard    string `json:"guard"`
	jwt.RegisteredClaims
}

func GenerateToken(username, guard string) (string, error) {
	var secret = []byte(conf.Get("jwt.secret"))
	var timing = time.Duration(conf.GetInt("jwt.expire"))
	nowTime := jwt.NewNumericDate(time.Now())
	expireTime := jwt.NewNumericDate(time.Now().Add(time.Hour * timing))
	claims := Claims{
		Username: username,
		Guard:    guard,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  nowTime,
			ExpiresAt: expireTime,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenClaims.SignedString(secret)
}
