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
		auth.POST("/check-user", s.CheckUser)
		auth.POST("/check-password", s.CheckPassword)
		auth.POST("/check-code", s.CheckCode)
		auth.POST("/send-code-again", s.SendCodeAgain)
		auth.POST("/fill-profile", s.FillProfile)
	}

	category := s.router.Group("/category")
	{
		category.POST("", s.CreateCategory)
		category.GET("", s.GetCategories)
	}

	product := s.router.Group("/product")
	{
		product.POST("", s.CreateProduct)
		product.GET("/id", s.GetProductById)
		product.GET("", s.GetProducts)
		product.PUT("", s.UpdateProduct)    //TODO ДОКРУТИТЬ
		product.DELETE("", s.DeleteProduct) //TODO ДОКРУТИТЬ
	}

	profile := s.router.Group("/profile")
	{
		profile.GET("", s.GetProfileInfo)
		profile.PUT("", s.EditProfile)   //TODO ДОКРУТИТЬ swagger
		profile.DELETE("", s.DeleteUser) //TODO ДОКРУТИТЬ swagger
	}
	orders := s.router.Group("/orders")
	{
		orders.GET("", s.GetOrders)
		orders.POST("", s.CreateOrder)
	}

	s.router.Static("/media", "./media")

}

func (s *Server) Run(addr string) error { return s.router.Run(addr) }
