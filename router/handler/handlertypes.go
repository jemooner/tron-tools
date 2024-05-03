package handler

// 获取最新区块高度返回结构体
type BlockResponse struct {
	Block interface{} `json:"blockHeight"`
}

// GetBlockInfoBody 获取区块信息接收struct
type GetBlockInfoBody struct {
	Number string `form:"number" json:"number" binding:"required"`
}

func (c *GetBlockInfoBody) GetBlockInfoParamsString() string {
	temp := ""
	temp += "number=" + c.Number
	return temp
}
