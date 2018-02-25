package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sj14/web-demo/infrastructure/repositories/blobs/filesystem"
	"github.com/sj14/web-demo/infrastructure/repositories/database/postgres"
	"github.com/sj14/web-demo/interfaces/web/controller"
	"github.com/sj14/web-demo/interfaces/web/controller/mainctl"
	"github.com/sj14/web-demo/interfaces/web/controller/postctl"
	"github.com/sj14/web-demo/interfaces/web/controller/profilectl"
	"github.com/sj14/web-demo/interfaces/web/controller/userctl"
	"github.com/sj14/web-demo/interfaces/web/sessions"
	"github.com/sj14/web-demo/usecases"
)

func main() {
	postgresRepo := postgres.NewPostgresStore()
	defer postgresRepo.CloseConn()

	fsRepo := filesystem.NewFilesystemStore()

	userUsecases := usecases.NewUserUsecases(postgresRepo)
	postUsecases := usecases.NewPostUsecases(postgresRepo)
	imageUsecases := usecases.NewImageUsecases(fsRepo)

	cookieStore, err := sessions.NewCookie("sj-web-demo", []byte("TODOTODOTODOTODO"))
	if err != nil {
		log.Fatal("Not able to create CookieStore: ", err)
	}

	mainCtl := mainctl.NewMainController(false, cookieStore, userUsecases, postUsecases, imageUsecases)
	profileCtl := profilectl.NewProfileController(mainCtl)
	userCtl := userctl.NewUserController(mainCtl)
	postCtl := postctl.NewPostController(mainCtl)

	router := mux.NewRouter()
	routerInteractor := controller.NewRouterInteractor(
		mainCtl,
		profileCtl,
		userCtl,
		postCtl,
		[]byte("asd"),
		false,
	)

	routerInteractor.InitializeRoutes(router)
	log.Println("listening on port " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
