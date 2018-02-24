package profilectl

import (
	"net/http"
	"strings"
	"time"

	"log"

	"gitlab.com/sj14/web-clean/src/usecases"
)

func (interactor *ProfileController) ShowLogin(w http.ResponseWriter, r *http.Request) {
	// Check if user is already logged in and redirect to his/her profile
	if ok := interactor.Cookie.IsLoggedIn(w, r); ok {
		interactor.Cookie.AddFlashInfo(w, r, "You are already logged in")
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	}

	m := map[string]interface{}{
		"Title": "Login",
		"Email": r.URL.Query().Get("email"),
	}
	interactor.ProcessTemplate(w, r, "user_login", m)
}

func (interactor *ProfileController) PostLogin(w http.ResponseWriter, r *http.Request) {
	type LoginForm struct {
		Email         string
		PasswordPlain string
	}

	details := LoginForm{
		Email:         r.FormValue("email"),
		PasswordPlain: r.FormValue("password"),
	}
	// Convert email to lowercase characters
	details.Email = strings.ToLower(details.Email)

	userId, err := interactor.UserUsecases.FindUserIdByEmail(details.Email)
	if err != nil {
		log.Println("Not able to find an userid to the given email address", err)
		// Throttle against Brute-Force Attacks
		time.Sleep(2 * time.Second)
		// interactor.ErrorHandler(w, r, http.StatusUnauthorized)
		interactor.Cookie.AddFlashDanger(w, r, "Die Mailaddresse oder das Kennwort ist falsch")
		http.Redirect(w, r, "/login?email="+details.Email, http.StatusSeeOther)
		return
	}

	user, err := interactor.UserUsecases.FindUserById(userId)
	if err != nil {
		log.Println("Not able to find user with the given user id", userId, err)
		// Throttle against Brute-Force Attacks
		time.Sleep(2 * time.Second)
		interactor.ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	log.Println("Trying to login with user: %v", user)

	// Check the password
	match, err := interactor.UserUsecases.LoginUser(user.ID, details.PasswordPlain)
	if match == false {
		// Throttle against Brute-Force Attacks
		time.Sleep(2 * time.Second)
		//interactor.ErrorHandler(w, r, http.StatusUnauthorized)

		if err == usecases.ErrUserDisabled {
			interactor.Cookie.AddFlashDanger(w, r, "Account has been deactivated")
			// TODO: allow user to resend activation link
		} else {
			interactor.Cookie.AddFlashDanger(w, r, "Wrong email or password")
		}

		http.Redirect(w, r, "/login?email="+details.Email, http.StatusSeeOther)
		return
	} else {
		log.Println("Password correct")
		interactor.Cookie.SetLoggedIn(user.ID, w, r)
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	}
}
