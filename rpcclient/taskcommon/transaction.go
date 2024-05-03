package taskcommon


type TrcTransaction struct {
	Result Result `json:"result"`
	Transaction Transaction `json:"transaction"`
}
type Result struct {
	Result bool `json:"result"`
}
type Value struct {
	Data string `json:"data"`
	OwnerAddress string `json:"owner_address"`
	ContractAddress string `json:"contract_address"`
}
type Parameter struct {
	Value Value `json:"value"`
	TypeURL string `json:"type_url"`
}
type Contract struct {
	Parameter Parameter `json:"parameter"`
	Type string `json:"type"`
}
type RawData struct {
	Contract []*Contract `json:"contract"`
	RefBlockBytes string `json:"ref_block_bytes"`
	RefBlockHash string `json:"ref_block_hash"`
	Expiration int64 `json:"expiration"`
	FeeLimit int `json:"fee_limit"`
	Timestamp int64 `json:"timestamp"`
}
type Transaction struct {
	RawData RawData `json:"raw_data"`
	TxID string `json:"txID"`
	RawDataHex string `json:"raw_data_hex"`
	Visible bool `json:"visible"`
	Signature []string `json:"signature"`
}


type TrxTransaction struct {
	Visible bool `json:"visible"`
	TxID string `json:"txID"`
	RawData TrxRawData `json:"raw_data"`
	RawDataHex string `json:"raw_data_hex"`
	Signature  []string `json:"signature"`
}

type TrxRawData struct {
	Contract      []*TrxContract `json:"contract"`
	RefBlockBytes string          `json:"ref_block_bytes"`
	RefBlockHash  string          `json:"ref_block_hash"`
	Expiration    int64           `json:"expiration"`
	Timestamp     int64           `json:"timestamp"`
}

type TrxContract struct {
	Parameter TrxParameter `json:"parameter"`
	Type      string    `json:"type"`
}

type TrxParameter struct {
	Value   TrxValue `json:"value"`
	TypeUrl string         `json:"type_url"`
}

type TrxValue struct {
	Amount       int64  `json:"amount"`
	OwnerAddress string `json:"owner_address"`
	ToAddress    string `json:"to_address"`
}