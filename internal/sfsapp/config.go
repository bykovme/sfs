package sfsapp

// Config - structure of config file
type Config struct {
	ServerPort     string `json:"server_port"`
	FilesPath  string `json:"files_path"`
	SignatureSalt string `json:"signature_salt"`
}