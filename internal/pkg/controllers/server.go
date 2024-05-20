package controllers

import (
	"RolePlayModule/internal/pkg/storage/pg"
	"RolePlayModule/internal/utils/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router  *gin.Engine
	storage *pg.Storage
	cfg     *config.Config
}

func NewServer(storage *pg.Storage, cfg *config.Config) *Server {
	router := gin.Default()
	server := &Server{
		router:  router,
		storage: storage,
		cfg:     cfg,
	}
	server.initRoutes()
	return server
}

func (s *Server) initRoutes() {

}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
