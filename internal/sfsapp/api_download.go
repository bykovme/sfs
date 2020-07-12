package sfsapp

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/bykovme/sfs/messages"
	"github.com/bykovme/sfs/models"
)

func (app *App) handleDownload(w http.ResponseWriter, r *http.Request) {
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
	file, err := os.Open(fileInfo)
	if err != nil {
		messages.PrepareErrorReply(w, messages.MSGAPI10006FileInfoNotExists+": "+err.Error(), http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(file)
	var fInfo models.FileInfo
	err = decoder.Decode(&fInfo)
	if err != nil {
		messages.PrepareErrorReply(w, messages.MSGAPI10006FileInfoNotExists+": "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = file.Close()
	if err != nil {
		messages.PrepareErrorReply(w, messages.MSGAPI10006FileInfoNotExists+": "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(fInfo.Filename))
	//w.Header().Set("Content-Type", fInfo.MimeType)
	http.ServeFile(w, r, fileData)
}
