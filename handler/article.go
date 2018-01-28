package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/fibreactive/articlelate/models"
	"github.com/gin-gonic/gin"
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
	h.render(http.StatusOK, c, gin.H{
		"title": "Home",
		"page":  page,
	}, "index.html")
}

func (h *Handler) ArticleDetail(c *gin.Context) {
	articleID := c.Param("article_id")
	article := h.as.GetByID(articleID)
	status := 200
	templateName := "article.html"
	var title string
	if article == nil {
		title = "Content not found"
		status = http.StatusNotFound
		templateName = "404.html"
	} else {
		title = article.Title
		article.Comments = h.cs.GetByArticle(article)
	}
	h.render(status, c, gin.H{
		"title":   title,
		"payload": article,
	}, templateName)
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
		//TODO comeback
		c.Redirect(http.StatusSeeOther, "/article/"+article.ID.Hex()+"/view")
		return
	}
	h.render(http.StatusOK, c, gin.H{
		"title":   "Create New Article",
		"payload": "",
	}, "create_article.html")
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
	var Err error
	articleID := c.Param("article_id")
	article := h.as.GetByID(articleID)
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
			c.Redirect(303, "/article/"+articleID+"/view")
			return
		}
		Err = errors.New("Internal server error")
	}
	h.render(http.StatusOK, c, gin.H{
		"title":   article.Title,
		"payload": article,
		"error":   Err,
	}, "edit_article.html")
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
