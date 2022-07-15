package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// custom claims
type Claims struct {
	Account string      `json:"account"`
	Role string         `json:"role"`
	jwt.StandardClaims
}

// jwt secret key
var jwtSecret = []byte("secret")

func Login(c *gin.Context) {
	// validate request body
	var body struct{
		Account string
		Password string
	}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// check account and password is correct
	if body.Account == "Kenny" && body.Password == "123456" {
		now := time.Now()
		jwtId := body.Account + strconv.FormatInt(now.Unix(), 10)
		role := "Member"

		// set claims and sign
		claims := Claims{
			Account:        body.Account,
			Role:           role,
			StandardClaims: jwt.StandardClaims{
				Audience:  body.Account,
				ExpiresAt: now.Add(20 * time.Second).Unix(),
				Id:        jwtId,
				IssuedAt:  now.Unix(),
				Issuer:    "ginJWT",
				NotBefore: now.Add(10 * time.Second).Unix(),
				Subject:   body.Account,
			},
		}
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err := tokenClaims.SignedString(jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		return
	}

	// incorrect account or password
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "Unauthorized",
	})
}