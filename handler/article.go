package handler

import (
	"net/http"
	"strconv"

	"github.com/fibreactive/articlelate/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Index(c *gin.Context) {
	articles := h.as.GetAll()
	if articles == nil {
		c.Header("Content-Type", "text/html")
		c.String(http.StatusInternalServerError, "<h1>Internal server error</h1>")
		return
	}
	paginator := NewPaginator(articles, 5)
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = minPage
	}
	posts, err := paginator.Page(page)
	if err == EmptyPage {
		posts, _ = paginator.Page(paginator.MaxPage)
	}
	h.render(http.StatusOK, c, gin.H{
		"title": "Home",
		"page":  posts,
	}, "index.html")
}

func (h *Handler) GetArticle(c *gin.Context) {
	articleID := c.Param("article_id")
	article := h.as.GetByID(articleID)
	if article == nil {
		h.render(http.StatusNotFound, c, gin.H{
			"title":   "Oops",
			"payload": article,
		}, "404.html")
		return
	}
	article.Comments = h.cs.GetByArticle(article)
	h.render(http.StatusOK, c, gin.H{
		"title":   article.Title,
		"payload": article,
	}, "article.html")
}

func (h *Handler) CreateArticle(c *gin.Context) {
	getFunc := func() {
		h.render(http.StatusOK, c, gin.H{
			"title":   "Create New Article",
			"payload": "",
		}, "create_article.html")
	}
	postFunc := func() {
		var req ArticleForm
		if err := Bind(c, &req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		user := getUserFromContext(c)
		article := models.NewArticle(user, req.Title, req.Content)
		if err := h.as.Create(article); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		//TODO comeback
		c.Redirect(http.StatusSeeOther, "/article/"+article.ID.Hex()+"/view")
	}
	resolveGetOrPost(c, getFunc, postFunc)
}

func (h *Handler) DeleteArticle(c *gin.Context) {
	articleID := c.Param("article_id")
	article := h.as.GetByID(articleID)
	user := getUserFromContext(c)
	if user == nil || user.ID != article.Author.ID {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	err := h.as.Delete(articleID)
	if err != nil {
		h.render(http.StatusNotFound, c, gin.H{
			"title":   "Oops",
			"payload": article,
		}, "404.html")
		return
	}
	c.Redirect(303, "/")
}

func (h *Handler) UpdateArticle(c *gin.Context) {
	articleID := c.Param("article_id")
	article := h.as.GetByID(articleID)
	user := getUserFromContext(c)
	if user == nil || user.ID != article.Author.ID {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	getFunc := func() {
		h.render(http.StatusOK, c, gin.H{
			"title":   article.Title,
			"payload": article,
		}, "edit_article.html")
	}
	postFunc := func() {
		var req ArticleForm
		if err := Bind(c, &req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		article.Title = req.Title
		article.Content = req.Content
		err := h.as.Update(article)
		if err != nil {
			c.String(http.StatusInternalServerError, "<h1>Internal Server Error</h1>")
			return
		}
		c.Redirect(303, "/article/"+articleID+"/view")
	}
	resolveGetOrPost(c, getFunc, postFunc)
}

func (h *Handler) LikeArticle(c *gin.Context) {
	var req LikeRequest
	articleID := c.Param("article_id")
	article := h.as.GetByID(articleID)
	user := getUserFromContext(c)
	if user == nil || article == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	switch ParseAction(req.Action) {
	case Like:
		user.LikeArticle(article)
	case Unlike:
		user.UnlikeArticle(article)
	default:
		c.JSON(http.StatusForbidden, gin.H{"error": "Error processing request"})
		return
	}

	h.as.Update(article)
	c.JSON(http.StatusOK, gin.H{"likes": len(article.Likes)})
}
