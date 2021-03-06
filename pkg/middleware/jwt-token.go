package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// MyClaims Customer jwt.StandardClaims
type MyClaims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}

// GenToken Create a new token
func GenToken(account string) (string, error) {
	c := MyClaims{
		account,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
			Issuer: "Flynn",
		},
	}
	// Choose specific algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// Choose specific Signature
	return token.SignedString("SecretKey")
}