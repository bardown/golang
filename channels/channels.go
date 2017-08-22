package main

import (
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Player struct {
	Name   string
	Number int
}

func main() {
	s := &http.Server()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go func() {
		<-sig
		time.Sleep(30 * time.Second)
		os.Exit(0)
	}()
	s.ListenAndServ()
}
