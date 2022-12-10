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

func (handler *ginModule) configure(res chan any) {
	engine := gin.Default()
	services.InitRouter(engine)
	*handler = ginModule{engine}
	res <- handler
}

func (handler *ginModule) start(res chan any) {
	handler.engine.Run(":8080")
	res <- handler
}
