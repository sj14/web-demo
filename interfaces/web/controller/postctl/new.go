package postctl

import (
	"net/http"
	"strconv"

	"time"
)

func (interactor *PostController) ShowNewPost(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{
		"Title": "New Post",
	}
	interactor.ProcessTemplate(w, r, "post_new", m)
}

func (interactor *PostController) PostNewPost(w http.ResponseWriter, r *http.Request) {
	userID, ok := interactor.Cookie.SessionGetUserId(r)
	if !ok {
		interactor.ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	text := r.FormValue("text")

	postID, err := interactor.PostUsecases.PublishPost(userID, text, time.Now())
	if err != nil {
		interactor.Cookie.AddFlashDanger(w, r, "Error while creating the post")
		http.Redirect(w, r, "/post/new", http.StatusSeeOther)
		return
	}

	postIdStr := strconv.FormatInt(postID, 10)
	http.Redirect(w, r, "/post/"+postIdStr, http.StatusFound)
}
