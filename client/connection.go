package main

import (
	. "file-server/utils"
	"net"
)

// Define la estructura de conexión.
type Connection struct {
	Network  string
	Address  string
	Conn     net.Conn
	emitter  *Emitter 
	listener *Listener 
}

// Crea una nueva conexión.
func NewConnection() *Connection {
	// Define los parámetros de nuestro servidor.
	return &Connection{
		Network: "tcp",
		Address: ":20000",
	}
}

// Inicializa y se conecta al servidor.
func (connection *Connection) Start() error {
	c, err := net.Dial("tcp", ":20000")
	// Devuelve un mensaje de error si no se establece la conexión-
	if err != nil {
		PrintError(err.Error())
		return err
	}
	// Conexión exitosa al servidor
	PrintSuccess("----Connected to server----")
	connection.Conn = c // Establece conexión con el protocolo tcp
	connection.emitter = NewEmitter(connection.Conn)
	connection.listener = NewListener(connection.Conn)
	go connection.emitter.subscriptionListener(connection.listener.Responses)
	return nil
}

// Escucha la conexión
func (connection *Connection) HandleResponse() {
	connection.listener.Listen()
}

// Detiene y cierra la conexión 
func (connection *Connection) close() {
	connection.listener.Stop()
	connection.Conn.Close()
}
