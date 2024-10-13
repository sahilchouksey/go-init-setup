package api

import (
	"github.com/sahilchouksey/go-init-setup/database"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2/log"
)

type APIServer struct {
	app           *gin.Engine
	listenAddress string
	store         database.Storage
}

func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		app:           gin.New(),
		listenAddress: listenAddress,
	}
}

func (s *APIServer) GetEngine() *gin.Engine {
	return s.app
}

func (s *APIServer) Run() error {
	log.Info("Starting API Server")
	log.Infof("Listening on %s", s.listenAddress)

	return s.app.Run(":" + s.listenAddress)
}
