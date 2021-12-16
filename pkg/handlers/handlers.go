package handlers

import (
	"go-web/pkg/config"
	"go-web/pkg/models"
	"go-web/pkg/render"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

// TemplateData holds data sent to templates


var Repo * Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewRepo
func NewHandlers(r *Repository) {
	Repo = r
}


func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

