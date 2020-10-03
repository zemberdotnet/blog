package web

import ()

type application struct {
}

func Entry() {

	app := &application{}
	r := app.routes()

	r.Run(":1337")

}
