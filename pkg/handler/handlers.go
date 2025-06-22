package handler

import (
	"encoding/json"
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
	render.RenederTemplate(w,r,"home.html", &models.TemplateData{
		StringMap: data,
	})
}

func (m *Reposotory)About(w http.ResponseWriter, r *http.Request){
	render.RenederTemplate(w,r,"about.html",&models.TemplateData{

	})
}
// Generals is the handler for the generals quarters page
// it renders the generals-quarters template
// this is where the user can see the generals quarters
func (m *Reposotory)Generals(w http.ResponseWriter, r *http.Request){
	render.RenederTemplate(w,r,"generals.html",&models.TemplateData{

	})
}
// Major is the handler for the majors page
// it renders the majors template
// this is where the user can see the majors
func (m *Reposotory)Major(w http.ResponseWriter, r *http.Request){
	render.RenederTemplate(w,r,"majors.html",&models.TemplateData{
	})
}
// Availability is the handler for the search availability page
// it renders the search-availability template
// it is used to search for available rooms
func (m *Reposotory)Availability(w http.ResponseWriter, r *http.Request){
	render.RenederTemplate(w,r,"search-availability.html",&models.TemplateData{
	})
}

func (m *Reposotory)PostAvailability(w http.ResponseWriter, r *http.Request){
	start:=r.FormValue("start")
	end:=r.FormValue("end")
	w.Write([]byte("Posted the availability from " + start + " to " + end))
}

type JsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Reposotory)AvailabilityJson(w http.ResponseWriter, r *http.Request){
	resp := JsonResponse{
		OK: true,
		Message: "Available",
	}
	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
// MakeReservation is the handler for the make reservation page
// it renders the make-reservation template
// this is where the user can make a reservation
func (m *Reposotory)MakeReservation(w http.ResponseWriter, r *http.Request){
	render.RenederTemplate(w,r,"make-reservation.html",&models.TemplateData{
	})
}
// Contact is the handler for the contact page
// it renders the contact template
// this is where the user can contact us
func (m *Reposotory)Contact(w http.ResponseWriter, r *http.Request){
	render.RenederTemplate(w,r,"contact.html",&models.TemplateData{
	})
}