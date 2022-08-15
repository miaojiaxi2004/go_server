package config

type Server struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
