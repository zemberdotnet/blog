package web

import (
	"github.com/gin-gonic/gin"
	"github.com/semihalev/gin-stats"
	"path/filepath"
)

func (app *application) routes() *gin.Engine {
	r := gin.New()

	fp, err := filepath.Abs("./ui/html/")
	if err != nil {
		// if we don't find the filepath we aren't going to get very far
		panic(err)
	}
	// Loading templates
	r.LoadHTMLGlob(filepath.Join(fp, "*.tmpl"))

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
}
