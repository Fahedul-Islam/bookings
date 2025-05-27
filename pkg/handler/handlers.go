package handler

import (
	"net/http"

	"github.com/fahedul-islam/bookings/pkg/config"
	"github.com/fahedul-islam/bookings/pkg/models"
	"github.com/fahedul-islam/bookings/pkg/render"
)

var Repo *Reposotory
type Reposotory struct {
	App *config.AppConfig
}

//create a new reposotory 
func NewRepo(a *config.AppConfig) *Reposotory{
	return &Reposotory{
		App :a,
	}
}

// sets the reposotory for the handler
func NewHandler(r *Reposotory) {
	Repo=r
}

func (m *Reposotory)Home(w http.ResponseWriter, r *http.Request){
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(),"remote_ip",remoteIP)

	data:= make(map[string]string)
	data["test"]="got the data"
	data["remote_ip"]=remoteIP
	render.RenederTemplate(w,"home.html", &models.TemplateData{
		StringMap: data,
	})
}

func (m *Reposotory)About(w http.ResponseWriter, r *http.Request){
	render.RenederTemplate(w,"about.html",&models.TemplateData{

	})
}
func (m *Reposotory)Generals(w http.ResponseWriter, r *http.Request){
	render.RenederTemplate(w,"generals.html",&models.TemplateData{

	})
}