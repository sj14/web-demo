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

const projectName = "sj-web-demo"

var (
	inProduction bool
	port         string
	dbURL        string
	cookiePass   string
	sys          string
)

func init() {
	port = getenv("PORT", "8080")
	dbURL = getenv("DATABASE_URL", "user=postgres password=example dbname=demo sslmode=disable")
	sys = getenv("SYS", "DEV")
	cookiePass = getenv("COOKIE_PASS", "CHANGEMECHANGEMECHANGEME")
	if sys == "PROD" {
		inProduction = true
	}
}

func main() {
	postgresRepo := postgres.NewPostgresStore(dbURL)
	defer postgresRepo.CloseConn()

	fsRepo := filesystem.NewFilesystemStore()

	userUsecases := usecases.NewUserUsecases(postgresRepo)
	postUsecases := usecases.NewPostUsecases(postgresRepo)
	imageUsecases := usecases.NewImageUsecases(fsRepo)

	cookieStore, err := sessions.NewCookie(projectName, []byte(cookiePass))
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
		projectName,
		inProduction,
	)

	routerInteractor.InitializeRoutes(router)
	log.Println("listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// https://stackoverflow.com/a/40326580/7125878
func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
