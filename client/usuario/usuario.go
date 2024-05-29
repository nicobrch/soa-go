package main

import (
	"fmt"
	client "soa/client"
)

func main() {
	const busAddr = "localhost:5000"

	c := client.NewClient(busAddr)
	defer c.Conn.Close()

	service := "users"
	data := "init"

	transaction, err := client.FormatSOATransaction(service, data)

	if err != nil {
		fmt.Println("Error formatting transaction:", err)
		return
	}

	c.SendTransaction(transaction)

	response := c.ReceiveMessage()
	fmt.Println("Received response:", response)
}
