package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("auth")
		user := getUserFromContext(c)
		if user == nil || err != nil {
			c.SetCookie("auth", "", -1, "/", "", false, true)
			c.Redirect(303, "/u/login")
		} else {
			c.Set("user", user)
			c.Set("token", token)
		}
		c.Next()
	}
}

func (h *Handler) AttachUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("auth")
		user := h.us.FindByToken(token)
		c.Set("user", user)
		c.Next()
	}
}
