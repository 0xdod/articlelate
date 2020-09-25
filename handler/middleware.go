package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("auth")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		user := h.us.FindByToken(token)
		if user == nil {
			c.SetCookie("auth", "", 0, "/", "localhost", false, true)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("user", user)
		c.Set("token", token)
		c.Next()
	}
}
