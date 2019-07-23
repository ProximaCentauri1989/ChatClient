package main

import (
	"ChatClient/client"
	"ChatClient/config"
	"ChatClient/failer"
	"ChatClient/server"
)

func main() {
	connection, err := server.StartServer(config.Network, config.Address)

	if err != nil && IsAlreadyBinded(err) {
		client.ConnectToServer(config.Network, config.Address)
	} else {
		failer.FailOnError(err, "start server")
	}

	server.HandleConnection(connection)
}
