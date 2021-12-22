package server

import (
	"fmt"
	"github.com/iyear/pure-live/app/server/internal/logger"
	"github.com/iyear/pure-live/app/server/internal/router"
	"github.com/iyear/pure-live/global"
	"github.com/iyear/pure-live/pkg/conf"
	"github.com/iyear/pure-live/pkg/db"
	"github.com/iyear/pure-live/pkg/request"
	"github.com/iyear/pure-live/pkg/util"
	"github.com/q191201771/naza/pkg/nazalog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
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

	sqlite, err := db.Init(path.Join(conf.C.Server.Path, "data.db"))
	if err != nil {
		zap.S().Fatalw("failed to init database", "error", err)
	}
	global.DB = sqlite
	zap.S().Infof("init database succ...")

	if conf.C.Socks5.Enable {
		request.SetSocks5(conf.C.Socks5.Host, conf.C.Socks5.Port, conf.C.Socks5.User, conf.C.Socks5.Password)
	}

	zap.S().Infof("server runs on :%d,debug: %v", conf.C.Server.Port, conf.C.Server.Debug)
	engine := router.Init()

	if err = engine.Run(fmt.Sprintf(":%d", conf.C.Server.Port)); err != nil {
		zap.S().Fatalw("failed to run gin engine", "error", err, "port", conf.C.Server.Port)
		return
	}
}
