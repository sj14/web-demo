package mainctl

import "net/http"

func (interactor *MainController) GetCSRFTest(w http.ResponseWriter, r *http.Request) {
	m := struct {
		Title string
	}{
		"CSRF Test",
	}
	interactor.ProcessTemplate(w, r, "csrf_test", m)
}

func (interactor *MainController) PostCSRF(w http.ResponseWriter, r *http.Request) {
	interactor.Cookie.AddFlashSuccess(w, r, "POST sent successfully")
	http.Redirect(w, r, "/csrf", http.StatusFound)
}
