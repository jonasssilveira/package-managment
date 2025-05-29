package infra

import (
	"github.com/gin-gonic/gin"
)

type Rest interface {
	Get(string, func(ctx *gin.Context))
	Post(string, func(ctx *gin.Context))
	Run(string)
}
