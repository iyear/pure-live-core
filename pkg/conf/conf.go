package conf

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
)

var C config

func Init(file string) {
	c := viper.New()
	c.SetConfigFile(file)
	// If a config file is found, read it in.
	if err := c.ReadInConfig(); err != nil {
		log.Fatalf("cannot read in config, file: %s", c.ConfigFileUsed())
		return
	}

	if err := c.Unmarshal(&C); err != nil {
		zap.S().Fatalw("cannot unmarshal config", "file", c.ConfigFileUsed(), "error", err)
		return
	}

	zap.S().Infof("read in config succ...")
}
