package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sj14/web-demo/interfaces/web/controller"
	"github.com/sj14/web-demo/interfaces/web/controller/userctl"
	"github.com/sj14/web-demo/interfaces/web/controller/mainctl"
	"github.com/sj14/web-demo/usecases"
	"github.com/sj14/web-demo/infrastructure/repositories/database/postgres"
	"net/http"
	"log"
)

func main() {
	fmt.Println("Hello World")

	postgresRepo := postgres.NewPostgresStore()
	defer postgresRepo.CloseConn()

	userUsease := usecases.NewUserUsecases(postgresRepo)

	mainCtl := mainctl.NewMainController(false, userUsease)

	userCtl := userctl.NewUserController(mainCtl)

	router := mux.NewRouter()
	routerInteractor := controller.NewRouterInteractor(
		mainCtl,
		userCtl,
		[]byte("asd"),
		false,
	)

	routerInteractor.InitializeRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
