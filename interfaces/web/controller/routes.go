package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sj14/web-demo/interfaces/web/controller/mainctl"
	"github.com/sj14/web-demo/interfaces/web/controller/userctl"
)

func NewRouterInteractor(
	mainController mainctl.MainController,
	userController userctl.UserController,
	csrfTokenSecret []byte,
	inProductionMode bool,
) RouterInteractor {

	return RouterInteractor{
		mainController,
		userController,
		csrfTokenSecret,
		inProductionMode}
}

type RouterInteractor struct {
	mainController   mainctl.MainController
	userController   userctl.UserController
	csrfTokenSecret  []byte
	inProductionMode bool
}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

func (interactor *RouterInteractor) InitializeRoutes(router *mux.Router) {
	log.Println("Initialize Routes")

	router.StrictSlash(true)

	/////////////////////////////////
	// Main Handler CSRF Protected //
	/////////////////////////////////
	http.Handle("/", handlers.RecoveryHandler()(router))

	//CSRF := csrf.Protect(
	//	[]byte(interactor.csrfTokenSecret),
	//	csrf.CookieName("sj-web-demo_csrf"),
	//	csrf.Secure(interactor.inProductionMode), // if in Production mode, secure is set to true
	//	csrf.ErrorHandler(http.HandlerFunc(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		userID, ok := interactor.mainController.Cookie.SessionGetUserId(r)
	//		if ok != true {
	//			userID = -1
	//		}
	//		interactor.mainController.Cookie.AddFlashDanger(w, r, "CSRF authentification failed")
	//		interactor.mainController.ErrorHandler(w, r, http.StatusForbidden)
	//	})),
	//	))
	//
	//http.Handle("/", CSRF(router))

	// Page Not Found
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		interactor.mainController.ErrorHandler(w, r, http.StatusNotFound)
	})

	////////////////////////
	// Serve static files //
	////////////////////////
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("interfaces/web/files/static/"))))

	serveSingle("/robots.txt", "interfaces/web/files/robots.txt")
	serveSingle("/favicon.ico", "interfaces/web/files/favicon.ico")

	router.HandleFunc("/", interactor.mainController.GetHome).Methods(http.MethodGet)
	router.HandleFunc("/panic", interactor.mainController.GetPanic).Methods(http.MethodGet)

}
