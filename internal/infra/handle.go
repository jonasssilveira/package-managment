package infra

import (
	"github.com/gin-gonic/gin"
)

type Handle struct {
	Engine *gin.Engine
	server Server
}

func NewHandle(server Server) Handle {
	return Handle{Engine: gin.New(), server: server}
}

func (h Handle) Get(path string, webHandle func(ctx *gin.Context)) {
	h.Engine.GET(path, webHandle)
}

func (h Handle) Post(path string, webHandle func(ctx *gin.Context)) {
	h.Engine.POST(path, webHandle)
}

func (h Handle) Delete(path string, webHandle func(ctx *gin.Context)) {
	h.Engine.DELETE(path, webHandle)
}


func (h Handle) Run() {
	err := h.Engine.Run(":" + h.server.port)
	if err != nil {
		return
	}
}
