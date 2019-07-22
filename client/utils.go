package client

import (
	"bufio"
	"chat_client/config"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

//Scans messages from connection and pushes it to console
func PrintMessages(conn net.Conn) {
	serverAddr := conn.RemoteAddr().String()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		fmt.Fprintln(os.Stdout, serverAddr+config.Separator+input.Text())
	}
	log.Println("Server connection lost")
	os.Exit(1)
}

//Sends greetings to server
func SayHello(c net.Conn) (int, error) {
	return c.Write([]byte("Client says hello!!!\n"))
}

func ReadFromCopyTo(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		config.FailOnError(err, "while messaging with server")
	}
}
