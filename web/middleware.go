package web

import (
	"github.com/gin-gonic/gin"
)

// secureHeaders adds basic security headers to the response
func (app *application) secureHeaders(c *gin.Context) {
	c.Writer.Header().Set("X-XSS-Protection", "1: mode=block")
	c.Writer.Header().Set("X-Frame-Options", "deny")
	// next
}
