package main

import (
	"encoding/json"
	. "file-server/structs"
	. "file-server/utils"
	"io/ioutil"
	"net"
	"os"
)

// Define la estructura del Emitter para la conexión
type Emitter struct {
	Connection    net.Conn
	subscriptions []string
}

// Devuelve un nuevo Emitter para la conexión.
func NewEmitter(connection net.Conn) *Emitter {
	return &Emitter{
		Connection: connection,
	}
}

// Emite nueva suscripción a canal especifico
func (emitter *Emitter) Subscribe(channel string) {
	if emitter.isSubscribed(channel) {
		PrintWarning("You are subscribed to", channel, "channel")
		return
	}

	message := Message{Action: SUBSCRIBE, Channel: channel}

	emitter.emit(message)
}

// Emite una nueva desuscripción a canal especifico
func (emitter *Emitter) Unsubscribe(channel string) {
	if !emitter.isSubscribed(channel) {
		PrintWarning("You are not subscribed to", channel, "channel")
		return
	}
	message := Message{Action: UNSUBSCRIBE, Channel: channel}
	emitter.emit(message)
}

// Emite el envio de un archivo a canal especifico
func (emitter *Emitter) SendFile(channel string, filePath string) {
	// Verifica si el cliente está suscrito
	if !emitter.isSubscribed(channel) {
		PrintWarning("You are not subscribed to", channel, "channel")
		return
	}

	file, err := os.Open(filePath) // Función integrada para leer un archivo en una ruta
	// Devuelve un mensaje de error si no se pudo leer el archivo
	if err != nil {
		PrintError(err.Error())
		return
	}
	defer file.Close() // Cierre

	infoFile, _ := file.Stat()          // Estructura del archivo
	fileByte, _ := ioutil.ReadAll(file) // Lectura información del archivo
	fileMessage := FileMessage{Name: infoFile.Name(), Size: infoFile.Size(), Content: fileByte, Mode: infoFile.Mode()}
	fileMessageByte, _ := json.Marshal(fileMessage) // Codificación a json del mensaje contenedor de las caracteristicas del archivo

	message := Message{Action: SEND, Channel: channel, Message: fileMessageByte}
	emitter.emit(message) // Emite un nuevo mensaje
}

// Emite un mensaje a la conexión.
func (emitter *Emitter) emit(message Message) {
	data, _ := json.Marshal(message)
	// Si el archivo pesa más de 5mb no hace el envio
	if len(data) > MAX_SIZE {
		PrintWarning("You can not upload more than 5 MB")
		return
	}
	// Escribe la información en la conexión
	emitter.Connection.Write(data)
}

// Devuelve verdadero si el cliente está suscrito al canal especificado
func (emitter *Emitter) isSubscribed(channel string) bool {
	isSubscribed := false

	for _, subscription := range emitter.subscriptions {
		if subscription == channel {
			isSubscribed = true
			break
		}
	}
	return isSubscribed
}

// Se llama a OnEntry cuando se recibe una entrada, luego se devuelve una nuevo Emitter
func (emitter *Emitter) OnEntry(options []string) {
	switch options[0] {
	case SUBSCRIBE:
		emitter.Subscribe(options[1])
	case UNSUBSCRIBE:
		emitter.Unsubscribe(options[1])
	case SEND:
		emitter.SendFile(options[1], options[2])
	case EXIT:
		break
	}

}

//  Devuelve un identificador de cliente.
func (emitter *Emitter) Identifier() string {
	return "client"
}

func (emitter *Emitter) subscriptionListener(responses map[string]chan string) {
	go func() {
		for channel := range responses[SUBSCRIBE] {
			if !emitter.isSubscribed(channel) {
				// Agrega a las suscripciones del Emitter un nuevo canal especificado
				emitter.subscriptions = append(emitter.subscriptions, channel)
			}
		}
	}()

	go func() {
		for response := range responses[UNSUBSCRIBE] {
			// Elimina de las suscripciones del Emitter un canal especificado
			position := -1
			for i, subscription := range emitter.subscriptions {
				if subscription == response {
					position = i
					break
				}
			}
			if position != -1 {
				emitter.subscriptions = append(emitter.subscriptions[:position], emitter.subscriptions[position+1:]...)
			}
		}
	}()
}
