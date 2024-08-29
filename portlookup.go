package portlookup

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// PortService contains the service name and description for a port
type PortService struct {
	ServiceName string
	Description string
}

// CSV: https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.csv
// LoadCSV loads the CSV data and returns a map of port numbers to PortService
func LoadCSV(filename string) (map[int]PortService, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV file: %w", err)
	}

	portMap := make(map[int]PortService)

	for i, record := range records[1:] { // skip the header
		if len(record) < 4 {
			return nil, fmt.Errorf("invalid record format in line %d: expected at least 4 fields, got %d", i+2, len(record))
		}

		port, err := strconv.Atoi(record[1])
		if err != nil {
			continue // Skip invalid port numbers
		}

		// Check if the port is already in the map
		if _, exists := portMap[port]; !exists {
			serviceName := record[0]
			description := record[3]

			if serviceName == "" {
				serviceName = "Unknown Service"
			}
			if description == "" {
				description = "No description available."
			}

			portMap[port] = PortService{
				ServiceName: serviceName,
				Description: description,
			}
		}
	}

	return portMap, nil
}

// LookupServiceByPort returns pointers to the service name and description for a given port
func LookupServiceByPort(portMap map[int]PortService, port int) (*string, *string) {
	if service, exists := portMap[port]; exists {
		return &service.ServiceName, &service.Description
	}
	return nil, nil
}
