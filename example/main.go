package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/NightRang3r/portlookup"
)

func main() {
	// Check if the port number is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide a port number.")
		return
	}

	// Convert the command-line argument to an integer
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid port number.")
		return
	}

	// Load the CSV file once and store the result in a map
	portMap, err := portlookup.LoadCSV("service-names-port-numbers.csv")
	if err != nil {
		log.Fatal("Error loading CSV:", err)
	}

	// Look up the service name and description for the given port
	serviceName, description := portlookup.LookupServiceByPort(portMap, port)

	// Check if the service was found
	if serviceName == nil && description == nil {
		fmt.Println("No service found for the given port.")
		return
	}

	// Use the service name and description as needed
	fmt.Printf("Service Name: %s\nDescription: %s\n", *serviceName, *description)
}
