package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tron-tools/config"
	"tron-tools/pkg/go-logger"
	"tron-tools/router/apicommon"
	"tron-tools/rpcclient/jsonrpc"
	"tron-tools/rpcclient/taskcommon"
)

var rc *taskcommon.RequestChain

// 初始化节点连接
func CreateConnet() *taskcommon.RequestChain {
	clent := jsonrpc.NewETHTP(config.Conf.Tron.NodeUrl)
	rc = taskcommon.NewRequestChain(clent)
	return rc
}

// 获取最新区块号
func LastBlock(ctx *gin.Context) {
	//获取最新区块号
	num, err := rc.GetLastBlockNum()
	if err != nil {
		logger.Error("LastBlock", "block get failed", err.Error())
		apicommon.ReturnErrorResponse(ctx, 1, err.Error(), "")
		return
	}
	var BlockHeight BlockResponse
	BlockHeight.Block = num
	apicommon.ReturnSuccessResponse(ctx, 0, "success", BlockHeight)
}

// 根据区块高度获取区块信息
func BlockInfo(ctx *gin.Context) {
	params := new(GetBlockInfoBody)
	if err := ctx.ShouldBind(params); err != nil {
		logger.Error("BlockInfo", "Parameter is wrong", err.Error())
		apicommon.ReturnErrorResponse(ctx, 1, err.Error(), "")
		return
	}
	logger.Debug("GetBlockInfo", "paramInfo", params.GetBlockInfoParamsString())
	
	//将string到uint64
	numnerInt, err := strconv.ParseUint(params.Number, 10, 64)
	//根据区块号获取区块信息
	res, err := rc.GetBlockByNumber(numnerInt)
	if err != nil {
		logger.Error("GetBlockByNumber", "Get BlockInfo failed", err.Error())
		apicommon.ReturnErrorResponse(ctx, 1, err.Error(), "")
		return
	}
	apicommon.ReturnSuccessResponse(ctx, 0, "success", res)
}
