package handlers

import (
	"net/http"

	"github.com/jinyanomura/ezres/pkg/config"
	"github.com/jinyanomura/ezres/pkg/models"
	"github.com/jinyanomura/ezres/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}