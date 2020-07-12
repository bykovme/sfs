package sfsapp

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bykovme/sfs/messages"
	"github.com/bykovme/sfs/models"
	"github.com/gabriel-vasile/mimetype"
)

func (app *App) handleUpload(w http.ResponseWriter, r *http.Request) {
	var fUploaded models.FileUploaded
	var err error
	fUploaded.FileInfo, err = app.uploadTrack(r)
	if err != nil {
		messages.PrepareErrorReply(w, messages.MSGAPI10004UploadFailed+": "+err.Error(), http.StatusInternalServerError)
		return
	}

	fUploaded.Status = messages.CSuccess
	messages.PrepareSuccessReply(w, fUploaded)
}

func (app *App) uploadTrack(r *http.Request) (fInfo models.FileInfo, err error) {
	fullPath := app.LoadedConfig.FilesPath

	err = r.ParseMultipartForm(32 << 20) // 32MB is the default used
	if err != nil {
		return fInfo, err
	}

	fhs := r.MultipartForm.File["file2transfer"]

	if len(fhs) != 1 {
		return fInfo, errors.New("wrong length of file form")
	}
	fh := fhs[0]

	fInfo.Filename = fh.Filename
	fInfo.UniqueId = GenerateId()
	fullPath = fullPath + "/" + fInfo.UniqueId
	fullFile := fullPath + "/" + CDefaultFilename

	reader, errReader := fh.Open()
	if errReader != nil {
		return fInfo, errReader
	}
	fInfo.Size = fh.Size

	mime, err := mimetype.DetectReader(reader)
	if err != nil {
		return fInfo, err
	}
	fInfo.MimeType = mime.String()
	fInfo.FileExt = mime.Extension()

	_, err = ioutil.ReadDir(fullPath)
	if err != nil {
		err := os.MkdirAll(fullPath, 0755)
		if err != nil {
			return fInfo, err
		}
		_, err = ioutil.ReadDir(fullPath)
		if err != nil {
			return fInfo, err
		}
	}

	writer, errWriter := os.Create(fullFile)
	if errWriter != nil {
		return fInfo, errWriter
	}

	_, err = io.Copy(writer, reader)
	if err != nil {
		return fInfo, err
	}

	hash := sha256.New()
	if _, err := io.Copy(hash, reader); err != nil {
		return fInfo, err
	}
	fInfo.Checksum = hex.EncodeToString(hash.Sum(nil))

	err = writer.Close()
	if err != nil {
		return fInfo, err
	}

	err = reader.Close()
	if err != nil {
		return fInfo, err
	}
	fInfoBytes, err := json.Marshal(fInfo)
	if err != nil {
		return fInfo, err
	}

	err = ioutil.WriteFile(fullPath+"/"+CInfoFilename, fInfoBytes, 0744)
	if err != nil {
		return fInfo, err
	}

	return fInfo, nil
}
