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
		auth.POST("/check-code", s.CheckCode)          //TODO Докрутить jwt и swagger
		auth.POST("/send-code-again", s.SendCodeAgain) //TODO Докрутить swagger
		auth.POST("/fill-profile", s.FillProfile)      //TODO Докрутить jwt и swagger
	}

	category := s.router.Group("/category")
	{
		category.POST("", s.CreateCategory) //TODO Swagger
		category.GET("", s.GetCategories)   //TODO Swagger
	}

	product := s.router.Group("/product")
	{
		product.POST("", s.CreateProduct)   //TODO ДОКРУТИТЬ swagger
		product.GET("", s.GetProducts)      //TODO ДОКРУТИТЬ swagger
		product.PUT("", s.UpdateProduct)    //TODO ДОКРУТИТЬ swagger
		product.DELETE("", s.DeleteProduct) //TODO ДОКРУТИТЬ swagger
	}

	profile := s.router.Group("/profile")
	{
		profile.GET("/", s.GetProfileInfo) //TODO ДОКРУТИТЬ swagger
		profile.PUT("/", s.EditProfile)    //TODO ДОКРУТИТЬ swagger
		profile.DELETE("/", s.DeleteUser)  //TODO ДОКРУТИТЬ swagger
	}

	s.router.Static("/media", "./media") // ✔

}

func (s *Server) Run(addr string) error { return s.router.Run(addr) }
