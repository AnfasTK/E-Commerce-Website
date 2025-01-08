package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId uint
	Email  string `json:"username"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

var JwtSecretKey = []byte(os.Getenv("SECRETKEY"))

func GenerateJWT(userId uint, email string, role string) (string, error) {
	claims:=Claims{
		UserId: uint(userId),
		Email: email,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err:=token.SignedString(JwtSecretKey)
	if err != nil{
		return "",err
	}
	return tokenString,nil
}
