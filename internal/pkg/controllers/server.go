package controllers

import (
	"RolePlayModule/internal/pkg/storage/pg"
	"RolePlayModule/internal/utils/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

func (s *Server) InitSwagger() {
	swagURL := s.cfg.NgrokUrl + "/swagger/doc.json"
	url := ginSwagger.URL(swagURL)
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func (s *Server) initRoutes() {
	auth := s.router.Group("/auth")
	{
		auth.POST("/check-user", s.CheckUser)          //TODO Докрутить реализацию через бд
		auth.POST("/check-code", s.CheckCode)          //TODO Докрутить реализацию через бд
		auth.POST("/send-code-again", s.SendCodeAgain) //TODO Докрутить реализацию через бд
		auth.POST("/fill-profile", s.FillProfile)      //TODO Докрутить реализацию через бд
	}

	product := s.router.Group("/product")
	{
		product.POST("/", s.CreateProduct)   //TODO ДОКРУТИТЬ
		product.GET("/", s.GetProducts)      //TODO ДОКРУТИТЬ
		product.PUT("/", s.UpdateProduct)    //TODO ДОКРУТИТЬ
		product.DELETE("/", s.DeleteProduct) //TODO ДОКРУТИТЬ
	}

	profile := s.router.Group("/profile")
	{
		profile.GET("/", s.GetProfileInfo) //TODO ДОКРУТИТЬ
		profile.PUT("/", s.EditProfile)    //TODO ДОКРУТИТЬ
		profile.DELETE("/", s.DeleteUser)  //TODO ДОКРУТИТЬ
	}

}

func (s *Server) Run(addr string) error { return s.router.Run(addr) }
