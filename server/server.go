package server

import (
	"fmt"
	"github.com/iyear/pure-live/conf"
	"github.com/iyear/pure-live/db"
	"github.com/iyear/pure-live/logger"
	"github.com/iyear/pure-live/router"
	"github.com/iyear/pure-live/util"
	"github.com/q191201771/naza/pkg/nazalog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Run(cfgFile string) {
	// lal包中的nazalog
	_ = nazalog.Init(func(option *nazalog.Option) {
		option.IsToStdout = conf.C.Server.Debug
	})

	conf.Init(cfgFile)
	logger.Init(util.IF(conf.C.Server.Debug, zapcore.DebugLevel, zapcore.InfoLevel).(zapcore.LevelEnabler))

	zap.S().Infof("init server...")
	if err := os.MkdirAll(conf.C.Server.Path, 0774); err != nil {
		zap.S().Fatalw("failed to mkdir", "error", err)
	}

	if err := db.Init(); err != nil {
		zap.S().Fatalw("failed to init database", "error", err)
	}
	zap.S().Infof("init database succ...")

	zap.S().Infof("server runs on :%d,debug: %v", conf.C.Server.Port, conf.C.Server.Debug)
	engine := router.Init()
	err := engine.Run(fmt.Sprintf(":%d", conf.C.Server.Port))
	if err != nil {
		zap.S().Fatalw("failed to run gin engine", "error", err, "port", conf.C.Server.Port)
		return
	}
}
