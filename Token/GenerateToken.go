package Token

import (
	"github.com/dgrijalva/jwt-go"
)

// GenerateTokenString
// 获取CLAIMS结构体对应的token值
func GenerateTokenString(claims jwt.Claims) (string, error) {
	//func NewWithClaims(method SigningMethod, claims Claims) *Token
	//参数解析：method表示加密方法，claims表示对应的claims结构体，后续跟着.SignedString([]byte("golang"))，其中的参数其实是代表私钥
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("golang"))
	return token, err
}
