package main

import (
	. "file-server/cli"
)

// Entry point for the server.
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
