package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var build = "develop"

func main() {
	log.Println("starting service", build)
	defer log.Println("service ended")

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Println("stopping service")
}
