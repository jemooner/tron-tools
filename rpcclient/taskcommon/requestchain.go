package taskcommon

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"strings"
	"tron-tools/config"
	"tron-tools/pkg/common"
	"tron-tools/pkg/crypto/base58"
	"tron-tools/pkg/go-logger"
	"tron-tools/rpcclient/jsonrpc"
	"tron-tools/rpcclient/types"
)

type RequestChain struct {
	Client *jsonrpc.Http
}

func NewRequestChain(c *jsonrpc.Http) *RequestChain {
	rc := new(RequestChain)
	rc.Client = c
	return rc
}

// GetLastBlockNum 获取最新区块
func (rc RequestChain) GetLastBlockNum() (interface{}, error) {
	res, err := rc.Client.GetNowBlock()
	if err != nil {
		logger.Error("GetBlockByNumber", "step", "GetBlockByNumber", "err", err.Error())
		return "", err
	}
	BlockInfo := new(types.Block)
	if err := json.Unmarshal(res, &BlockInfo); err != nil {
		logger.Error("GetBlockByNumber", "step", "Unmarshal transaction", "err", err.Error())
		return "", err
	}
	return BlockInfo.BlockHeader.RawData.Number, nil
}

// GetBlockByNumber 根据块号获取区块信息
func (rc RequestChain) GetBlockByNumber(number uint64) (interface{}, error) {
	res, err := rc.Client.GetBlockByNumber(number)
	if err != nil {
		logger.Error("GetBlockByNumber", "step", "GetBlockByNumber", "err", err.Error())
		return types.Block{}, err
	}
	BlockInfo := new(types.Block)
	if err := json.Unmarshal(res, &BlockInfo); err != nil {
		logger.Error("GetBlockByNumber", "step", "Unmarshal transaction", "err", err.Error())
		return types.Block{}, err
	}
	//循环修改Hex格式地址为Base58格式
	ts := BlockInfo.Transactions
	for k := 0; k < len(ts); k++ {
		// 判断是否为TRX交易，如果是，则将HEX格式的form及to地址修改为base58格式
		if ts[k].RawData.Contract[0].Type == "TransferContract" {
			ts[k].RawData.Contract[0].Parameter.Value.OwnerAddress = HexToAddress(ts[k].RawData.Contract[0].Parameter.Value.OwnerAddress)
			ts[k].RawData.Contract[0].Parameter.Value.ToAddress = HexToAddress(ts[k].RawData.Contract[0].Parameter.Value.ToAddress)
			//判断合约地址是否为TRC20-USDT，如果是则进行以下流程
		} else if ts[k].RawData.Contract[0].Type == "TriggerSmartContract" {
			ts[k].RawData.Contract[0].Parameter.Value.OwnerAddress = HexToAddress(ts[k].RawData.Contract[0].Parameter.Value.OwnerAddress)
			ts[k].RawData.Contract[0].Parameter.Value.ContractAddress = HexToAddress(ts[k].RawData.Contract[0].Parameter.Value.ContractAddress)
			data := ts[k].RawData.Contract[0].Parameter.Value.Data
			//判断合约地址是否为TRC20-USDT，如果是则进行以下流程
			if ts[k].RawData.Contract[0].Parameter.Value.ContractAddress == config.Conf.Tron.ContractAddr {
				//判断data方法是否为a9059cbb（transfer）方法，如果是则进行以下流程
				if data[:8] == "a9059cbb" {
					//截取HEX格式的to地址并转为base58后，放到该交易体中的to_address中
					to := "41" + data[32:72]
					ts[k].RawData.Contract[0].Parameter.Value.ToAddress = HexToAddress(to)
					//截取16进制的交易金额后转为10进制的bigint整数，放到该交易体中的amount中
					amount := "0x" + data[73:136]
					ts[k].RawData.Contract[0].Parameter.Value.Amount = common.HexToBigInt(amount)
					//判断data方法是否为23b872dd（transferFrom）方法，如果是则进行以下流程
				} else if data[:8] == "23b872dd" {
					//截取HEX格式的from地址并转为base58后，放到该交易体中的owner_address中
					from := "41" + data[32:72]
					ts[k].RawData.Contract[0].Parameter.Value.OwnerAddress = HexToAddress(from)
					//截取HEX格式的to地址并转为base58后，放到该交易体中的to_address中
					to := "41" + data[96:136]
					ts[k].RawData.Contract[0].Parameter.Value.ToAddress = HexToAddress(to)
					//截取16进制的交易金额后转为10进制的bigint整数，放到该交易体中的amount中
					amount := "0x" + data[136:200]
					ts[k].RawData.Contract[0].Parameter.Value.Amount = common.HexToBigInt(amount)
				}
			}
		}
	}
	return BlockInfo, nil
}

// GetBlockByNumber 根据块号获取区块信息
func (rc RequestChain) GetBlockByNumberInfo(number uint64) (interface{}, error) {
	//=========================获取指定区块信息================================================================================
	res, err := rc.Client.GetBlockByNumber(number)
	if err != nil {
		logger.Error("GetBlockByNumber", "step", "GetBlockByNumber", "err", err.Error())
		return "", err
	}
	BlockInfo := new(types.Block)
	if err := json.Unmarshal(res, &BlockInfo); err != nil {
		logger.Error("GetBlockByNumber", "step", "Unmarshal transaction", "err", err.Error())
		return "", err
	}
	//=========================for循环遍历区块信息中的Transactions==============================================================
	//for循环遍历区块信息中的Transactions
	ts := BlockInfo.Transactions
	for k := 0; k < len(ts); k++ {
		// 判断是否为TRX交易，如果是，则将HEX格式的form及to地址修改为base58格式
		if ts[k].RawData.Contract[0].Type == "TransferContract" {
			ts[k].RawData.Contract[0].Parameter.Value.OwnerAddress = HexToAddress(ts[k].RawData.Contract[0].Parameter.Value.OwnerAddress)
			ts[k].RawData.Contract[0].Parameter.Value.ToAddress = HexToAddress(ts[k].RawData.Contract[0].Parameter.Value.ToAddress)
		}
		// 判断是否为TRC20交易，如果是，则将HEX格式的form及合约地址修改为base58格式
		if ts[k].RawData.Contract[0].Type == "TriggerSmartContract" {
			ts[k].RawData.Contract[0].Parameter.Value.OwnerAddress = HexToAddress(ts[k].RawData.Contract[0].Parameter.Value.OwnerAddress)
			ts[k].RawData.Contract[0].Parameter.Value.ContractAddress = HexToAddress(ts[k].RawData.Contract[0].Parameter.Value.ContractAddress)
			data := ts[k].RawData.Contract[0].Parameter.Value.Data
			//判断合约地址是否为TRC20-USDT，如果是则进行以下流程
			if ts[k].RawData.Contract[0].Parameter.Value.ContractAddress == config.Conf.Tron.ContractAddr {
				//判断data方法是否为a9059cbb（transfer）方法，如果是则进行以下流程
				if data[:8] == "a9059cbb" {
					//截取HEX格式的to地址并转为base58后，放到该交易体中的to_address中
					to := "41" + data[32:72]
					ts[k].RawData.Contract[0].Parameter.Value.ToAddress = HexToAddress(to)
					//截取16进制的交易金额后转为10进制的bigint整数，放到该交易体中的amount中
					amount := "0x" + data[73:136]
					ts[k].RawData.Contract[0].Parameter.Value.Amount = common.HexToBigInt(amount)
					//判断data方法是否为23b872dd（transferFrom）方法，如果是则进行以下流程
				} else if data[:8] == "23b872dd" {
					//截取HEX格式的from地址并转为base58后，放到该交易体中的owner_address中
					from := "41" + data[32:72]
					ts[k].RawData.Contract[0].Parameter.Value.OwnerAddress = HexToAddress(from)
					//截取HEX格式的to地址并转为base58后，放到该交易体中的to_address中
					to := "41" + data[96:136]
					ts[k].RawData.Contract[0].Parameter.Value.ToAddress = HexToAddress(to)
					//截取16进制的交易金额后转为10进制的bigint整数，放到该交易体中的amount中
					amount := "0x" + data[136:200]
					ts[k].RawData.Contract[0].Parameter.Value.Amount = common.HexToBigInt(amount)
				}
			}
		}
	}
	return BlockInfo, nil
}

// 将波场Hex格式地址转为Base58格式
func HexToAddress(s string) string {
	if HasHexPrefix(s) {
		s = s[2:]
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	addrw, err := Hex2Bytes(s)
	if err != nil {
		return "nil"
	}
	return EncodeCheck(addrw)
}

// hasHexPrefix 验证str以“0x”或“0X”开头。
func HasHexPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

// Hex2Bytes 返回十六进制字符串str所代表的字节。
func Hex2Bytes(str string) ([]byte, error) {
	h, err := hex.DecodeString(str)
	return h, err
}

// HexToBigInt
func HexToBigInt(hex string) *big.Int {
	if strings.HasPrefix(hex, "0x") {
		hex = hex[2:]
	}
	n := new(big.Int)
	n, _ = n.SetString(hex[:], 16)
	return n
}

func EncodeCheck(input []byte) string {
	h256h0 := sha256.New()
	h256h0.Write(input)
	h0 := h256h0.Sum(nil)

	h256h1 := sha256.New()
	h256h1.Write(h0)
	h1 := h256h1.Sum(nil)

	inputCheck := input
	inputCheck = append(inputCheck, h1[:4]...)
	return base58.Encode(inputCheck, base58.BitcoinAlphabet)
}
