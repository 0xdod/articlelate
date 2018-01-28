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

func (h *Handler) UpdateComment(c *gin.Context) {
	id := c.Param("comment_id")
	comment := h.cs.GetByID(id)
	user := getUserFromContext(c)
	if user == nil || user.ID != comment.Author.ID {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if c.Request.Method == "POST" {
		var req CommentForm
		if err := Bind(c, &req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		comment.Content = req.Content
		if err := h.cs.Update(comment); err != nil {
			c.String(http.StatusInternalServerError, "<h1>Internal Server Error</h1>")
			return
		}
		c.Redirect(303, "/article/"+comment.Article.ID.Hex()+"/view")
	}
	h.render(http.StatusOK, c, gin.H{
		"title":   "Edit comment",
		"payload": comment,
	}, "edit_comment.html")
}

func (h *Handler) LikeComment(c *gin.Context) {
	var req LikeRequest
	id := c.Param("comment_id")
	comment := h.cs.GetByID(id)
	user := getUserFromContext(c)
	if user == nil || comment == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	switch ParseAction(req.Action) {
	case Like:
		user.Like(comment)
	case Unlike:
		user.Unlike(comment)
	default:
		c.JSON(http.StatusForbidden, gin.H{"error": "Error processing request"})
		return
	}
	h.cs.Update(comment)
	c.JSON(http.StatusOK, gin.H{"likes": len(comment.Likes)})
}

func (h *Handler) DeleteComment(c *gin.Context) {
	commentId := c.Param("comment_id")
	comment := h.cs.GetByID(commentId)
	user := getUserFromContext(c)
	if user == nil || user.ID != comment.Author.ID {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if err := h.cs.Delete(commentId); err != nil {
		h.render(http.StatusInternalServerError, c, gin.H{
			"title":   "Oops",
			"payload": nil,
		}, "404.html")
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, "/article"+comment.Article.ID.Hex()+"/view")
}
