package template

import (
	"fmt"
	"html/template"
	"path"
)

type TemplateLink struct {
	Title string
	Link  string
}

func NewTemplate(page string) *template.Template {
	templ := template.Must(template.ParseFiles("ui/base.html"))
    template.Must(templ.ParseFiles(path.Join("ui/pages/", fmt.Sprintf("%s.html", page))))
    template.Must(templ.ParseGlob("ui/partials/*.html"))

    return templ
}
