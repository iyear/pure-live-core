package conf

import (
	"github.com/spf13/viper"
)

var (
	Account account
)

// InitAccount init account config
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
