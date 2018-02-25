package userctl

import (
	"net/http"

	"strconv"

	"log"

	"encoding/base64"

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

	profilePicture, err := interactor.ImageUsecases.RetrieveUserProfilePicture(user.ID)
	if err != nil {
		log.Println(err)
	}

	profilePictureBase64 := base64.StdEncoding.EncodeToString(profilePicture)

	m := map[string]interface{}{
		"User":           user,
		"UrlQuery":       r.URL.RawQuery,
		"ProfilePicture": profilePictureBase64,
	}
	interactor.ProcessTemplate(w, r, "user_show", m)
}
