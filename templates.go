package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageView struct {
	PageTitle string
	User      User
	Data      interface{}
	Refresh   bool
}

func renderTemplate(w http.ResponseWriter, t string, v *PageView) {
	var templates = template.Must(
		template.ParseFiles(
			"web/templates/head.html",
			"web/templates/footer.html",
			"web/templates/home.html",
			"web/templates/issues_view.html",
			"web/templates/issues.html",
			"web/templates/charts.html",
			"web/templates/project_add.html",
			"web/templates/resources.html",
			"web/templates/signin.html",
			"web/templates/signup.html",
			"web/templates/upgrade.html",
			"web/templates/manage.html",
			"web/templates/canceled.html",
		))

	err := templates.ExecuteTemplate(w, t+".html", v)
	if err != nil {
		log.Println(err)
	}
}
