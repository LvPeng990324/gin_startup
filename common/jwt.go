package common

import (
	"contract_system_server/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("a_secret_create")

type Claims struct {
	UserId uint
	jwt.RegisteredClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // 7天过期
	claims := &Claims{
		UserId: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // 过期时间
			Issuer:    "contract_system_server",           // 签发人
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey) // 签名

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
