package Token

import (
	"GO_JWT/claims"
	"github.com/dgrijalva/jwt-go"
)

// ParseTokenString
// 解析Token字符串，将其转换为claims结构体对象
func ParseTokenString(tokenString string) (interface{}, error) {
	//func jwt.ParseWithClaims(tokenString string, claims jwt.Claims, keyFunc jwt.Keyfunc) (*jwt.Token, error)
	//参数解析：tokenString token字符串，claims 目标结构体对象，keyFunc表示通过这个token返回私钥
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &claims.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})

	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if value, ok := tokenClaims.Claims.(*claims.Claims); ok && tokenClaims.Valid {
			return value, nil
		}
	}

	return nil, err
}
