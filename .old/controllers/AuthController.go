/**
 * Created by VoidArtanis on 11/2/2017
 */

package controllers

import (
	"github.com/VoidArtanis/go-rest-boilerplate/middlewares"
	"github.com/VoidArtanis/go-rest-boilerplate/shared"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (this AuthController) HandleLogin(c *gin.Context) {
	userId := "123"
	username := "Beast"
	roles := []string{shared.RoleAdmin, shared.RoleProUser}

	// do user auth here

	//issue token
	token, err := middlewares.GenerateToken([]byte(middlewares.SigningKey), userId, username, roles)
	if err != nil {

	}
	c.JSON(200, token)
}
