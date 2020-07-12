package messages

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"unicode"

	"github.com/bykovme/sfs/models"
)

// PrepareSuccessReply - prepare successful reply
func PrepareSuccessReply(w http.ResponseWriter, preparedStruct interface{}) {

	b, err := json.Marshal(preparedStruct)
	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		_, err := io.WriteString(w, string(b[:]))
		if err != nil {
			log.Println(err)
		}
		return
	}
	log.Println(err)
}

// PrepareErrorReply - preparing json error reply
func PrepareErrorReply(w http.ResponseWriter, msgID string, httpStatus int) {

	b, err := json.Marshal(parseError(msgID))
	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(httpStatus)
		_, err = io.WriteString(w, string(b[:]))
		if err != nil {
			log.Println(err.Error())
		}
		return
	}
	log.Println(err.Error())
}

// CheckErrorFormat - check error format
func CheckErrorFormat(s string) bool {
	if len(s) < 7 {
		return false
	}

	errCodeLength := cErrorCodeLength
	for _, r := range s {
		if errCodeLength == 0 {
			if unicode.IsSpace(r) {
				return true
			}
		}
		if !unicode.IsDigit(r) {
			return false
		}
		errCodeLength--
	}
	return false
}

func parseError(msgID string) (errResp models.ErrorResponse) {
	if msgID == "" {
		return parseError(MSGNumUnknown)
	}

	if CheckErrorFormat(msgID) == false {
		errResp.Status = CError
		errResp.MsgNo = MSGNumUnknown
		errResp.MsgText = msgID
		return errResp
	}

	errResp.Status = CError
	errResp.MsgNo = msgID[:5]
	errResp.MsgText = msgID[6:]
	return errResp
}


