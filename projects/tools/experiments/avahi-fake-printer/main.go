package main

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"github.com/holoplot/go-avahi"
	"log"
)

func main() {
	// Connect to the system bus
	conn, err := dbus.SystemBus()
	if err != nil {
		log.Fatalf("Failed to connect to system bus: %v", err)
	}

	// Create a new Avahi server
	server, err := avahi.ServerNew(conn)
	if err != nil {
		log.Fatalf("Failed to create Avahi server: %v", err)
	}

	// Create a new entry group
	group, err := server.EntryGroupNew()
	if err != nil {
		log.Fatalf("Failed to create Avahi entry group: %v", err)
	}

	// Define the printer service details
	printerServiceName := "Brother MFC-2710DW"
	serviceType := "_ipp._tcp"
	domainName := "local"
	hostName := ""
	port := 631

	// Define the TXT records to match the Brother MFC-2710DW capabilities
	txtRecords := []string{
		"txtvers=1",
		"qtotal=1",
		"ty=Brother MFC-2710DW",
		"product=(Brother MFC-2710DW)",
		"priority=20",
		"adminurl=http://[fe80::1%25en0]/",
		"note=Office Printer",
		"rp=ipp/print",
		"pdl=application/pdf,application/postscript,image/jpeg",
		"Color=T",
		"Duplex=T",
		"Copies=T",
		"Collate=T",
		"Scan=T",
		"Fax=T",
	}
	// Convert TXT records to the format required by Avahi
	var txt [][]byte
	for _, record := range txtRecords {
		txt = append(txt, []byte(record))
	}
	// Add the printer service to the entry group
	err = group.AddService(
		avahi.InterfaceUnspec, // Interface
		avahi.ProtoUnspec,     // Protocol
		0,                     // Flags
		printerServiceName,    // Name
		serviceType,           // Type
		domainName,            // Domain
		hostName,              // Host
		uint16(port),          // Port
		txt,                   // TXT records
	)
	if err != nil {
		log.Fatalf("Failed to add service: %v", err)
	}

	// Commit the entry group to announce the service
	err = group.Commit()
	if err != nil {
		log.Fatalf("Failed to commit entry group: %v", err)
	}

	fmt.Printf("Printer '%s' announced on the network.\n", printerServiceName)

	// Keep the program running to maintain the service announcement
	select {}
}
