package config

// yaml配置结构体
type Config struct {
	LogCfg    string     `json:"log"`
	Console   *Console   `json:"console"`
	Tron      *Tron      `json:"tron"`
	Redis     *Redis     `json:"redis"`
	SecretKey *SecretKey `json:"secretKey"`
}

type Console struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Port    string `json:"port"`
}

type Tron struct {
	NodeUrl      string `json:"nodeurl"`
	ContractAddr string `json:"contractAddr"`
}
type Redis struct {
	Password string `json:"password"`
	Host     string `json:"host"`
	Database int    `json:"database"`
}
type SecretKey struct {
	SingKey      string `json:"singKey"`
	SingSWitcher bool   `json:"singSWitcher"`
}
