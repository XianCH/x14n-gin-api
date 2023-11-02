package model

import "github.com/golang-jwt/jwt"

type UserCliams struct {
	*jwt.StandardClaims
	Uid uint
}
