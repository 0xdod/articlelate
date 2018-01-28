package handler

import (
	"net/http"

	"github.com/fibreactive/articlelate/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	us service.UserService
	as service.ArticleService
	cs service.CommentService
}

func NewHandler() *Handler {
	us := service.NewUserStore()
	as := service.NewArticleStore()
	cs := service.NewCommentStore()
	return &Handler{
		us: us,
		as: as,
		cs: cs,
	}
}

func (h *Handler) NotFound(c *gin.Context) {
	h.render(http.StatusNotFound, c, gin.H{
		"title":   "Content not found",
		"payload": "",
	}, "404.html")
}
