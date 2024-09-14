package main

import (
	"fmt"
	"log"
	"go.bug.st/serial"
)

func main(){
	//list available serial ports
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}

	//if there weren't any serial ports available, throw a wobbly
	if len(ports) == 0 {
		log.Fatal("No serial ports found")
	}

	//if you found some ports, spit it out
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
}