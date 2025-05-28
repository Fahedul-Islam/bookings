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
// Generals is the handler for the generals quarters page
// it renders the generals-quarters template
// this is where the user can see the generals quarters
func (m *Reposotory)Generals(w http.ResponseWriter, r *http.Request){
	render.RenederTemplate(w,"generals.html",&models.TemplateData{

	})
}
// Major is the handler for the majors page
// it renders the majors template
// this is where the user can see the majors
func (m *Reposotory)Major(w http.ResponseWriter, r *http.Request){
	render.RenederTemplate(w,"majors.html",&models.TemplateData{
	})
}
// BookNow is the handler for the search availability page
// it renders the search-availability template
// it is used to search for available rooms
func (m *Reposotory)BookNow(w http.ResponseWriter, r *http.Request){
	render.RenederTemplate(w,"search-availability.html",&models.TemplateData{
	})
}
// MakeReservation is the handler for the make reservation page
// it renders the make-reservation template
// this is where the user can make a reservation
func (m *Reposotory)MakeReservation(w http.ResponseWriter, r *http.Request){
	render.RenederTemplate(w,"make-reservation.html",&models.TemplateData{
	})
}
// Contact is the handler for the contact page
// it renders the contact template
// this is where the user can contact us
func (m *Reposotory)Contact(w http.ResponseWriter, r *http.Request){
	render.RenederTemplate(w,"contact.html",&models.TemplateData{
	})
}