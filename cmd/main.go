package main

import (
	"chat_client/client"
	"chat_client/config"
	"chat_client/server"
)

func main() {
	connection, err := server.StartServer(config.Network, config.Address)

	if err != nil && IsAlreadyBinded(err) {
		client.ConnectToServer(config.Network, config.Address)
	} else {
		config.FailOnError(err, "start server")
	}

	server.HandleConnection(connection)
}
