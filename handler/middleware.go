package handler

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Sessions(name string) gin.HandlerFunc {
	return sessions.Sessions(name, cookieStore)
}

func (h *Handler) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		var auth string
		a := session.Get("auth")
		auth, ok := a.(string)
		if ok {
			auth = a.(string)
			user := h.us.FindByToken(auth)
			c.Set("user", user)
		} else {
			c.Redirect(http.StatusSeeOther, "/u/login")
		}
		c.Next()
	}
}
