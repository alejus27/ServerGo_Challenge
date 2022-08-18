package main

import (
	"encoding/json"
	. "file-server/structs"
	. "file-server/utils"
)

type Broadcaster struct {
}

var singleInstance *Broadcaster

// El Broadcaster permite que el emisor envíe mensajes y los suscriptores escuchen estos
func BroadcastInstance() *Broadcaster {
	if singleInstance == nil {
		singleInstance = &Broadcaster{}
	}
	return singleInstance
}

// Emisión de mensaje a todos los suscriptores
func (broadcaster *Broadcaster) Broadcast(subscribers []*Client, message Message) {
	for _, subscriber := range subscribers {
		broadcaster.emit(subscriber, message)
	}
}


func (broadcaster *Broadcaster) emit(subscriber *Client, message Message) {
	data, err := json.Marshal(message)
	if err != nil {
		PrintError(err.Error())
	}
	/*if len(data) > MAX_SIZE {
		PrintWarning("You can not upload more than 5 MB")
		return
	}*/
	subscriber.Connection.Write(data)
}
