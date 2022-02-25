package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/jinyanomura/ezres-web/pkg/config"
	"github.com/jinyanomura/ezres-web/pkg/driver"
	"github.com/jinyanomura/ezres-web/pkg/helpers"
	"github.com/jinyanomura/ezres-web/pkg/models"
	"github.com/jinyanomura/ezres-web/pkg/render"
	"github.com/jinyanomura/ezres-web/pkg/repository"
	"github.com/jinyanomura/ezres-web/pkg/repository/dbrepo"
)

type Repository struct {
	App *config.AppConfig
	DB repository.DatabaseRepo
}

var Repo *Repository

func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB: dbrepo.NewPostgresDBRepo(a, db.SQL),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

func (m *Repository) Document(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "document.page.html", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}

func (m *Repository) Example(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	date := time.Now()

	layout := "2006-01-02"

	stringMap["today"] = date.Format(layout)
	stringMap["maxDate"] = date.AddDate(0, 0, 7).Format(layout)

	if r.URL.Query().Get("y") != "" {
		year, _ := strconv.Atoi(r.URL.Query().Get("y"))
		month, _ := strconv.Atoi(r.URL.Query().Get("m"))
		day, _ := strconv.Atoi(r.URL.Query().Get("d"))
		date = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	}

	stringMap["date"] = date.Format(layout)

	tables, err := m.DB.GetAllTables()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	data := make(map[string]interface{})
	for i, t := range tables {
		restrictions, err := m.DB.GetRestrictionsByDay(t.ID, date)
		if err != nil {
			helpers.ServerError(w, err)
			return
		}
		tables[i].Restrictions = restrictions
	}

	data["tables"] = tables

	render.RenderTemplate(w, r, "example.page.html", &models.TemplateData{
		Data: data,
		StringMap: stringMap,
	})
}