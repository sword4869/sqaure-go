package controller

import (
	"test/config"

	"github.com/gin-gonic/gin"
)

var g *gin.Engine

func init() {
	config.InitConfig("../config")
	g = Router()
}
