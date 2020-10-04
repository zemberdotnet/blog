package web

import (
	"github.com/gin-gonic/gin"
	"github.com/semihalev/gin-stats"
)

func (app *application) routes() *gin.Engine {
	r := gin.New()

	// Loading static files
	r.Static("/static/", "./ui/static/")

	// Creating out subgroups
	app.website(r)

	return r
}
func (app *application) website(r *gin.Engine) {
	// this may be bad
	web := r.Group("")

	// consideration
	// should we put stats before or after verification
	web.Use(stats.RequestStats())
	web.Use(app.secureHeaders)

	web.GET("/", app.home)
	web.GET("/bookshelf", app.bookshelf)
	web.GET("/resume", app.resume)
	web.GET("/writing", app.writing)
}
