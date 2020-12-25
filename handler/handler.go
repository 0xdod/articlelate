package handler

import (
	"net/http"

	"github.com/fibreactive/articlelate/service"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var dh *Handler

// move secret to env variable
var cookieStore = cookie.NewStore([]byte("hello"))

//var c = mgm.CollectionByName("sessions").Collection

//var _ = mongo.NewStore(c, 3600, true, []byte("secret"))

type Handler struct {
	us service.UserService
	ps service.PostService
	cs service.CommentService
}

func NewHandler() *Handler {
	us := service.NewUserStore()
	ps := service.NewPostMongo()
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
