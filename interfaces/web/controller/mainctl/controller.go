package mainctl

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/sj14/web-demo/interfaces/web/sessions"
	"github.com/sj14/web-demo/interfaces/web/view"
	"github.com/sj14/web-demo/usecases"
)

func NewMainController(
	inProductionMode bool,
	cookieStore sessions.Cookie,
	userUsecases usecases.UserUsecases,
	postUsecases usecases.PostUsecases,
) MainController {

	return MainController{inProductionMode,
		cookieStore,
		userUsecases,
		postUsecases,
	}
}

type MainController struct {
	inProductionMode bool
	Cookie           sessions.Cookie
	UserUsecases     usecases.UserUsecases
	PostUsecases     usecases.PostUsecases
}

type Interface interface{}

// ProcessTemplate the named template
func (interactor *MainController) ProcessTemplate(w http.ResponseWriter, r *http.Request, tmpl string, data Interface) {
	infoFlashesMsgs := interactor.Cookie.PopFlashesInfo(w, r)
	successFlashesMsgs := interactor.Cookie.PopFlashesSuccess(w, r)
	warningFlashesMsgs := interactor.Cookie.PopFlashesWarning(w, r)
	dangerFlashesMsgs := interactor.Cookie.PopFlashesDanger(w, r)

	loggedIn := interactor.Cookie.IsLoggedIn(w, r)

	payload := struct {
		CtlData        Interface
		InfoFlashes    []string
		SuccessFlashes []string
		WarningFlashes []string
		DangerFlashes  []string
		CSRFField      template.HTML
		LoggedIn       bool
	}{
		data,
		infoFlashesMsgs,
		successFlashesMsgs,
		warningFlashesMsgs,
		dangerFlashesMsgs,
		csrf.TemplateField(r),
		loggedIn,
	}

	err := view.Templates.ExecuteTemplate(w, tmpl, payload)
	if err != nil {
		log.Println("Template not found: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (interactor *MainController) ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	switch status {
	case http.StatusNotFound:
		log.Println("URL Path not found: ", r.URL.Path)
		m := struct {
			Title string
			Path  string
		}{
			"Ups",
			r.RequestURI,
		}
		interactor.ProcessTemplate(w, r, "404_NotFound", m)
	case http.StatusBadRequest:
		m := struct {
			Title string
			Path  string
		}{
			"Ups",
			r.RequestURI,
		}
		interactor.ProcessTemplate(w, r, "400_BadRequest", m)
	case http.StatusUnauthorized:
		m := struct {
			Title string
			Path  string
		}{
			"Ups",
			r.RequestURI,
		}
		interactor.ProcessTemplate(w, r, "401_Unauthorized", m)
	case http.StatusForbidden:
		m := struct {
			Title string
			Path  string
		}{
			"Ups",
			r.RequestURI,
		}
		interactor.ProcessTemplate(w, r, "403_Forbidden", m)
	case http.StatusInternalServerError:
		m := struct {
			Title string
			Path  string
		}{
			"Ups",
			r.RequestURI,
		}
		interactor.ProcessTemplate(w, r, "500_InternalServerError", m)
	default:
		log.Println("Error Page not found:", status)
		m := struct {
			Title string
			Path  string
		}{
			"Ups",
			r.RequestURI,
		}
		interactor.ProcessTemplate(w, r, "500_InternalServerError", m)
	}
}
