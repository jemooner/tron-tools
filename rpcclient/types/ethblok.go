package types

import "math/big"

// ============================发送交易返回结果========================
type TxResult struct {
	Result bool   `json:"result"`
	Txid   string `json:"txid"`
}

// ============================发送交易错误返回========================
type TxErrorResult struct {
	Code    string `json:"code"`
	Txid    string `json:"txid"`
	Message string `json:"message"`
}

// ============================tron通过区块号获取区块信息========================
type Block struct {
	BlockID      string         `json:"blockID"`
	BlockHeader  BlockHeader    `json:"block_header"`
	Transactions []Transactions `json:"transactions"`
}
type BlockHeader struct {
	RawData          BlockRawData `json:"raw_data"`
	WitnessSignature string       `json:"witness_signature"`
}
type BlockRawData struct {
	Number         int    `json:"number"`
	TxTrieRoot     string `json:"txTrieRoot"`
	WitnessAddress string `json:"witness_address"`
	ParentHash     string `json:"parentHash"`
	Version        int    `json:"version"`
	Timestamp      int64  `json:"timestamp"`
}

// ============================tron通过交易哈希获取交易信息（固始化）========================
type Transactions struct {
	Ret        []Ret    `json:"ret"`
	Signature  []string `json:"signature"`
	TxID       string   `json:"txID"`
	RawData    RawData  `json:"raw_data"`
	RawDataHex string   `json:"raw_data_hex"`
}
type Ret struct {
	ContractRet string `json:"contractRet"`
}
type Value struct {
	Amount          *big.Int `json:"amount"`
	AssetName       string   `json:"asset_name"`
	Data            string   `json:"data"`
	OwnerAddress    string   `json:"owner_address"`
	ContractAddress string   `json:"contract_address"`
	ToAddress       string   `json:"to_address"`
}
type Parameter struct {
	Value   Value  `json:"value"`
	TypeURL string `json:"type_url"`
}
type Contract struct {
	Parameter Parameter `json:"parameter"`
	Type      string    `json:"type"`
}
type RawData struct {
	Contract      []Contract `json:"contract"`
	RefBlockBytes string     `json:"ref_block_bytes"`
	RefBlockHash  string     `json:"ref_block_hash"`
	Epiration     int64      `json:"epiration"`
	FeeLimit      int        `json:"fee_limit"`
	Timestamp     int64      `json:"timestamp"`
}

// ============================tron通过交易哈希获取交易信息（合约交易）========================
type TransactionsInfo struct {
	ID                   string                 `json:"id"`
	Fee                  int                    `json:"fee"`
	BlockNumber          int                    `json:"blockNumber"`
	BlockTimeStamp       int64                  `json:"blockTimeStamp"`
	ContractResult       []string               `json:"contractResult"`
	ContractAddress      string                 `json:"contract_address"`
	Receipt              Receipt                `json:"receipt"`
	Log                  []Log                  `json:"log"`
	InternalTransactions []InternalTransactions `json:"internal_transactions"`
	PackingFee           int                    `json:"packingFee"`
}
type Receipt struct {
	EnergyUsage      int    `json:"energy_usage"`
	EnergyFee        int    `json:"energy_fee"`
	EnergyUsageTotal int    `json:"energy_usage_total"`
	NetFee           int    `json:"net_fee"`
	Result           string `json:"result"`
}
type Log struct {
	Address string   `json:"address"`
	Topics  []string `json:"topics"`
	Data    string   `json:"data"`
}
type CallValueInfo struct {
}
type InternalTransactions struct {
	Hash              string          `json:"hash"`
	CallerAddress     string          `json:"caller_address"`
	TransferToAddress string          `json:"transferTo_address"`
	CallValueInfo     []CallValueInfo `json:"callValueInfo"`
	Note              string          `json:"note"`
}

// ============================tron通过地址获取地址余额========================
type AccountBalance struct {
	Data    []Data `json:"data"`
	Success bool   `json:"success"`
	Meta    Meta   `json:"meta"`
}
type Keys struct {
	Address string `json:"address"`
	Weight  int    `json:"weight"`
}
type OwnerPermission struct {
	Keys           []Keys `json:"keys"`
	Threshold      int    `json:"threshold"`
	PermissionName string `json:"permission_name"`
}
type AccountResource struct {
	LatestConsumeTimeForEnergy int64 `json:"latest_consume_time_for_energy"`
}
type ActivePermission struct {
	Operations     string `json:"operations"`
	Keys           []Keys `json:"keys"`
	Threshold      int    `json:"threshold"`
	ID             int    `json:"id"`
	Type           string `json:"type"`
	PermissionName string `json:"permission_name"`
}

type Data struct {
	LatestOprationTime    uint64              `json:"latest_opration_time"`
	OwnerPermission       OwnerPermission     `json:"owner_permission"`
	FreeNetUsage          int                 `json:"free_net_usage"`
	AccountResource       AccountResource     `json:"account_resource"`
	ActivePermission      []ActivePermission  `json:"active_permission"`
	Address               string              `json:"address"`
	Balance               uint64              `json:"balance"`
	CreateTime            uint64              `json:"create_time"`
	Trc20                 []map[string]string `json:"trc20"`
	LatestConsumeFreeTime int64               `json:"latest_consume_free_time"`
}

type Meta struct {
	At       int64 `json:"at"`
	PageSize int   `json:"page_size"`
}

// ============================查询地址trx余额========================
type GetAccountInfo struct {
	Address          string                        `json:"address"`
	Balance          uint64                        `json:"balance"`
	CreateTime       int64                         `json:"create_time"`
	AccountResource  AccountInfoAccountResource    `json:"account_resource"`
	OwnerPermission  AccountInfoOwnerPermission    `json:"owner_permission"`
	ActivePermission []AccountInfoActivePermission `json:"active_permission"`
}
type AccountInfoAccountResource struct {
}
type AccountInfoKeys struct {
	Address string `json:"address"`
	Weight  int    `json:"weight"`
}
type AccountInfoOwnerPermission struct {
	PermissionName string            `json:"permission_name"`
	Threshold      int               `json:"threshold"`
	Keys           []AccountInfoKeys `json:"keys"`
}
type AccountInfoActivePermission struct {
	Type           string `json:"type"`
	ID             int    `json:"id"`
	PermissionName string `json:"permission_name"`
	Threshold      int    `json:"threshold"`
	Operations     string `json:"operations"`
	Keys           []Keys `json:"keys"`
}

// ============================查询地址trc20余额========================
type GetAddressTrcBalance struct {
	Result         TrcBalanceResult `json:"result"`
	EnergyUsed     int              `json:"energy_used"`
	ConstantResult []string         `json:"constant_result"`
	Transaction    Transaction      `json:"transaction"`
}
type TrcBalanceResult struct {
	Result bool `json:"result"`
}
type TrcBalanceRet struct {
}
type TrcBalanceValue struct {
	Data            string `json:"data"`
	OwnerAddress    string `json:"owner_address"`
	ContractAddress string `json:"contract_address"`
}
type TrcBalanceParameter struct {
	Value   TrcBalanceValue `json:"value"`
	TypeURL string          `json:"type_url"`
}
type TrcBalanceContract struct {
	Parameter Parameter `json:"parameter"`
	Type      string    `json:"type"`
}
type TrcBalanceRawData struct {
	Contract      []TrcBalanceContract `json:"contract"`
	RefBlockBytes string               `json:"ref_block_bytes"`
	RefBlockHash  string               `json:"ref_block_hash"`
	Expiration    int64                `json:"expiration"`
	Timestamp     int64                `json:"timestamp"`
}
type Transaction struct {
	Ret        []Ret             `json:"ret"`
	Visible    bool              `json:"visible"`
	TxID       string            `json:"txID"`
	RawData    TrcBalanceRawData `json:"raw_data"`
	RawDataHex string            `json:"raw_data_hex"`
}

//============================tron获取账户历史TRC20交易记录========================

type TrcTxInfo struct {
	Data    []TxData `json:"data"`
	Success bool     `json:"success"`
	Meta    Meta     `json:"meta"`
}
type TokenInfo struct {
	Symbol   string `json:"symbol"`
	Address  string `json:"address"`
	Decimals int    `json:"decimals"`
	Name     string `json:"name"`
}
type TxData struct {
	TransactionID  string    `json:"transaction_id"`
	TokenInfo      TokenInfo `json:"token_info"`
	BlockTimestamp int64     `json:"block_timestamp"`
	From           string    `json:"from"`
	To             string    `json:"to"`
	Type           string    `json:"type"`
	Value          string    `json:"value"`
}
