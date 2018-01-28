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
	if data != nil {
		data["user"] = getUserFromContext(c)
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
