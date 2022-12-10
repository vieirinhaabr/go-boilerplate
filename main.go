package main

import (
	"github.com/VoidArtanis/go-rest-boilerplate/routes"
	"github.com/VoidArtanis/go-rest-boilerplate/shared"
	"github.com/gin-gonic/gin"
)

var DB = make(map[string]string)

func main() {

	//Db Connect and Close
	shared.Init()
	defer shared.CloseDb()

	//Init Gin
	r := gin.Default()
	routes.InitRouter(r)

	//Run Server
	r.Run(":8080")
}
