package sfsapp

import (
	"net/http"
	"os"

	"github.com/bykovme/sfs/messages"
)

func (app *App) handleCheck(w http.ResponseWriter, r *http.Request) {
	fileId, _ := ShiftPath(r.URL.Path)
	fullPath := app.LoadedConfig.FilesPath + "/" + fileId

	fileData := fullPath + "/" + CDefaultFilename
	fileInfo := fullPath + "/" + CInfoFilename

	if _, err := os.Stat(fileData); os.IsNotExist(err) {
		messages.PrepareErrorReply(w, messages.MSGAPI10005FileNotExists+": "+err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := os.Stat(fileInfo); os.IsNotExist(err) {
		messages.PrepareErrorReply(w, messages.MSGAPI10006FileInfoNotExists+": "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	http.ServeFile(w, r, fileInfo)
}
