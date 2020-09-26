package main

import (
	"github.com/fibreactive/articlelate/handler"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, handler *handler.Handler) {

	//PRIVATE ROUTES
	private := router.Group("")
	private.Use(handler.Authorize())
	articlePrivate := private.Group("/article")
	userPrivate := private.Group("/u")
	commentPriv := private.Group("/comment")
	{
		userPrivate.POST("/logout", handler.Logout)
		userPrivate.GET("/article/create", handler.CreateArticle)
		userPrivate.POST("/article/create", handler.CreateArticle)
		articlePrivate.POST("/:article_id/comment", handler.CreateComment)
		articlePrivate.POST("/:article_id/delete", handler.DeleteArticle)
		articlePrivate.GET("/:article_id/edit", handler.UpdateArticle)
		articlePrivate.POST("/:article_id/edit", handler.UpdateArticle)
		articlePrivate.POST("/:article_id/like", handler.LikeArticle)
		commentPriv.POST("/:comment_id/delete", handler.DeleteComment)
		commentPriv.POST("/:comment_id/edit", handler.UpdateComment)
		commentPriv.POST("/:comment_id/like", handler.LikeComment)
	}

	//PUBLIC ROUTES
	public := router.Group("")
	a := public.Group("/article")
	u := public.Group("/u")
	{
		a.GET("/:article_id/view", handler.GetArticle)
		u.GET("/register", handler.Register)
		u.POST("/register", handler.Register)
		u.GET("/login", handler.Login)
		u.POST("/login", handler.Login)
		public.GET("/", handler.Index)
	}

	router.NoRoute(handler.NotFound)
}
