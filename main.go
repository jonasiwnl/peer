package main

import (
	"log"
	"os"
	"time"
)

func testData(data chan string) {
	time.Sleep(time.Second * 2)
	data <- "test"
	time.Sleep(time.Second)
	data <- "test2"
	time.Sleep(time.Second * 3)
	close(data)
}

func main() {
	output, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}

	// close output file on exit and check for its returned error
	defer func() {
		if err := output.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	data := make(chan string)
	go testData(data)

	p := Peer{data, output, false, ""}
	p.Run()

	log.Println("Exiting...")
}
