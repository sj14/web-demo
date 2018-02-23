package mainctl

import (
	"net/http"
)

func (interactor *MainController) GetPanic(w http.ResponseWriter, r *http.Request) {
	panic("Panic Test")

	m := struct {
		Title string
	}{
		"Panic",
	}

	interactor.ProcessTemplate(w, r, "panic", m)
}
