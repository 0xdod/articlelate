package handler

import (
	"net/http"

	"github.com/fibreactive/articlelate/service"
	"github.com/gin-gonic/gin"
)

var dh *Handler

type Handler struct {
	us service.UserService
	as service.ArticleService
	cs service.CommentService
}

func NewHandler() *Handler {
	us := service.NewUserStore()
	as := service.NewArticleStore()
	cs := service.NewCommentStore()
	dh = &Handler{
		us: us,
		as: as,
		cs: cs,
	}
	return dh
}

func (h *Handler) NotFound(c *gin.Context) {
	render(c, http.StatusNotFound, "404.html", gin.H{})
}
