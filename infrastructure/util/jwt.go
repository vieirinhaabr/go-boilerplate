package util

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

var (
	SigningKey = "$SolidSigningKey$"
)

type jwtMapper struct {
	UserId   interface{}
	Username interface{}
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func ValidateToken(authRoles []string, token string) (*jwtMapper, string) {
	b := "Bearer: "
	if !strings.Contains(token, b) {
		return nil, "Your request is not authorized"
	}

	t := strings.Split(token, b)
	if len(t) < 2 {
		return nil, "An authorization token was not supplied"
	}

	valid, err := jwt.Parse(t[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(SigningKey), nil
	})
	if err != nil {
		return nil, "Invalid authorization token"
	}

	roles := valid.Claims.(jwt.MapClaims)["roles"].([]interface{})
	var tokenRoles []string
	for _, v := range roles {
		tokenRoles = append(tokenRoles, v.(string))
	}
	for _, v := range authRoles {
		hasRole := contains(tokenRoles, v)
		if !hasRole {
			return nil, "Not authorized to perform action"
		}
	}

	maps := valid.Claims.(jwt.MapClaims)
	return &jwtMapper{
		UserId:   maps["user_id"],
		Username: maps["username"],
	}, ""
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
