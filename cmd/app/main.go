package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/tiagoangelozup/horusec-admin/internal/logger"

	"github.com/tiagoangelozup/horusec-admin/internal/http/router"
	"github.com/tiagoangelozup/horusec-admin/internal/server"
)

func main() {
	log := logger.WithPrefix("main")

	r, err := router.New()
	if err != nil {
		log.WithError(err).Fatal("failed to create HTTP request router")
	}

	srv := server.New(r).Start()

	waitForInterruptSignal()
	if err = srv.GracefullyShutdown(); err != nil {
		log.WithError(err).Fatal("server forced to shutdown")
	}

	log.Info("server exiting")
}

func waitForInterruptSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
