package userctl

import (
	"net/http"

	"strconv"

	"log"

	"github.com/gorilla/mux"
)

func (interactor *UserController) ShowUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	urlIdStr := vars["id"]

	// Parse ID from URL
	userID, err := strconv.ParseInt(urlIdStr, 10, 64)
	if err != nil {
		interactor.ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	user, err := interactor.UserUsecases.FindUserById(userID)
	if err != nil {
		log.Println(err)
		interactor.ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	m := map[string]interface{}{
		"User": user,
		//"Limit":    1,
		"UrlQuery": r.URL.RawQuery,
	}
	interactor.ProcessTemplate(w, r, "user_show", m)
}
