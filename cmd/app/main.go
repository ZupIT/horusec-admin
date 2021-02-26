package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tiagoangelozup/horusec-admin/internal/http/router"
	"github.com/tiagoangelozup/horusec-admin/internal/server"
)

func main() {
	r, err := router.New()
	checkErr(err)

	srv := server.New(r).Start()

	waitForInterruptSignal()
	err = srv.GracefullyShutdown()
	if err != nil {
		log.Fatal(fmt.Errorf("server forced to shutdown: %w", err))
	}

	log.Println("server exiting")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func waitForInterruptSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
