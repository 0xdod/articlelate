package handler

import (
	"errors"
	"net/http"

	"github.com/fibreactive/articlelate/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Signup(c *gin.Context) {
	if c.Request.Method == "POST" {
		var req SignUpForm
		if err := Bind(c, &req); err != nil {
			c.AbortWithStatus(422)
			return
		}
		user := models.NewUser(req.Username, req.Email, req.Password)
		if err := h.us.Create(user); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.Redirect(http.StatusSeeOther, "/u/login")
		return
	}
	user := getUserFromContext(c)
	if user != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	render(c, http.StatusOK, "register.html", gin.H{})
}

func (h *Handler) Login(c *gin.Context) {
	var Err error
	if c.Request.Method == "POST" {
		var req LoginForm
		if err := Bind(c, &req); err != nil {
			Err = errors.New("Login details incorrect try again")
		}
		user := h.us.Authenticate(req.Login, req.Password)
		if user == nil {
			Err = errors.New("Login details incorrect try again")
		} else {
			token, err := h.us.GenerateAuthToken(user)
			if err == nil {
				c.SetCookie("auth", token, 360000, "/", "", false, true)
				c.Redirect(http.StatusSeeOther, "/")
				return
			}
		}
	}
	user := getUserFromContext(c)
	if user != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	render(c, http.StatusOK, "login.html", gin.H{"error": Err})
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
	c.SetCookie("auth", "", -1, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/")
}

//TODO routes:
//edit (user profile), delete (user account)
