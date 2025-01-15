package webServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// WebServer general web server struct
type WebServer struct {
	engine *gin.Engine
	config Config
}

type Config struct {
	Address string
	Port    int
}

// New constructor from web server
func New(config Config) *WebServer {

	gin.SetMode(gin.DebugMode)

	engine := gin.Default()

	ws := WebServer{
		engine: engine,
		config: config,
	}

	return &ws
}

// Run web server
func (ws *WebServer) Run() error {
	address := fmt.Sprintf("%s:%d", ws.config.Address, ws.config.Port)
	err := ws.engine.Run(address)
	if err != nil {
		return err
	}
	return nil
}
