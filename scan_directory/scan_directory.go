package main

import (
	"log"
	"os"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("scan_directory: ")
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetFlags(0)

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println("start scanning")
}
