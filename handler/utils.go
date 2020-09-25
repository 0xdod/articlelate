package handler

import (
	"github.com/fibreactive/articlelate/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func (h *Handler) render(code int, c *gin.Context, data gin.H, templateName string) {

	// add user model
	if data != nil {
		token, _ := c.Cookie("auth")
		data["user"] = h.RetrieveUser(token)
	}

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(code, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(code, data["payload"])
	default:
		// Respond with HTML
		c.HTML(code, templateName, data)
	}

}

//Helper func to help choose which func to run based on the request method
//postFUnc handles POST request, similarly GetFunc, GET request
func resolveGetOrPost(c *gin.Context, getFunc, postFunc func()) {
	switch c.Request.Method {
	case "GET":
		getFunc()
	case "POST":
		postFunc()
	}
}

func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

//retrieve user from
func getUserFromContext(c *gin.Context) *models.User {
	u, exists := c.Get("user")
	if !exists {
		return nil
	}
	user := u.(*models.User)
	return user
}

func (h *Handler) RetrieveUser(cookie string) *models.User {
	return h.us.FindByToken(cookie)
}
