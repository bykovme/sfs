package models

// SuccessResponse - header for success response
type SuccessResponse struct {
	Status  string `json:"status"`
	MsgNo   string `json:"msgno"`
	MsgText string `json:"msg"`
}

// FileInfo -
type FileInfo struct {
	UniqueId string `json:"id"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
	Checksum string `json:"sha256_checksum"`
	Created  string `json:"created"`
	MimeType string `json:"mimetype"`
	FileExt  string `json:"ext"`
	Preview  bool   `json:"preview"`
}

type FileUploaded struct {
	SuccessResponse
	FileInfo
}

// ErrorResponse - error response structure
type ErrorResponse struct {
	Status  string `json:"status"`
	MsgNo   string `json:"msgno"`
	MsgText string `json:"msg"`
}

// AuthJSONPost - login post structure
type AuthJSONPost struct {
	LoginEntity int64 `json:"user_id"`
}

// AuthJSONReply - reply to successful login
type AuthJSONReply struct {
	SuccessResponse
	Token string `json:"token"`
}
