package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("auth")
		if err != nil || token == "" {
			c.SetCookie("auth", "", -1, "/", "", false, true)
			c.Redirect(303, "/u/login")
		} else {
			user := h.us.FindByToken(token)
			c.Set("user", user)
			c.Set("token", token)
		}
		c.Next()
	}
}
