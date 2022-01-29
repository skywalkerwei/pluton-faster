package config

// RedisConf redis配置
type RedisConf struct {
	Host string `json:"Host"`
	Pass string `json:"Pass"`
	Type string `json:"Type"`
}
