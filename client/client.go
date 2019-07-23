package client

import (
	"ChatClient/failer"
	"net"
	"os"
)

//Dials address and performs messaging with server
func ConnectToServer(network, address string) {
	conn, err := net.Dial(network, address)
	defer conn.Close()

	failer.FailOnError(err, "connect to server")
	go PrintMessages(conn)
	_, err = SayOneMessage(conn, "Client says hello!!!\n")
	failer.FailOnError(err, "say hello to server")

	ReadFromCopyTo(conn, os.Stdin)
	_, err = SayOneMessage(conn, "Client says goodbye!!!\n")
}
