package web

import (
	"github.com/gin-contrib/multitemplate"
	"html/template"
	"path/filepath"
)

func newTemplateRender(dir string) multitemplate.Render {
	r := multitemplate.New()

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		panic(err)
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			panic(err)
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			panic(err)
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			panic(err)
		}
		r.Add(name, ts)

	}
	return r

}
