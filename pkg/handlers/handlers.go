package handlers

import (
	"fmt"
	"net/http"

	"github.com/rolandtannous/gobookings/models"
	"github.com/rolandtannous/gobookings/pkg/config"
	"github.com/rolandtannous/gobookings/pkg/render"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// The repository used by the handlers
var Repo *Repository

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// newHandlers sets the rpository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	remoteIP := r.RemoteAddr // grabs the IP address of teh visitor

	fmt.Println("the remote ip address is", remoteIP)
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	//render template
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// about is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_IP"] = remoteIP
	fmt.Println("stringMap['remote_ip'] is equal to in about handler", stringMap["remote_ip"])

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// we are sending &TemplateData{} cause we can't send nil
