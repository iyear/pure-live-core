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
	"log"
	"os"
	"path"
)

func Run(serverConf string, accountConf string) {
	// lal包中的nazalog
	_ = nazalog.Init(func(option *nazalog.Option) {
		option.IsToStdout = conf.Server.Debug
	})

	if err := conf.InitServer(serverConf); err != nil {
		log.Fatalf("failed to read server config: %s", err)
	}
	if err := conf.InitAccount(accountConf); err != nil {
		log.Fatalf("failed to read account config: %s", err)
	}

	logger.Init(util.IF(conf.Server.Debug, zapcore.DebugLevel, zapcore.InfoLevel).(zapcore.LevelEnabler))

	zap.S().Infof("read config succ...")

	if err := os.MkdirAll(conf.Server.Path, 0774); err != nil {
		zap.S().Fatalw("failed to mkdir", "error", err)
	}

	sqlite, err := db.Init(path.Join(conf.Server.Path, "data.db"))
	if err != nil {
		zap.S().Fatalw("failed to init database", "error", err)
	}
	global.DB = sqlite
	zap.S().Infof("init database succ...")

	if conf.Server.Socks5.Enable {
		request.SetSocks5(conf.Server.Socks5.Host, conf.Server.Socks5.Port, conf.Server.Socks5.User, conf.Server.Socks5.Password)
	}

	zap.S().Infof("server runs on :%d,debug: %v", conf.Server.Port, conf.Server.Debug)
	engine := router.Init()

	if err = engine.Run(fmt.Sprintf(":%d", conf.Server.Port)); err != nil {
		zap.S().Fatalw("failed to run gin engine", "error", err, "port", conf.Server.Port)
		return
	}
}
