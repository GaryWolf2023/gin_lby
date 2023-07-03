package utils

import (
	"Gin_one/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var stSignKey = []byte(viper.GetString("jwt.signKey"))

type JwtCustClaims struct {
	ID   int
	Name string
	jwt.RegisteredClaims
}

func CreateToken(id int, name string) (string, error) {
	iJwtCustClaims := JwtCustClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.ViperConfig.GetDuration("jwt.tokenExpire") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustClaims)
	return token.SignedString(stSignKey)
}

func ParseToken(tokenString string) (JwtCustClaims, error) {
	iJwtCustClaims := JwtCustClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &iJwtCustClaims, func(token *jwt.Token) (interface{}, error) {
		return stSignKey, nil
	})
	if err == nil && !token.Valid {
		err = errors.New("token is invalid")
	}
	return iJwtCustClaims, err
}

func IsTokenValid(tokenString string) bool {
	_, err := ParseToken(tokenString)
	return err == nil
}
