package config

// viper 会根据标签 value 值把 config.yaml 的数据赋予给结构体
type App struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	Port    int    `mapstructure:"port" json:"port" yaml:"port"`
	AppName string `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
	AppUrl  string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`
	DbType  string `mapstructure:"db_type" json:"db_type" yaml:"db_type"`
}
