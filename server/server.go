package server

import (
	"bufio"
	"chat_client/config"
	"fmt"
	"log"
	"net"
	"os"
)

//Create listener on a given address
func StartServer(network, address string) (net.Listener, error) {
	return net.Listen(network, address)
}

//Waits for client and starts connection handler
func HandleConnection(listener net.Listener) {
	for {
		log.Print("Waiting for new connection...\n")
		conn, err := listener.Accept()
		config.FailOnError(err, "accepting connection")
		log.Print("Chat session has been started...\n")
		StartCommunication(conn)
		log.Print("Chat session has been disconnected...\n")
	}
}

/*Performs everything you need to exchange messages with a client
  Notes: We use main channel and auxiliary channels to terminate goroutines: ConnectionWriter and RedirectToChannel.
  When messageChannel is closed its produce a loop exit in ConnectionWriter.
  Then ConnectionWriter terminates if writerCanceler receive its value.
  The RedirectToChannel routine terminates when redirecterCanceler receive its value
*/
func StartCommunication(connection net.Conn) {
	//channels
	writerCanceler := make(chan bool)
	redirecterCanceler := make(chan bool)
	messageChannel := make(chan string)

	//start connection writer
	go ConnectionWriter(connection, messageChannel, writerCanceler)

	//send to channel client's address
	who := connection.RemoteAddr().String()
	messageChannel <- "You " + who

	//start input redirector
	go RedirectToChannel(messageChannel, redirecterCanceler)

	//Start reading messages from client
	input := bufio.NewScanner(connection)
	for input.Scan() {
		fmt.Fprintln(os.Stdout, who+config.Separator+input.Text())
	}

	fmt.Fprintln(os.Stdout, who+" disconnected. Press any key to continue")

	//Close main channel, close connection, stop goroutines by sending true to them
	close(messageChannel)
	connection.Close()
	redirecterCanceler <- true
	writerCanceler <- true
}
