package structs

import (
	"encoding/json"
	"os"
)

// Define la estructura del mensaje
type Message struct {
	Action  string          `json:"action"`
	Channel string          `json:"channel"`
	Message json.RawMessage `json:"message"`
}

// Define la estructura del mensaje que contiene el archivo
type FileMessage struct {
	Name    string
	Size    int64
	Content []byte
	Mode    os.FileMode
}
