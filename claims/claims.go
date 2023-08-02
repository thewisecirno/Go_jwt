package claims

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	ID       int64
	Username string
	jwt.StandardClaims
}
