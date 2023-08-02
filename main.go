package main

import (
	"GO_JWT/Token"
	"GO_JWT/claims"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	TokenString, err := Token.GenerateTokenString(claims.Claims{
		ID:             1,
		Username:       "Tom",
		StandardClaims: jwt.StandardClaims{},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", TokenString)

	value, err := Token.ParseTokenString(TokenString)
	fmt.Printf("%#v", value)
}
