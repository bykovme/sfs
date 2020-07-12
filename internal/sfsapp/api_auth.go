package sfsapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bykovme/sfs/messages"
	"github.com/bykovme/sfs/models"
)

func (app *App) handleAuth(w http.ResponseWriter, r *http.Request) {
	var authReply models.AuthJSONReply
	if r.Method != "POST" {
		messages.PrepareErrorReply(w, "wrong request type (must be POST)", http.StatusBadRequest)
		return
	}
	var auth models.AuthJSONPost
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &auth)
	if err != nil {
		messages.PrepareErrorReply(w, err.Error(), http.StatusBadRequest)
		return
	}

	authReply.Token, err = app.createToken(auth.LoginEntity)

	if err != nil {
		messages.PrepareErrorReply(w, err.Error(), http.StatusUnauthorized)
		return
	}
	authReply.Status = messages.CSuccess
	messages.PrepareSuccessReply(w, authReply)
}
