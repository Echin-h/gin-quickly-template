package auth

import (
	"errors"
	"gin-quickly-template/config"
	"github.com/golang-jwt/jwt"
	"time"
)

type Info struct {
	Uid  string
	Name string

	IsRefreshToken bool
}

type JWTClaims struct {
	Info Info
	jwt.StandardClaims
}

const (
	AccessTokenExpireIn  = time.Hour * 12 * 7
	RefreshTokenExpireIn = time.Hour * 24 * 30
)

// GenToken 生成JWT
func GenToken(info Info, expire ...time.Duration) (string, error) {
	if len(expire) == 0 {
		expire = append(expire, AccessTokenExpireIn)
	}
	c := JWTClaims{
		Info: info,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expire[0]).Unix(),
			Issuer:    config.GetConfig().Auth.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(config.GetConfig().Auth.Secret))
}

func genTokenByTest(info Info, expire ...time.Duration) (string, error) {
	if len(expire) == 0 {
		expire = append(expire, AccessTokenExpireIn)
	}
	c := JWTClaims{
		Info: info,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expire[0]).Unix(),
			Issuer:    "MJCLOUDS",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte("<random>"))
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(config.GetConfig().Auth.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
