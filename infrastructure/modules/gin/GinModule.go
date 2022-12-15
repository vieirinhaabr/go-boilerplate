package ginModule

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/config/environment"
	"go-boilerplate/infrastructure/modules/gin/routes"
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

func (handler *Module) Start() {
	err := handler.engine.Run(":" + environment.Vars.Gin.Port)
	if err != nil {
		panic(err)
	}
}

func (handler *Module) Finish() {}
