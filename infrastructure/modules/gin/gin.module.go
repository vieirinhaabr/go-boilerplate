package gin

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/infrastructure/modules/gin/routes"
)

type ginModule struct {
	engine *gin.Engine
}

func NewGinModule() *ginModule {
	return &ginModule{}
}

func (handler *ginModule) Configure() {
	engine := gin.Default()
	services.InitRouter(engine)
	*handler = ginModule{engine}
}

func (handler *ginModule) Start() {
	handler.engine.Run(":8080")
}
