package handlers

import (
	"bookings/pkg/config"
	"bookings/pkg/models"
	"bookings/pkg/render"
	"net/http"
)

var Repo Repository

type Repository struct {
	App_Config    *config.AppConfig
	Template_Data *models.TemplateData
}

func NewRepo(app_config *config.AppConfig) {
	Repo.App_Config = app_config
}
func (repo *Repository) About(response_writer http.ResponseWriter, request *http.Request) {
	test_data := map[string]string{}
	test_data["about"] = "This is the about page"
	render.RenderTemplate(response_writer, "about.page.tmpl", Repo.App_Config, &models.TemplateData{
		StringMap: test_data,
	})
}

func (repo *Repository) Home(response_writer http.ResponseWriter, request *http.Request) {
	test_data := map[string]string{}
	test_data["home"] = "This is the home page"
	render.RenderTemplate(response_writer, "home.page.tmpl", Repo.App_Config, &models.TemplateData{
		StringMap: test_data,
	})
}
