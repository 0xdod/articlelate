package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/fibreactive/articlelate/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *Handler) PostList(c *gin.Context) {
	var posts []*models.Post
	// if s in query string
	// find in db posts with string s
	search := c.Query("s")
	if search != "" {
		filter := bson.M{"$text": bson.M{"$search": search}}
		posts = h.ps.Filter(filter)
	} else {
		posts = h.ps.GetAll()
	}
	var page *Page
	paginator := NewPaginator(posts, 5)
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
	render(c, http.StatusOK, "post_list.html", gin.H{"page": page, "search": search})
}

func (h *Handler) PostDetail(c *gin.Context) {
	slug := c.Param("slug")
	u := c.Param("u")
	filter := bson.M{"author.username": u, "slug": slug}
	post := h.ps.Get(filter)
	status := http.StatusOK
	templateName := "post_detail.html"
	if post == nil {
		status = http.StatusNotFound
		templateName = "404.html"
	} else {
		post.Comments = h.cs.GetByPost(post)
	}
	render(c, status, templateName, gin.H{"post": post})
}

func (h *Handler) CreatePost(c *gin.Context) {
	user := getUserFromContext(c)
	if c.Request.Method == "POST" {
		var req PostForm
		if err := Bind(c, &req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		post := models.NewPost(user, req.Title, req.Content)
		if err := h.ps.Create(post); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		//@TODO comeback
		c.Redirect(http.StatusSeeOther, post.GetAbsoluteURL())
		return
	}
	render(c, http.StatusOK, "create_post.html", gin.H{})
}

func (h *Handler) DeletePost(c *gin.Context) {
	slug := c.Param("slug")
	u := c.Param("u")
	filter := bson.M{"author.username": u, "slug": slug}
	post := h.ps.Get(filter)
	user := getUserFromContext(c)
	if user == nil || user.ID != post.Author.ID {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	err := h.ps.Delete(post.ID)
	if err != nil {
		render(c, http.StatusNotFound, "404.html", gin.H{"post": post})
		return
	}
	c.Redirect(303, "/")
}

func (h *Handler) UpdatePost(c *gin.Context) {
	var Err error
	slug := c.Param("slug")
	u := c.Param("u")
	filter := bson.M{"author.username": u, "slug": slug}
	post := h.ps.Get(filter)
	user := getUserFromContext(c)
	if user == nil || user.ID != post.Author.ID {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if c.Request.Method == "POST" {
		var req PostForm
		if err := Bind(c, &req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		post.Title = req.Title
		post.Content = req.Content
		err := h.ps.Update(post)
		if err == nil {
			c.Redirect(303, post.GetAbsoluteURL())
			return
		}
		Err = errors.New("Internal server error")
	}
	render(c, http.StatusOK, "edit_post.html", gin.H{
		"post":  post,
		"error": Err,
	})
}

func (h *Handler) LikePost(c *gin.Context) {
	var req LikeRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	postID := req.PostID
	post := h.ps.GetByID(postID)
	user := getUserFromContext(c)
	if user == nil || post == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	switch ParseAction(req.Action) {
	case Like:
		user.Like(post)
	case Unlike:
		user.Unlike(post)
	default:
		c.JSON(http.StatusForbidden, gin.H{"error": "Error processing request"})
		return
	}
	h.ps.Update(post)
	c.JSON(http.StatusOK, gin.H{"likes": len(post.Likes)})
}
