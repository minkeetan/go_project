package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("scan_directory: ")
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetFlags(0)

	/*To set the log output */
	// file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.SetOutput(file)

	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())

			if !info.IsDir() {
				/* To read file at once, use os.readfile*/
				// dat, err := os.ReadFile(path)
				// if err != nil {
				// 	log.Fatal(err)
				// }

				// fmt.Print(string(dat))

				/* To read file line by line, use bufio scanner*/
				f, _ := os.Open(path)
				scanner := bufio.NewScanner(f)

				for scanner.Scan() {
					fmt.Println(scanner.Text())
				}
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}

}
