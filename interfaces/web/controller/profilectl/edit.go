package profilectl

import (
	"net/http"

	"log"

	"github.com/sj14/web-demo/usecases"
)

func (interactor *ProfileController) ShowEditProfile(w http.ResponseWriter, r *http.Request) {
	userId, ok := interactor.Cookie.SessionGetUserId(r)
	if !ok {
		interactor.ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	user, err := interactor.UserUsecases.FindUserById(userId)
	if err != nil {
		log.Println("Cookie UserId not found in DB, error: ", err)
		interactor.Cookie.Logout(w, r)
		interactor.ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	//profilePicture, err := interactor.ImageUsecases.RetrieveUserProfilePicture(userId)
	//if err != nil {
	//	// ignore failure in retrieving user picture
	//}
	//
	//profilePictureBase64 := base64.StdEncoding.EncodeToString(profilePicture)

	m := map[string]interface{}{
		"Title": "Profil",
		"User":  user,
		//"ProfilePicture": profilePictureBase64,
		//"LoggedIn": loggedIn,
	}
	interactor.ProcessTemplate(w, r, "profile_edit", m)
}

// TODO: After changing the email address, do another verification of the address?
func (interactor *ProfileController) PostEditProfile(w http.ResponseWriter, r *http.Request) {
	userID, ok := interactor.Cookie.SessionGetUserId(r)
	if !ok {
		interactor.ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	user, err := interactor.UserUsecases.FindUserById(userID)
	if err != nil {
		// Something went horrible wrong. The userId from the session (Cookie) was not found in the DB.
		// This should really not happen. Get sure and log the usercontroller out.
		interactor.Cookie.Logout(w, r)
		interactor.ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	// HTML Form Values
	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")

	err = interactor.UserUsecases.UpdateUserExceptPassword(user)
	if err == usecases.ErrEmailInUse {
		interactor.Cookie.AddFlashDanger(w, r, "E-Mail already in use")
		http.Redirect(w, r, "/profile/edit", http.StatusSeeOther)
		return
	} else if err != nil {
		interactor.Cookie.AddFlashDanger(w, r, "Error while saving the data")
		http.Redirect(w, r, "/profile/edit", http.StatusSeeOther)
		return
	}

	interactor.Cookie.AddFlashSuccess(w, r, "Sucessfully updates your data")
	http.Redirect(w, r, "/profile/edit", http.StatusSeeOther)
}

func (interactor *ProfileController) PostEditPassword(w http.ResponseWriter, r *http.Request) {
	userID, ok := interactor.Cookie.SessionGetUserId(r)
	if !ok {
		interactor.ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	curPassword := r.FormValue("cur_password")
	newPassword := r.FormValue("new_password")

	err := interactor.UserUsecases.UpdateUserPasswordOnly(userID, curPassword, newPassword)
	if err != nil {
		if err == usecases.ErrPasswordNotMatch {
			interactor.Cookie.AddFlashDanger(w, r, "Your current password was not correct")
			http.Redirect(w, r, "/profile/edit", http.StatusSeeOther)
			return
		}
		interactor.Cookie.AddFlashDanger(w, r, "Internal error while changing the password")
		http.Redirect(w, r, "/profile/edit", http.StatusSeeOther)
		return
	}
	interactor.Cookie.AddFlashSuccess(w, r, "Your password has been changed")
	http.Redirect(w, r, "/profile/edit", http.StatusSeeOther)
	return
}
