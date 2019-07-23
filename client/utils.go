package client

import (
	"ChatClient/config"
	"ChatClient/failer"
	"bufio"
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
}

//Sends greetings to server
func SayOneMessage(c net.Conn, msg string) (int, error) {
	return c.Write([]byte(msg))
}

func ReadFromCopyTo(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		failer.FailOnError(err, "while messaging with server")
	}
}
