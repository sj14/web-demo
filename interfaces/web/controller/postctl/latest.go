package postctl

import (
	"log"
	"net/http"
	"strconv"
)

// https://stackoverflow.com/questions/41136000/creating-load-more-button-in-golang-with-templates
func (interactor *PostController) GetPostsLatest(w http.ResponseWriter, r *http.Request) {
	urlLimitStr := r.URL.Query().Get("limit")
	if urlLimitStr == "" {
		urlLimitStr = "10"
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

	posts, err := interactor.PostUsecases.FindNewestPosts(urlLimitInt, urlOffsetInt)
	if err != nil {
		log.Println(err)
		interactor.Cookie.AddFlashDanger(w, r, "Error while loading posts")
	}

	m := map[string]interface{}{
		"Posts": posts,
		//"LimitAndOffset": 1,
	}

	interactor.ProcessTemplate(w, r, "posts_list", m)
}
