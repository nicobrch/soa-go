package main

import (
	"arqui-soa/services"
	"fmt"
	"net"
)

const service = "testf"

func main() {
	conn, err := net.Dial("tcp", "soabus:5000")
	if err != nil {
		fmt.Println("Error creating connection:", err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error on defer function:", err)
			return
		}
	}(conn)

	data := []byte("00010sinit" + service)
	fmt.Println("Initializing service:", string(data))

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	initMsg, err := services.ReadRequest(conn)
	if err != nil {
		fmt.Println("Error reading request:", err)
	}
	if initMsg.Data[0].Key == "status" && initMsg.Data[0].Value == "OK" {
		fmt.Println("Initialization succesful.")
	} else {
		fmt.Println("Could not initialize the service.")
	}

	for {
		fmt.Println("Waiting for transactions...")
		request, err := services.ReadRequest(conn)
		if err != nil {
			fmt.Println("Error reading request:", err)
		}
		if request.Service != service {
			continue
		}

		fmt.Println("Received client transaction:", request)

		var data []services.DataField

		value := services.DataField{
			Key:   "respuesta",
			Value: "se recibio mensaje en testf exitosamente",
		}
		data = append(data, value)

		response := services.CreateMessage(service, data)

		_, err = conn.Write(response)
		if err != nil {
			fmt.Println("Error writing response:", err)
			return
		}
	}
}
