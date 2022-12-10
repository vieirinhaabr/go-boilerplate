package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var (
	SigningKey = "$SolidSigningKey$"
)

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func validateToken(tokenString string, key string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	return token, err
}

func AuthHandler(authRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		b := "Bearer: "
		if !strings.Contains(token, b) {
			c.JSON(403, gin.H{"message": "Your request is not authorized"})
			c.Abort()
			return
		}

		t := strings.Split(token, b)
		if len(t) < 2 {
			c.JSON(403, gin.H{"message": "An authorization token was not supplied"})
			c.Abort()
			return
		}

		valid, err := validateToken(t[1], SigningKey)
		if err != nil {
			c.JSON(403, gin.H{"message": "Invalid authorization token"})
			c.Abort()
			return
		}

		roles := valid.Claims.(jwt.MapClaims)["roles"].([]interface{})
		var tokenRoles []string
		for _, v := range roles {
			tokenRoles = append(tokenRoles, v.(string))
		}
		for _, v := range authRoles {
			hasRole := contains(tokenRoles, v)
			if !hasRole {
				c.JSON(403, gin.H{"message": "Not authorized to perform action"})
				c.Abort()
			}
		}

		c.Set("userId", valid.Claims.(jwt.MapClaims)["user_id"])
		c.Set("username", valid.Claims.(jwt.MapClaims)["username"])
		c.Next()
	}
}

func GenerateToken(key []byte, userId string, username string, roles []string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := make(jwt.MapClaims)
	claims["user_id"] = userId
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour*72).UnixNano() / int64(time.Millisecond)

	claims["roles"] = roles

	token.Claims = claims

	tokenString, err := token.SignedString(key)
	return tokenString, err
}
