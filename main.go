package main

import (
	"fmt"
	"log"

	"go.bug.st/serial/enumerator"
)

const NanoVID = "0483"
const NanoPID = "5740"

var workingNanoPort string

func main() {

	//Find available serial ports
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}

	//If there weren't any serial ports available, throw a wobbly
	if len(ports) == 0 {
		log.Fatal("No serial ports found")
	}

	//Figure out which device is a NanoVNA and store that port as the working port.
	//If multiple NanoVNAs are discovered, ask the user to determine which one they want to use.

	var nanoPorts []string
	for _, port := range ports {

		//Check how many ports have the right VID and PID
		if port.IsUSB && port.VID == NanoVID && port.PID == NanoPID {

			//Add port addresses to the list of plugged in NanoVNAs
			nanoPorts = append(nanoPorts, port.Name)
		}
	}

	//Check how many NanoVNAs were found.  If more than one was found, ask the user to choose one.
	if len(nanoPorts) > 1 {
		fmt.Printf("More than one NanoVNA is plugged in.\n")
		fmt.Printf("Which one would you like to use?\n")
		fmt.Printf("Type one number and press enter:\n")
		for i, nanoPort := range nanoPorts {
			fmt.Printf("\t%v) %v\n", i+1, nanoPort)
		}

		//Create variable to store user's answer
		var userChoice int
		fmt.Scanln(&userChoice)

		//set the working port to be that which the user chose.
		workingNanoPort = nanoPorts[userChoice-1]

	} else if len(nanoPorts) == 1 {
		//If only one NanoVNA is found, just use that one.
		workingNanoPort = nanoPorts[0]
	} else {
		//No NanvoVNAs found plugged in so throw a wobbly
		log.Fatal("No NanoVNAs discovered.  Please plug in a NanoVNA and try again.\n")
	}

	fmt.Printf("The NanoVNA that this program is using is connected at:\n\t%v\n", workingNanoPort)

	//set serial mode
	// mode := &serial.Mode{
	// 	BaudRate: 115200,
	// }

	//open serial port
	// port, err := serial.Open(nanoPort, mode)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("port opened: %v\n", port)
}
