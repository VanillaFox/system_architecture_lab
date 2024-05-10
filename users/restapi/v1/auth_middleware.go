package v1

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const tokenType = "Bearer"

var ErrAuthFail = errors.New("authorization fail")

var jwtSecretKey []byte

func InitJwtSecretKey(val []byte) {
	jwtSecretKey = val
}

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	splitedHeader := strings.Split(authHeader, " ")

	if len(splitedHeader) != 2 || splitedHeader[0] != tokenType {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": ErrAuthFail.Error()})

		return
	}

	token, err := jwt.Parse(splitedHeader[1], func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": ErrAuthFail.Error()})

		return
	}

	_, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": ErrAuthFail.Error()})

		return
	}

	c.Next()
}
