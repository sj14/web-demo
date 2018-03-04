package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sj14/web-demo/infrastructure/repositories/blobs/filesystem"
	"github.com/sj14/web-demo/infrastructure/repositories/database/postgres"
	"github.com/sj14/web-demo/interfaces/web/controller"
	"github.com/sj14/web-demo/interfaces/web/controller/graphqlctl"
	"github.com/sj14/web-demo/interfaces/web/controller/mainctl"
	"github.com/sj14/web-demo/interfaces/web/controller/postctl"
	"github.com/sj14/web-demo/interfaces/web/controller/profilectl"
	"github.com/sj14/web-demo/interfaces/web/controller/userctl"
	"github.com/sj14/web-demo/interfaces/web/sessions"
	"github.com/sj14/web-demo/usecases"
)

type config struct {
	projectName  string
	inProduction bool
	port         string
	dbURL        string
	cookiePass   string
	sys          string
	dataDir      string
}

var cfg = &config{}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg.projectName = "sj-web-demo"
	cfg.port = getenv("PORT", "8080")
	cfg.dbURL = getenv("DATABASE_URL", "user=postgres password=example dbname=demo sslmode=disable")
	cfg.sys = getenv("SYS", "DEV")
	cfg.cookiePass = getenv("COOKIE_PASS", "CHANGEMECHANGEMECHANGEME")
	if cfg.sys == "PROD" {
		cfg.inProduction = true
	}
	cfg.dataDir = getenv("DATA_DIR", "data") // e.g. uploaded files
	log.Println("CONFIG:", cfg)
}

func main() {
	postgresRepo := postgres.NewPostgresStore(cfg.dbURL)
	defer postgresRepo.CloseConn()

	fsRepo := filesystem.NewFilesystemStore(cfg.dataDir)

	userUsecases := usecases.NewUserUsecases(postgresRepo)
	postUsecases := usecases.NewPostUsecases(postgresRepo)
	imageUsecases := usecases.NewImageUsecases(fsRepo)

	cookieStore, err := sessions.NewCookie(cfg.projectName, []byte(cfg.cookiePass))
	if err != nil {
		log.Fatal("Not able to create CookieStore: ", err)
	}

	mainCtl := mainctl.NewMainController(false, cookieStore, userUsecases, postUsecases, imageUsecases)
	profileCtl := profilectl.NewProfileController(mainCtl)
	userCtl := userctl.NewUserController(mainCtl)
	postCtl := postctl.NewPostController(mainCtl)
	graphqlCtl := graphqlctl.NewGraphQLController(mainCtl)

	router := mux.NewRouter()
	routerInteractor := controller.NewRouterInteractor(
		mainCtl,
		profileCtl,
		userCtl,
		postCtl,
		graphqlCtl,
		[]byte("asd"),
		cfg.projectName,
		cfg.inProduction,
	)

	routerInteractor.InitializeRoutes(router)
	log.Println("listening on port " + cfg.port)
	log.Fatal(http.ListenAndServe(":"+cfg.port, nil))
}

// https://stackoverflow.com/a/40326580/7125878
func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
