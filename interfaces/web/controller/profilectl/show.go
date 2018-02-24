package profilectl

import (
	"log"
	"net/http"
)

func (interactor *ProfileController) ShowProfile(w http.ResponseWriter, r *http.Request) {

	userId, ok := interactor.Cookie.SessionGetUserId(r)
	if !ok {
		interactor.ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	user, err := interactor.UserUsecases.FindUserById(userId)
	if err != nil {
		// Something went horrible wrong. The userId from the session (Cookie) was not found in the DB.
		// This should really not happen. Get sure and log the usercontroller out.
		log.Println("Cookie UserId not found in DB, error: ", err)
		interactor.Cookie.Logout(w, r)
		interactor.ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	//profilePicture, err := interactor.ImageUsecases.RetrieveUserProfilePicture(userId)
	if err != nil {
		log.Println("No Profile Picture found")
		// TODO: set sample profile picture
	}

	//profilePictureBase64 := base64.StdEncoding.EncodeToString(profilePicture)

	m := map[string]interface{}{
		"Title": "Profil",
		"User":  user,
		//"ProfilePicture": profilePictureBase64,
		//"LoggedIn": loggedIn,
	}
	interactor.ProcessTemplate(w, r, "user_profile", m)
}
