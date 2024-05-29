package client

import (
	"fmt"
	"net"
	"os"
)

type Client struct {
	Conn net.Conn
}

func NewClient(busAddr string) *Client {
	conn, err := net.Dial("tcp", busAddr)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		os.Exit(1)
	}
	return &Client{Conn: conn}
}

func FormatSOATransaction(service string, data string) (string, error) {
	if len(service) > 5 {
		return "", fmt.Errorf("service name must be no longer than 5 characters")
	}
	message := service + data
	length := len(message)
	formattedMessage := fmt.Sprintf("%05d%s", length, message)
	return formattedMessage, nil
}

func (c *Client) SendTransaction(transaction string) {
	fmt.Println("Sending transaction:", transaction)
	_, err := c.Conn.Write([]byte(transaction))
	if err != nil {
		fmt.Println("Error sending transaction:", err)
		os.Exit(1)
	}
}

func (c *Client) ReceiveMessage() string {
	buffer := make([]byte, 1024)
	n, err := c.Conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
		os.Exit(1)
	}
	return string(buffer[:n])
}
