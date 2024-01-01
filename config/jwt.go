package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("hbqwjbxu7228388edhjxjsnax01")

type JWTClaim struct{
	Username string
	jwt.RegisteredClaims
}