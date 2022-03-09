package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/zserge/lorca"
)

func main() {
	go func() {
		gin.SetMode(gin.DebugMode)
		router := gin.Default()
		router.GET("/", func(c *gin.Context) {
			c.Writer.Write([]byte("hello go"))
		})
		router.Run(":8080")
	}()
	var ui lorca.UI
	ui, _ = lorca.New("http://localhost:8080", "", 800, 600, "--disable-sync", "--disable-translate")
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ui.Done():
	case <-chSignal:
	}
	ui.Close()
}
