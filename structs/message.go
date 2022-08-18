package structs

import (
	"encoding/json"
	"os"
)

type Message struct {
	Action  string          
	Channel string          
	Message json.RawMessage 
}

type FileMessage struct {
	Name    string
	Size    int64
	Content []byte
	Mode    os.FileMode
}
