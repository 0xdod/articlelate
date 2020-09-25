package handler

import (
	"net/http"

	"github.com/fibreactive/articlelate/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	getFunc := func() {
		h.render(http.StatusOK, c, nil, "register.html")
	}
	postFunc := func() {
		var req SignUpForm
		if err := Bind(c, &req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		user := models.NewUser(req.Username, req.Email, req.Password)
		if err := h.us.Create(user); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		token, err := h.us.GenerateAuthToken(user)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.SetCookie("auth", token, 0, "/", "localhost", false, true)
		c.Redirect(http.StatusSeeOther, "/")

	}
	resolveGetOrPost(c, getFunc, postFunc)
}

func (h *Handler) Login(c *gin.Context) {
	getFunc := func() {
		h.render(http.StatusOK, c, nil, "login.html")
	}
	postFunc := func() {
		var req LoginForm
		if err := Bind(c, &req); err != nil {
			c.SetCookie("message", "Login details incorrect try again", 1, "/", "localhost", false, false)
			c.Redirect(http.StatusSeeOther, "/u/login")
			return
		}
		user := h.us.Authenticate(req.Login, req.Password)
		if user == nil {
			c.SetCookie("message", "Login details incorrect try again", 1, "/", "localhost", false, false)
			c.Redirect(http.StatusSeeOther, "/u/login")
			return
		}
		token, err := h.us.GenerateAuthToken(user)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.SetCookie("auth", token, 0, "/", "localhost", false, true)
		c.Redirect(http.StatusSeeOther, "/")
	}
	resolveGetOrPost(c, getFunc, postFunc)
}

func (h *Handler) Logout(c *gin.Context) {
	token, err := c.Cookie("auth")
	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	if err := h.us.RemoveAuthToken(token); err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.SetCookie("auth", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/")
}

//TODO routes:
//edit (user profile), delete (user account)
