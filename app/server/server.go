package server

import (
	"context"
	"fmt"
	"github.com/iyear/pure-live-core/app/server/internal/config"
	"github.com/iyear/pure-live-core/app/server/internal/logger"
	"github.com/iyear/pure-live-core/app/server/internal/router"
	"github.com/iyear/pure-live-core/global"
	"github.com/iyear/pure-live-core/pkg/conf"
	"github.com/iyear/pure-live-core/pkg/db"
	"github.com/iyear/pure-live-core/pkg/request"
	"github.com/iyear/pure-live-core/pkg/util"
	"github.com/q191201771/naza/pkg/nazalog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func Run(serverConf string, accountConf string) {
	// lal包中的nazalog
	_ = nazalog.Init(func(option *nazalog.Option) {
		option.IsToStdout = config.Server.Debug
	})

	if err := config.InitServer(serverConf); err != nil {
		log.Fatalf("failed to read server config: %s", err)
	}
	if err := conf.InitAccount(accountConf); err != nil {
		log.Fatalf("failed to read account config: %s", err)
	}

	logger.Init(util.IF(config.Server.Debug, zapcore.DebugLevel, zapcore.InfoLevel).(zapcore.LevelEnabler))

	zap.S().Infof("read config succ...")

	if err := os.MkdirAll(config.Server.Path, 0774); err != nil {
		zap.S().Fatalw("failed to mkdir", "error", err)
	}

	sqlite, err := db.Init(filepath.Join(config.Server.Path, "data.db"))
	if err != nil {
		zap.S().Fatalw("failed to init database", "error", err)
	}
	global.DB = sqlite
	zap.S().Infof("init database succ...")

	if config.Server.Socks5.Enable {
		request.SetSocks5(config.Server.Socks5.Host, config.Server.Socks5.Port, config.Server.Socks5.User, config.Server.Socks5.Password)
	}

	zap.S().Infof("server runs on :%d,debug: %v", config.Server.Port, config.Server.Debug)

	handler := router.Init()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Server.Port),
		Handler: handler,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Infow("failed to start to listen and serve", "error", err, "port", config.Server.Port)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.S().Info("shutdown server...")

	ctx, stop := context.WithTimeout(context.Background(), 5*time.Second)
	defer stop()

	if err = s.Shutdown(ctx); err != nil {
		zap.S().Fatalw("server forced to shutdown", "error", err)
	}

	zap.S().Infow("server exited...")
}
