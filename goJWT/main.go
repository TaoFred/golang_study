package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func main() {
	tokenString, err := GenRegisteredClaims()
	if err != nil {
		panic(err)
	}
	fmt.Println(tokenString)
}

// 用于签名的密钥
var secret = []byte("taoyuanming.com")

func GenRegisteredClaims() (string, error) {
	// 创建 Claims
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 过期时间
		Issuer:    "admin",
	}

	// 生产token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生产签名字符串
	return token.SignedString(secret)

}

// ParseRegisteredClaims 解析jwt
func ParseRegisteredClaims(tokenString string) bool {
	//  解析token
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return false
	}
	return token.Valid
}
