package util

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}
