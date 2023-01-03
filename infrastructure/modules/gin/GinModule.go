package ginModule

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/config/environment"
	"go-boilerplate/infrastructure/modules/gin/routes"
	"go-boilerplate/utils"
)

type Module struct {
	engine *gin.Engine
}

func NewGinModule() *Module {
	return &Module{}
}

func (handler *Module) Configure() {
	utils.Log.Warn("[GIN] ᛃ Configuring\n")

	engine := gin.Default()
	router.SetRoutes(engine)
	*handler = Module{engine}
}

func (handler *Module) Start() error {
	utils.Log.Warn("[GIN] ▶ Starting\n")

	var config = environment.Vars.Gin

	utils.Log.Warn("[GIN] Listening on %s\n", config.Port)
	return handler.engine.Run(":" + config.Port)
}

func (handler *Module) Finish() {
	utils.Log.Warn("[GIN] ■ Finishing\n")
}
