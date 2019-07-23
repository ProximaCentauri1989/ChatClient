package failer

import (
	"log"
	"os"
)

func FailOnError(err error, action string) {
	if err != nil {
		log.Println("Error:", err, "\nAction: ", action)
		os.Exit(1)
	}
}
