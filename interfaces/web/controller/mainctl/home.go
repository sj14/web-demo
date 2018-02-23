package mainctl

import (
	"log"
	"net/http"
)

func (interactor *MainController) GetHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Println("URL Path not found: ", r.URL.Path)
		w.WriteHeader(http.StatusNotFound)

		m := map[string]interface{}{
			"Path": r.RequestURI,
		}
		interactor.ProcessTemplate(w, r, "404_NotFound", m)
		return
	}

	m := map[string]interface{}{
		"Title": "Home",
	}
	interactor.ProcessTemplate(w, r, "index", m)
}
