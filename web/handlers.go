package web

import (
	"github.com/gin-gonic/gin"
)

func (app *application) Test(c *gin.Context) {

	c.SecureJSON(200, gin.H{
		"message": "Hello",
	})
}
