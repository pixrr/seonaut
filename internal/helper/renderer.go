package helper

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/stjudewashere/seonaut/internal/user"
)

type PageView struct {
	PageTitle string
	User      user.User
	Data      interface{}
	Refresh   bool
}

type Renderer struct {
	translationMap map[string]interface{}
}

func NewRenderer(m map[string]interface{}) *Renderer {
	return &Renderer{
		translationMap: m,
	}
}

func (r *Renderer) RenderTemplate(w http.ResponseWriter, t string, v *PageView) {
	var templates = template.Must(
		template.New("").Funcs(template.FuncMap{
			"trans": r.trans,
		}).ParseFiles(
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
		))

	err := templates.ExecuteTemplate(w, t+".html", v)
	if err != nil {
		log.Printf("RenderTemplate: %v\n", err)
	}
}

func (r *Renderer) trans(s string) string {
	t, ok := r.translationMap[s]
	if !ok {
		log.Printf("trans: %s translation not found\n", s)
		return s
	}

	return fmt.Sprintf("%v", t)
}
