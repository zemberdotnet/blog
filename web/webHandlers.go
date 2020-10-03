package web

import (
	"github.com/gin-gonic/gin"
)

// bad lets mote

func (app *application) home(c *gin.Context) {
	c.HTML(200, "home.page.tmpl", nil)
}
