package main

import (
	"arqui-soa/clients"
	"fmt"
	"net"
)

const service = "testf"

func main() {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Create data fields to send
	var dataFields []clients.DataField
	value := clients.DataField{
		Key:   "hola",
		Value: service,
	}
	dataFields = append(dataFields, value)
	msg := clients.CreateMessage(service, dataFields)

	_, err = conn.Write(msg)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Read and process data from the server
	response, er := clients.ReadData(conn)
	if er != nil {
		fmt.Println("Error reading data:", er)
	}
	fmt.Println("Received service response:")
	fmt.Println("Service:", response.Service)
	fmt.Println("Status:", response.Status)
	fmt.Println("Data:", response.Data)
}
