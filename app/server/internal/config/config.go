package config

import (
	"github.com/spf13/viper"
)

var Server server

func InitServer(path string) error {
	c := viper.New()
	c.SetConfigFile(path)
	if err := c.ReadInConfig(); err != nil {
		return err
	}
	if err := c.Unmarshal(&Server); err != nil {
		return err
	}

	return nil
}

type server struct {
	Port   int    `mapstructure:"port"`
	Debug  bool   `mapstructure:"debug"`
	Path   string `mapstructure:"path"`
	Socks5 socks5 `mapstructure:"socks5"`
}

type socks5 struct {
	Enable   bool   `mapstructure:"enable"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}
