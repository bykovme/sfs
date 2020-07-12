package sfsapp

import (
	"net/http"
)

func (app *App) handleDownload(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filename))
	//w.Header().Set("Content-Type", "application/octet-stream")
	//http.ServeFile(w, r, filePath)
	// http.ServeFile(w,r,fileData)
}
