package utils

import (
	"TinyTikTok/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	models.CommonEntity
	Name string `json:"name"`
	jwt.StandardClaims
}

// 密匙
var jwtSigningKey = []byte("tinytiktok")

// GenerateToken 生成 token
func GenerateToken(name string, commonEntity models.CommonEntity) (string, error) {
	userClaims := &UserClaims{
		CommonEntity:   commonEntity,
		Name:           name,
		StandardClaims: jwt.StandardClaims{},
	}
	// 使用 HS256 签名方法创建一个 JWT，并使用 jwtSigningKey 进行签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := token.SignedString(jwtSigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

// ParseToken 解析 token
func ParseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)

	// 解析JWT，验证签名
	claim, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) { // 回调函数，用于提供用于验证 JWT 签名的密钥。这个函数会在 JWT 解析的过程中被调用
		return jwtSigningKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}
	if !claim.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return userClaim, nil
}
