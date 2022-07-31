
package main

import (
	. "file-server/cli"
)

//var wg sync.WaitGroup

// Starts a client and emits the response.
func main() {

	connection := NewConnection()
	err := connection.Start()
	if err != nil {
		return
	}
	cli := NewCli("client")
	cli.Register(connection.emitter)
	go cli.Start()

	go func() {
		for active := range cli.Active {
			if !active {
				connection.close()
				break
			}
		}
	}()

	connection.HandleResponse()

}
