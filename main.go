package main

import (
	"flag"
	"log"
	"os"
	"time"
)

func exampleData(data chan string) {
	time.Sleep(time.Second * 2)
	data <- "test"
	time.Sleep(time.Second)
	data <- "test2"
	time.Sleep(time.Second * 3)
	data <- "test3"
	data <- "test4"
	time.Sleep(time.Second)
	data <- "test5"
	time.Sleep(time.Second * 2)
	close(data)
}

func main() {
	// cmd line
	outputFlag := flag.String("output", "output.txt", "output file")
	exampleFlag := flag.Bool("example", false, "run example")
	flag.Parse()

	output, err := os.Create(*outputFlag)
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
	if *exampleFlag {
		go exampleData(data)
	}

	p := Peer{data, output, false, ""}
	p.Run()

	log.Println("Exiting...")
}
