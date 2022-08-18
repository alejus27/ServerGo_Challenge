package main

import (
	"encoding/json"
	. "file-server/structs"
	. "file-server/utils"
	"github.com/google/uuid"
	"net"
)

type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyAll(message Message)
}

// Define la estructura del cliente.
type Client struct {
	ID         uuid.UUID
	Address    string
	Connection net.Conn
	Observers  []Observer
}

// Crea un nuevo cliente.
func NewClient(connection net.Conn) *Client {
	return &Client{
		ID:         uuid.New(),
		Address:    connection.RemoteAddr().String(),
		Connection: connection,
	}
}

// Control de clientes.
func (client *Client) handle(disconnected chan *Client) {
	var message Message
	for {
		b := make([]byte, MAX_SIZE)
		bs, err := client.Connection.Read(b)

		if err != nil {
			// Cliente se desconecta del servidor y se emite un mensaje.
			PrintError(err.Error(), "1 subscriber went offline")
			disconnected <- client
			for _, obs := range client.Observers {
				if obs.Identifier() == message.Channel {
					go obs.OnDisconnect(client)
				}
			}
			break
			
		} else {
			err = json.Unmarshal(b[:bs], &message)
			if err != nil {
				PrintError(err.Error())
				continue
			}
			go client.NotifyAll(message)
		}
	}
}

// Agrega un observador al cliente.
func (client *Client) Register(observer Observer) {
	client.Observers = append(client.Observers, observer)
}

// Quita un observador al cliente.
func (client *Client) Unregister(observer Observer) {
	for i, obs := range client.Observers {
		if obs.Identifier() == observer.Identifier() {
			client.Observers = append(client.Observers[:i], client.Observers[i+1:]...)
		}
	}
}


// Notifica a todos los observadores de un mensaje.
func (client *Client) NotifyAll(message Message) {
	for _, obs := range client.Observers {
		if obs.Identifier() == message.Channel {
			obs.OnMessage(client, message)
		}
	}
}
