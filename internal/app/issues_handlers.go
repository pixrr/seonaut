package app

import (
	"log"
	"net/http"
	"strconv"

	"github.com/mnlg/lenkrr/internal/issue"
	"github.com/mnlg/lenkrr/internal/project"
	"github.com/mnlg/lenkrr/internal/user"
)

type IssuesGroupView struct {
	ProjectView *project.ProjectView
	MediaChart  Chart
	StatusChart Chart
	IssueCount  *issue.IssueCount
}

type IssuesView struct {
	ProjectView   *project.ProjectView
	Eid           string
	Project       project.Project
	PaginatorView issue.PaginatorView
}

func (app *App) serveIssues(user *user.User, w http.ResponseWriter, r *http.Request) {
	pid, err := strconv.Atoi(r.URL.Query().Get("pid"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)

		return
	}

	pv, err := app.projectService.GetProjectView(pid, user.Id)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	issueCount := app.issueService.GetIssuesCount(pv.Crawl.Id)

	ig := IssuesGroupView{
		ProjectView: pv,
		MediaChart:  NewChart(issueCount.MediaCount),
		StatusChart: NewChart(issueCount.StatusCount),
		IssueCount:  issueCount,
	}

	v := &PageView{
		Data:      ig,
		User:      *user,
		PageTitle: "ISSUES_VIEW",
	}

	app.renderer.renderTemplate(w, "issues", v)
}

func (app *App) serveIssuesView(user *user.User, w http.ResponseWriter, r *http.Request) {
	eid := r.URL.Query().Get("eid")
	if eid == "" {
		log.Println("serveIssuesView: eid parameter missing")
		http.Redirect(w, r, "/", http.StatusSeeOther)

		return
	}

	pid, err := strconv.Atoi(r.URL.Query().Get("pid"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)

		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("p"))
	if err != nil {
		page = 1
	}

	pv, err := app.projectService.GetProjectView(pid, user.Id)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)

		return
	}

	paginatorView, err := app.issueService.GetPaginatedReportsByIssue(pv.Crawl.Id, page, eid)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)

		return
	}

	view := IssuesView{
		ProjectView:   pv,
		Eid:           eid,
		Project:       pv.Project,
		PaginatorView: paginatorView,
	}

	v := &PageView{
		Data:      view,
		User:      *user,
		PageTitle: "ISSUES_DETAIL",
	}

	app.renderer.renderTemplate(w, "issues_view", v)
}
