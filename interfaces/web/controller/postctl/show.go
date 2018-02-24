package postctl

import (
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

func (interactor *PostController) ShowPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	urlIdStr := vars["id"]

	// Parse ID from URL
	postID, err := strconv.ParseInt(urlIdStr, 10, 64)
	if err != nil {
		interactor.ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	post, err := interactor.PostUsecases.FindPostByID(postID)
	if err != nil {
		interactor.ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	user, err := interactor.UserUsecases.FindUserById(post.UserID)
	if err != nil {
		interactor.ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	m := map[string]interface{}{
		"Post": post,
		"User": user,
	}
	interactor.ProcessTemplate(w, r, "post_show", m)
}
