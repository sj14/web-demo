package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sj14/web-demo/infrastructure/repositories/database/postgres"
	"github.com/sj14/web-demo/interfaces/web/controller"
	"github.com/sj14/web-demo/interfaces/web/controller/mainctl"
	"github.com/sj14/web-demo/interfaces/web/controller/profilectl"
	"github.com/sj14/web-demo/interfaces/web/controller/userctl"
	"github.com/sj14/web-demo/interfaces/web/sessions"
	"github.com/sj14/web-demo/usecases"
)

func main() {
	fmt.Println("Hello World")

	postgresRepo := postgres.NewPostgresStore()
	defer postgresRepo.CloseConn()

	userUsease := usecases.NewUserUsecases(postgresRepo)

	cookieStore, err := sessions.NewCookie("sj-web-demo", []byte("TODOTODOTODOTODO"))
	if err != nil {
		log.Fatal("Not able to create CookieStore: ", err)
	}

	mainCtl := mainctl.NewMainController(false, cookieStore, userUsease)
	profileCtl := profilectl.NewProfileController(mainCtl)
	userCtl := userctl.NewUserController(mainCtl)

	router := mux.NewRouter()
	routerInteractor := controller.NewRouterInteractor(
		mainCtl,
		profileCtl,
		userCtl,
		[]byte("asd"),
		false,
	)

	routerInteractor.InitializeRoutes(router)
	log.Println("listening on port " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
