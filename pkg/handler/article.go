package handler

import (
	blogging "blogging_app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateArticle(c *gin.Context) {
	var input blogging.Article

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Article.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetAllArticles(c *gin.Context) {
	articles, err := h.service.Article.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, articles)
}

func (h *Handler) GetArticleByID(c *gin.Context) {

}

func (h *Handler) UpdateArticle(c *gin.Context) {

}

func (h *Handler) DeleteArticle(c *gin.Context) {

}
