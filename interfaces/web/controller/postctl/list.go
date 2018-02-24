package postctl

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// https://stackoverflow.com/questions/41136000/creating-load-more-button-in-golang-with-templates
func (interactor *PostController) GetPostsList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	urlIdStr := vars["id"]

	// Parse ID from URL
	userID, err := strconv.ParseInt(urlIdStr, 10, 64)
	if err != nil {
		interactor.ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	urlLimitStr := r.URL.Query().Get("limit")
	if urlLimitStr == "" {
		urlLimitStr = "1"
	}

	urlLimitInt, err := strconv.ParseInt(urlLimitStr, 10, 64)
	if err != nil {
		log.Println(err)
	}

	urlOffsetStr := r.URL.Query().Get("offset")
	if urlOffsetStr == "" {
		urlOffsetStr = "0"
	}
	urlOffsetInt, err := strconv.ParseInt(urlOffsetStr, 10, 64)
	if err != nil {
		log.Println(err)
	}

	posts, err := interactor.PostUsecases.FindPostsByUserID(userID, urlLimitInt, urlOffsetInt)
	if err != nil {
		log.Println(err)
		interactor.Cookie.AddFlashDanger(w, r, "Error while loading posts")
	}

	m := map[string]interface{}{
		"Posts": posts,
		//"LimitAndOffset": 1,
	}

	//if len(posts) == 0 {
	//	w.WriteHeader(http.StatusNoContent)
	//	return
	//}
	interactor.ProcessTemplate(w, r, "posts_list", m)
}
