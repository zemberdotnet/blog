package web

import (
	"github.com/gin-gonic/gin"
)

// bad lets mote

func (app *application) home(c *gin.Context) {
	c.HTML(200, "home.page.tmpl", nil)
}

func (app *application) bookshelf(c *gin.Context) {
	c.HTML(200, "bookshelf.page.tmpl", nil)
}

func (app *application) resume(c *gin.Context) {
	c.HTML(200, "resume.page.tmpl", nil)
}

func (app *application) writing(c *gin.Context) {
	c.HTML(200, "writing.page.tmpl", nil)
}

func (app *application) essay(c *gin.Context) {
	c.HTML(200, "curatedaesthetics.page.tmpl", gin.H{
		"article": true,
	})
}
