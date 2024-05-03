package main

import (
	"tron-tools/config"
	"tron-tools/router"
	"tron-tools/router/handler"
)

func main() {
	// 初始化配置信息
	config.Initialize()
	// 初始化节点长链接
	handler.CreateConnet()
	// api服务启动
	router.RouterStart()
}
