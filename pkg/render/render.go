package render

import (
	"bookings/pkg/config"
	"bookings/pkg/models"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(response_writer http.ResponseWriter, template_string string, app_config *config.AppConfig, template_data *models.TemplateData) {
	if !app_config.UseCache {
		app_config.Template_Cache = CreateTemplateCache()
	}

	cached_template, exists := app_config.Template_Cache[template_string]
	if !exists {
		log.Fatal(fmt.Sprintf("%s template not found.", template_string))
	}
	buffer := new(bytes.Buffer)
	err := cached_template.Execute(buffer, template_data)

	if err != nil {
		log.Fatal(err)
	}

	buffer.WriteTo(response_writer)

}
func CreateTemplateCache() map[string]*template.Template {
	template_cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		log.Fatal(err)
	}

	for _, page := range pages {

		if err != nil {
			log.Fatal(err)
		}
		page_name := filepath.Base(page)
		template_set, err := template.New(page_name).ParseFiles(page)
		if err != nil {
			log.Fatal(err)
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			log.Fatal(err)
		}
		if len(layouts) > 0 {
			template_set, err = template_set.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				log.Fatal(err)
			}
			template_cache[page_name] = template_set
		}

	}
	return template_cache
}
