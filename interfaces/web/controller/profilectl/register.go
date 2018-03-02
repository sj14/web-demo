package profilectl

import (
	"net/http"

	"log"

	"github.com/sj14/web-demo/usecases"
)

func (interactor *ProfileController) ShowRegister(w http.ResponseWriter, r *http.Request) {
	// Check if user is already logged in and redirect to his/her profile
	if ok := interactor.Cookie.IsLoggedIn(w, r); ok {
		interactor.Cookie.AddFlashInfo(w, r, "Sie sind bereits angemeldet")
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	}

	m := map[string]interface{}{
		"Title": "Registrieren",
	}
	interactor.ProcessTemplate(w, r, "register", m)
}

func (interactor *ProfileController) PostRegister(w http.ResponseWriter, r *http.Request) {
	type RegisterForm struct {
		Name          string
		Email         string
		PasswordPlain string
	}

	details := RegisterForm{
		Name:          r.FormValue("name"),
		Email:         r.FormValue("email"),
		PasswordPlain: r.FormValue("password"),
	}

	_, err := interactor.MainController.UserUsecases.CreateUser(details.Name, details.Email, details.PasswordPlain)

	if err != nil {
		if err == usecases.ErrEmailInUse {
			interactor.Cookie.AddFlashDanger(w, r, "Die Mailaddresse wird bereits verwendet")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			log.Println("Failed to create new user", err)
			interactor.Cookie.AddFlashDanger(w, r, "Die Registrierung ist fehlgeschlagen")
			interactor.ErrorHandler(w, r, http.StatusInternalServerError)
		}
	}

	interactor.Cookie.AddFlashSuccess(w, r, "Successfully registered")
	http.Redirect(w, r, "/login", http.StatusSeeOther)

}
