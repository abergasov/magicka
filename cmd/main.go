package main

import (
	"flag"
	"fmt"
	"magicka/internal/logger"
	"magicka/internal/routes"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

var appPort = flag.String("port", "8080", "App listen port")

func main() {
	flag.Parse()
	log, err := logger.NewAppLogger()
	if err != nil {
		println(fmt.Sprintf("error init logger: %s", err))
		return
	}
	app := routes.InitAppRouter(*appPort, log)
	go func() {
		log.Info("starting service", zap.String("port", *appPort))
		if err = app.Run(); err != nil {
			log.Fatal("error start service", err)
		}
	}()

	// register app shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c // This blocks the main thread until an interrupt is received
	if err = app.Shutdown(); err != nil {
		log.Fatal("error shutdown service", err)
	}
	log.Info("service was successful shutdown")
}
