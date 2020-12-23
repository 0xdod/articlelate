package main

import (
	"github.com/fibreactive/articlelate/handler"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, handler *handler.Handler) {
	//PRIVATE ROUTES
	private := router.Group("")
	private.Use(handler.Authorize())
	postPrivate := private.Group("/p")
	userPrivate := private.Group("/u")
	commentPrivate := private.Group("/comment")
	{
		postPrivate.GET("/", handler.CreatePost)
		postPrivate.POST("/", handler.CreatePost)
		postPrivate.POST("/:u/:slug/comment", handler.CreateComment)
		postPrivate.POST("/:u/:slug/delete", handler.DeletePost)
		postPrivate.GET("/:u/:slug/edit", handler.UpdatePost)
		postPrivate.POST("/:u/:slug/edit", handler.UpdatePost)
		postPrivate.POST("/:u/:slug/like", handler.LikePost)
		commentPrivate.POST("/:comment_id/delete", handler.DeleteComment)
		commentPrivate.POST("/:comment_id/edit", handler.UpdateComment)
		commentPrivate.POST("/:comment_id/like", handler.LikeComment)
		userPrivate.POST("/logout", handler.Logout)
	}
	//PUBLIC ROUTES
	pub := router.Group("")
	p := pub.Group("/p")
	u := pub.Group("/u")
	{
		pub.GET("/", handler.PostList)
		p.GET("/:u/:slug", handler.PostDetail)
		u.GET("/register", handler.Signup)
		u.POST("/register", handler.Signup)
		u.GET("/login", handler.Login)
		u.POST("/login", handler.Login)
	}
	router.NoRoute(handler.NotFound)
}
