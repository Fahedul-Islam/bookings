package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/fahedul-islam/bookings/pkg/config"
	"github.com/fahedul-islam/bookings/pkg/models"
	"github.com/justinas/nosurf"
)
var app *config.AppConfig
func NewTemplate(a *config.AppConfig) {
	app=a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken=nosurf.Token(r)
	return td
}

func RenederTemplate(w http.ResponseWriter,r *http.Request, tmpl string, td *models.TemplateData){
	var tc map[string]*template.Template
	if app.UseCache {
		tc=app.TemplateCache
	} else {
		tc,_=CreateTemplateCache()
	}
	t,ok :=tc[tmpl]
	if !ok {
		log.Fatal(ok)
	}
	td=AddDefaultData(td,r)
	er := t.Execute(w,td)
	if er!=nil {
		log.Fatal(er)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error){
	mycache :=map[string]*template.Template{}

	pages,err :=filepath.Glob("./template/*.html")
	if err!=nil {
		return mycache,err
	}

	for _,page := range pages {
		name:= filepath.Base(page)
		ts, er := template.New(name).ParseFiles("./template/base.html", page)

		if er!=nil {
			return mycache,er
		}
		mycache[name]=ts
	}

	return mycache, nil
}