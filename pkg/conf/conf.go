package conf

import (
	"github.com/spf13/viper"
)

var (
	Server  server
	Account account
)

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

func InitAccount(path string) error {
	c := viper.New()
	c.SetConfigFile(path)
	if err := c.ReadInConfig(); err != nil {
		return err
	}
	if err := c.Unmarshal(&Account); err != nil {
		return err
	}

	return nil
}
