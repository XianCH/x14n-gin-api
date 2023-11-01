package middle

import "github.com/golang-jwt/jwt"

func CreateToken(uid uint) (string, error) {
	newWithClaim := jwt.NewWithClaims()
}
