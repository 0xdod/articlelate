package main

import (
	"github.com/fibreactive/articlelate/handler"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, handler *handler.Handler) {

	//PRIVATE ROUTES
	router.Use(handler.AttachUser())
	private := router.Group("")
	private.Use(handler.Authorize())
	articlePrivate := private.Group("/article")
	userPrivate := private.Group("/u")
	commentPriv := private.Group("/comment")
	{
		articlePrivate.GET("/", handler.CreateArticle)
		articlePrivate.POST("/", handler.CreateArticle)
		articlePrivate.POST("/:article_id/comment", handler.CreateComment)
		articlePrivate.POST("/:article_id/delete", handler.DeleteArticle)
		articlePrivate.GET("/:article_id/edit", handler.UpdateArticle)
		articlePrivate.POST("/:article_id/edit", handler.UpdateArticle)
		articlePrivate.POST("/:article_id/like", handler.LikeArticle)
		commentPriv.POST("/:comment_id/delete", handler.DeleteComment)
		commentPriv.POST("/:comment_id/edit", handler.UpdateComment)
		commentPriv.POST("/:comment_id/like", handler.LikeComment)
		userPrivate.POST("/logout", handler.Logout)
	}

	//PUBLIC ROUTES
	pub := router.Group("")
	a := pub.Group("/article")
	u := pub.Group("/u")
	{
		pub.GET("/", handler.ArticleList)
		a.GET("/:article_id/view", handler.ArticleDetail)
		u.GET("/register", handler.Register)
		u.POST("/register", handler.Register)
		u.GET("/login", handler.Login)
		u.POST("/login", handler.Login)
	}

	router.NoRoute(handler.NotFound)
}
