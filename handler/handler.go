package handler

import (
	"net/http"

	"github.com/fibreactive/articlelate/service"
	"github.com/gin-gonic/gin"
)

var dh *Handler

type Handler struct {
	us service.UserService
	ps service.PostService
	cs service.CommentService
}

func NewHandler() *Handler {
	us := service.NewUserStore()
	ps := service.NewPostStore()
	cs := service.NewCommentStore()
	dh = &Handler{
		us: us,
		ps: ps,
		cs: cs,
	}
	return dh
}

func (h *Handler) NotFound(c *gin.Context) {
	render(c, http.StatusNotFound, "404.html", gin.H{})
}
