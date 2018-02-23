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
	"github.com/gorilla/handlers"
	"os"
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
	log.Println("listening on port " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":" +	os.Getenv("PORT"), handlers.RecoveryHandler()(router)))
}
