package config

// FIXME zap通过mapstructure解析yml配置文件
// https://github.com/mitchellh/mapstructure

type ZapCfg struct {
	FileStdout    bool   `mapstructure:"fileStdout"`
	ConsoleStdout bool   `mapstructure:"consoleStdout"`
	Level         string `mapstructure:"level"`
	Lumber        `mapstructure:",squash"`
}

type Lumber struct {
	Path       string `mapstructure:"path"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxAge     int    `mapstructure:"maxAge"`
	MaxBackups int    `mapstructure:"maxBackups"`
	LocalTime  bool   `mapstructure:"localTime"`
	Compress   bool   `mapstructure:"compress"`
}
