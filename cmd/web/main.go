package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/fahedul-islam/bookings/pkg/config"
	"github.com/fahedul-islam/bookings/pkg/handler"
	"github.com/fahedul-islam/bookings/pkg/render"
)
var app config.AppConfig
const port=":8080"
var session *scs.SessionManager

func main(){

	app.InProduction =false
	session =scs.New()
	session.Lifetime = 24* time.Hour
	session.Cookie.Persist =true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure =app.InProduction

	app.Session=session
	tc,err:= render.CreateTemplateCache()

	if err!=nil {
		log.Fatal(err)
	}
	app.TemplateCache=tc
	app.UseCache =false
	
	render.NewTemplate(&app)

	repo:=handler.NewRepo(&app)
	handler.NewHandler(repo)

	fmt.Println("Server running on port ",port)
	
	srv:= &http.Server{
		Addr: port,
		Handler: routes(&app),
	}
	err=srv.ListenAndServe()
	log.Fatal(err)
}