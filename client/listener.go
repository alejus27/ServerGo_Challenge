package main

import (
	"encoding/json"
	. "file-server/structs"
	. "file-server/utils"
	"net"
	"os"
)

// Define la estructura del Listener para la conexión
type Listener struct {
	Connection net.Conn
	Responses  map[string]chan string
	active     bool
}

// Crea un nuevo Listener para la conexión dada.
func NewListener(connection net.Conn) *Listener {
	responses := make(map[string]chan string)
	responses[SUBSCRIBE] = make(chan string)
	responses[UNSUBSCRIBE] = make(chan string)
	return &Listener{
		Connection: connection,
		Responses:  responses,
		active:     true,
	}
}


//  El Listener escucha los mensajes entrantes
func (listener *Listener) Listen() {
	var response Message
	for listener.active {
		b := make([]byte, MAX_SIZE) // Define tamaño y tipo de datos
		bs, err := listener.Connection.Read(b) // Lectura de información desde la conexión 

		if err != nil { // Si el listener no recibe más mensajes se desconecta de la conexión
			PrintError(err.Error(), "Disconnected")
			listener.Stop()
			break
		} else {
			err = json.Unmarshal(b[:bs], &response)
			if err != nil {
				PrintError(err.Error())
				continue
			}

			switch response.Action {
			case SUBSCRIBE:
				listener.Subscribe(response)
			case UNSUBSCRIBE:
				listener.Unsubscribe(response)
			case SEND:
				listener.Send(response)
			}
		}
	}
}

// Detiene la conexión del Listener.
func (listener *Listener) Stop() {
	listener.active = false
}

// Suscripción.
func (listener *Listener) Subscribe(response Message) {
	PrintSuccess(string(response.Message))
	listener.Responses[SUBSCRIBE] <- response.Channel
}

// Cancelar la suscripción.
func (listener *Listener) Unsubscribe(response Message) {
	PrintSuccess(string(response.Message))
	listener.Responses[UNSUBSCRIBE] <- response.Channel
}

// Envía un mensaje al servidor
func (listener *Listener) Send(response Message) {
	var fileMessage FileMessage
	err := json.Unmarshal(response.Message, &fileMessage) //Conversión a json
	if err != nil {
		//PrintError(err.Error())
		PrintSuccess(string(response.Message))
		return
	}
	err = os.WriteFile("../storage/"+fileMessage.Name, fileMessage.Content, fileMessage.Mode) // Función integrada para guardar un archivo en una ruta
	PrintSuccess("New file saved in Storage")

	// Devuelve un mensaje de error si no se pudo guardar el archivo
	if err != nil {
		PrintError("Error saving file", err.Error())
	}

}
