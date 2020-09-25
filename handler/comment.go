package handler

import (
	"net/http"

	"github.com/fibreactive/articlelate/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateComment(c *gin.Context) {
	articleId := c.Param("article_id")
	user := getUserFromContext(c)
	if user == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	var form CommentForm
	if err := Bind(c, &form); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	article := h.as.GetByID(articleId)
	comment := models.NewComment(user, article, form.Content)
	err := h.cs.Create(comment)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Redirect(http.StatusSeeOther, "/article/"+article.ID.Hex()+"/view")
}
