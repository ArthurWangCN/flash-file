package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ArthurWangCN/flash-files/server"
	"github.com/zserge/lorca"
)

func main() {
	go server.Run()
	startBrowser()
}

func startBrowser() {
	var ui lorca.UI
	ui, _ = lorca.New("http://127.0.0.1:8080/static/index.html", "", 800, 600, "--disable-sync", "--disable-translate")
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ui.Done():
	case <-chSignal:
	}
	ui.Close()
}
