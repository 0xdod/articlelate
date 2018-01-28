package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/fibreactive/articlelate/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *Handler) ArticleList(c *gin.Context) {
	articles := h.as.GetAll()
	var page *Page
	paginator := NewPaginator(articles, 5)
	if paginator == nil {
		page = nil
	} else {

		pageNo, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil {
			pageNo = minPage
		}
		page, err = paginator.Page(pageNo)
		if err == EmptyPage {
			page, _ = paginator.Page(paginator.MaxPage)
		}
	}
	render(c, http.StatusOK, "post_list.html", gin.H{"page": page})
}

func (h *Handler) ArticleDetail(c *gin.Context) {
	slug := c.Param("slug")
	u := c.Param("u")
	filter := bson.M{"author.username": u, "slug": slug}
	article := h.as.Get(filter)
	status := http.StatusOK
	templateName := "post_detail.html"
	if article == nil {
		status = http.StatusNotFound
		templateName = "404.html"
	} else {
		article.Comments = h.cs.GetByArticle(article)
	}
	render(c, status, templateName, gin.H{"post": article})
}

func (h *Handler) CreateArticle(c *gin.Context) {
	user := getUserFromContext(c)
	if c.Request.Method == "POST" {
		var req ArticleForm
		if err := Bind(c, &req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		article := models.NewArticle(user, req.Title, req.Content)
		if err := h.as.Create(article); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		//@TODO comeback
		c.Redirect(http.StatusSeeOther, article.GetAbsoluteURL())
		return
	}
	render(c, http.StatusOK, "create_post.html", gin.H{})
}

func (h *Handler) DeleteArticle(c *gin.Context) {
	slug := c.Param("slug")
	u := c.Param("u")
	filter := bson.M{"author.username": u, "slug": slug}
	article := h.as.Get(filter)
	user := getUserFromContext(c)
	if user == nil || user.ID != article.Author.ID {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	err := h.as.Delete(article.ID)
	if err != nil {
		render(c, http.StatusNotFound, "404.html", gin.H{"post": article})
		return
	}
	c.Redirect(303, "/")
}

func (h *Handler) UpdateArticle(c *gin.Context) {
	var Err error
	slug := c.Param("slug")
	u := c.Param("u")
	filter := bson.M{"author.username": u, "slug": slug}
	article := h.as.Get(filter)
	user := getUserFromContext(c)
	if user == nil || user.ID != article.Author.ID {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if c.Request.Method == "POST" {
		var req ArticleForm
		if err := Bind(c, &req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		article.Title = req.Title
		article.Content = req.Content
		err := h.as.Update(article)
		if err == nil {
			c.Redirect(303, article.GetAbsoluteURL())
			return
		}
		Err = errors.New("Internal server error")
	}
	render(c, http.StatusOK, "edit_post.html", gin.H{
		"post":  article,
		"error": Err,
	})
}

func (h *Handler) LikeArticle(c *gin.Context) {
	var req LikeRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	articleID := req.ArticleID
	article := h.as.GetByID(articleID)
	user := getUserFromContext(c)
	if user == nil || article == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	switch ParseAction(req.Action) {
	case Like:
		user.Like(article)
	case Unlike:
		user.Unlike(article)
	default:
		c.JSON(http.StatusForbidden, gin.H{"error": "Error processing request"})
		return
	}
	h.as.Update(article)
	c.JSON(http.StatusOK, gin.H{"likes": len(article.Likes)})
}
