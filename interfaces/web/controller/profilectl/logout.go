package profilectl

import "net/http"

func (interactor *ProfileController) PostLogout(w http.ResponseWriter, r *http.Request) {
	interactor.Cookie.Logout(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
