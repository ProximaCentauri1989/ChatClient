package client

import (
	"chat_client/config"
	"net"
	"os"
)

//Dials address and performs messaging with server
func ConnectToServer(network, address string) {
	conn, err := net.Dial(network, address)
	defer conn.Close()

	config.FailOnError(err, "connect to server")
	go PrintMessages(conn)
	_, err = SayHello(conn)
	config.FailOnError(err, "say hello to server")

	ReadFromCopyTo(conn, os.Stdin)
}
