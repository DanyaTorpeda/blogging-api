package handler

import (
	"blogging_app/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		articles := api.Group("/articles")
		{
			articles.POST("/", h.CreateArticle)
			articles.GET("/", h.GetAllArticles)
			articles.GET("/:id", h.GetArticleByID)
			articles.PUT("/:id", h.UpdateArticle)
			articles.DELETE("/:id", h.DeleteArticle)
		}
	}

	return router
}
