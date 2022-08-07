package main

import (
	. "file-server/cli"
)

// Clase principal donde se inicializa el servidor
func main() {
	server := NewServer()
	cli := NewCli("server")
	cli.Register(server)
	go cli.Start()

	for active := range cli.Active {
		if !active {
			break
		}
	}

}
