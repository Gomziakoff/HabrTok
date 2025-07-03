package config

import "github.com/spf13/viper"

type ParserConfig struct {
	RssUrl   string `mapstructure:"rss_url"`
	Interval string `mapstructure:"interval"`
	DbDsn    string `mapstructure:"db_dsn"`
	EsHost   string `mapstructure:"es_host"`
	LogLevel string `mapstructure:"log_level"`
}

func LoadConfig() (*ParserConfig, error) {
	v := viper.New()
	v.SetConfigFile("config/config.yaml")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	var cfg ParserConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
