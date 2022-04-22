package config

import "fmt"

type MysqlCfg struct {
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	IP       string `json:"ip" yaml:"ip"`
	Port     string `json:"port" yaml:"port"`
	DB       string `json:"db" yaml:"db"`
}

func (m *MysqlCfg) GetDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", m.User, m.Password, m.IP, m.Port, m.DB)
}
