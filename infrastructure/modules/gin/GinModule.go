package ginModule

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/config/environment"
	"go-boilerplate/infrastructure/modules/gin/routes"
	"log"
)

type Module struct {
	engine *gin.Engine
}

func NewGinModule() *Module {
	return &Module{}
}

func (handler *Module) Configure() {
	engine := gin.Default()
	router.SetRoutes(engine)
	*handler = Module{engine}
}

func (handler *Module) Start() error {
	return handler.engine.Run(":" + environment.Vars.Gin.Port)
}

func (handler *Module) Finish() {
	log.Printf("[GIN] Finished\n")
}
