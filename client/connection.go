package main

import (
	. "file-server/utils"
	"net"
)

// Estructura
type Connection struct {
	Network  string
	Address  string
	Conn     net.Conn //Conexión de red genérica orientada a la transmisión
	emitter  *Emitter 
	listener *Listener 
}

// Parámetros para conexión con servidor.
func NewConnection() *Connection {
	return &Connection{
		Network: "tcp",
		Address: ":20000",
	}
}

// Conexión al servidor.
func (connection *Connection) Start() error {
	c, err := net.Dial("tcp", ":20000")

	if err != nil {
		PrintError(err.Error())
		return err
	}
	// Conexión exitosa al servidor
	PrintSuccess("---> Successful connection to the server <---")

	connection.Conn = c 
	connection.emitter = NewEmitter(connection.Conn)
	connection.listener = NewListener(connection.Conn)
	go connection.emitter.subscriptionListener(connection.listener.Responses)
	
	return nil
}

func (connection *Connection) HandleResponse() {
	connection.listener.Listen()
}

func (connection *Connection) close() {
	connection.listener.Stop()
	connection.Conn.Close()
}
