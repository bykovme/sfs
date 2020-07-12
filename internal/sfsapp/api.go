package sfsapp

import (
	"net/http"

	"github.com/bykovme/sfs/messages"
	"github.com/bykovme/sfs/models"
)

func (app *App) ServeAPI(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)

	switch head {
	case "ok":
		messages.PrepareSuccessReply(w, models.SuccessResponse{Status: messages.CSuccess, MsgNo: "0", MsgText: ""})
		return
	}

	switch head {
	case "v1":
		app.handleVersions(w, r)
	default:
		messages.PrepareErrorReply(w, messages.MSGAPI10002APINotFound, http.StatusMethodNotAllowed)
	}
}

func (app *App) handleVersions(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)

	switch head {
	case "ok":
		messages.PrepareSuccessReply(w, models.SuccessResponse{Status: messages.CSuccess, MsgNo: "0"})
		return
	case "auth":
		app.handleAuth(w, r)
		return
	}

	err := app.CheckHeaderToken(r)
	if err != nil {
		messages.PrepareErrorReply(w, messages.MSGAPI10001NotAuthorized+":"+err.Error(), http.StatusForbidden)
		return
	}
	switch head {
	case "upload":
		app.handleUpload(w, r)
	case "download":
		app.handleDownload(w, r)
	case "check":
		app.handleCheck(w, r)

	default:
		messages.PrepareErrorReply(w, messages.MSGAPI10002APINotFound, http.StatusMethodNotAllowed)
	}
}
