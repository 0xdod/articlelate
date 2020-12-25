package handler

import (
	"errors"
	"net/http"

	"github.com/fibreactive/articlelate/models"
	"github.com/gin-contrib/sessions"
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
			session := sessions.Default(c)
			token, err := h.us.GenerateAuthToken(user)
			if err == nil {
				session.Set("auth", token)
				err := session.Save()
				if err != nil {
					c.AbortWithStatus(http.StatusInternalServerError)
					return
				}
				c.Redirect(http.StatusSeeOther, "/")
				return
			}
		}
	}
	user := getUserFromContext(c)
	if user != nil {
		// display flash message that user is logged in
		// asking them to go to home or to logout
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	render(c, http.StatusOK, "login.html", gin.H{"error": Err})
}

func (h *Handler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Redirect(http.StatusSeeOther, "/")
}

//TODO routes:
//edit (user profile), delete (user account)
