package handler

import (
	// "net/http"

	"github.com/Beksultan15/go_project/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
    return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000"}
	// router.Use(cors.New(config))

	// auth := router.Group("/auth")
	//  {
	// 	// auth.POST("/signup",h.signup)
	// 	// auth.POST("/signin",h.signin)
	// 	// auth.GET("/logout",)
	//  }

	//  search := router.Group("/products")
	//  {
	// 	search.GET("/search/",h.Search)
	//  }


	 return router
}