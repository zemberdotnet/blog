package web

import ()

type application struct {
}

func Entry() {

	app := &application{}
	r := app.routes()

	renderer := newTemplateRender("./ui/html/")
	r.HTMLRender = renderer

	r.Run(":1337")

}
