package main

import (
	"context"
	"fmt"
	"go_cli/demo/database/postgres"
	"go_cli/demo/database/redis"
	"go_cli/demo/logger"
	"go_cli/demo/router"
	"go_cli/demo/setting"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"

	"go.uber.org/zap"
)

func main() {
	// settings
	if err := setting.Init(); err != nil {
		fmt.Printf("Config file initialization failed,err:%v\n\n", err)
		return
	}
	// logger
	if err := logger.Init(); err != nil {
		fmt.Printf("Logger initialization failed,err:%v\n\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("Logger initialization succeed ^_")
	// database(postgres)
	if err := postgres.Init(); err != nil {
		fmt.Printf("Database initialization failed,err:%v\n\n", err)
		return
	}
	defer postgres.Close()
	// redis
	if err := redis.Init(); err != nil {
		fmt.Printf("Redis initialization failed,err:%v\n\n", err)
		return
	}
	defer redis.Close()
	// router
	r := router.Init()
	serve := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.Get("app.port")),
		Handler: r,
	}

	// start a go runtime to start server
	go func() {
		if err := serve.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Info(fmt.Sprintf("listen:%s\n", err))
		}
	}()

	// close server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Fatal("Shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := serve.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server shutdown", zap.Error(err))
	}
	log.Println("Server existing")
}
