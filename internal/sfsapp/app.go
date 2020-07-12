package sfsapp

import (
	"log"
	"net/http"
	"os"

	"github.com/bykovme/sfs/messages"
)

// App - main application
type App struct {
	LoadedConfig  Config
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)
	switch head {
	case "api":
		app.ServeAPI(w, r)
	default:
		messages.PrepareErrorReply(w, messages.MSGAPI10002APINotFound, http.StatusMethodNotAllowed)
	}
}

// GetPort returns default port either from config or from environment variable
// environment variable has priority
func (app *App) GetPort() string {
	port := app.LoadedConfig.ServerPort

	portEnv := os.Getenv("PORT")

	if portEnv != "" {
		port = portEnv
		log.Println("Using port set in the environment, PORT=" + port)
	}
	return port
}