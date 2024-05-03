package router

import (
	"github.com/gin-gonic/gin"
	"tron-tools/config"
	"tron-tools/pkg/go-logger"
	"tron-tools/router/handler"
)

// RouterStart
func RouterStart() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	search := router.Group("/api")
	{
		// 获取最新区块高度
		search.POST("/block/lastBlock", handler.LastBlock)
		// 根据区块号获取区块信息
		search.POST("/block/blockInfo", handler.BlockInfo)
	}

	logger.Info("API server start", "listen", config.Conf.Console.Port)
	router.Run(config.Conf.Console.Port)
}
