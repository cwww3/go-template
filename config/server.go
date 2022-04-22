package config

type ServerCfg struct {
	IP   string `json:"ip" yaml:"ip"`
	Port string `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
}
