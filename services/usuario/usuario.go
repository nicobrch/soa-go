package main

import (
	"fmt"
	service "soa/services"
)

func Menu(data string) string {
	if data == "init" {
		msg := "Menu de opciones\n 1. Crear Usuario\n 2. Listar Usuarios\n 3. Salir\n"
		return msg
	} else if data == "1" {
		msg := "Ingrese el nombre , edad y correo del usuario:\n"
		return msg
	} else if data == "2" {
		msg := "Listado de usuarios\n"
		return msg
	} else if data == "3" {
		msg := "Saliendo..."
		return msg
	} else if len(service.GetDataFields(data)) == 3 {
		msg := "Usuario creado con éxito\n"
		return msg
	}
	return "Opción no válida"
}

func main() {
	const busAddr = "localhost:5000"

	s := service.NewService("users", busAddr)
	defer s.Conn.Close()

	response, err := s.InitService()
	if err != nil {
		fmt.Println("Error sending transaction:", err)
		return
	}

	fmt.Println("Received response:", response)
	// Logic to process data
	s.ProcessData(Menu)
}
