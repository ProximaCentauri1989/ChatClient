package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//Reads messages from channel and sends it to connection
func ConnectionWriter(conn net.Conn, ch <-chan string, quit <-chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			for msg := range ch {
				fmt.Fprintln(conn, msg)
			}
		}
	}
}

//Scans messages from console and sends it to channel
func RedirectToChannel(ch chan string, quit chan bool) {
	msgScanner := bufio.NewScanner(os.Stdin)
	for {
		msgScanner.Scan()
		select {
		case <-quit:
			return
		default:
			ch <- msgScanner.Text()
		}
	}
}
