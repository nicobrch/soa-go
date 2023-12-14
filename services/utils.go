package services

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func ReadRequest(conn net.Conn) (Request, error) {
	reader := bufio.NewReader(conn)
	length, err := reader.Peek(5)
	if err != nil {
		return Request{}, fmt.Errorf("error reading length: %v", err)
	}
	lengthStr := string(length)
	messageLength, err := strconv.Atoi(lengthStr)
	if err != nil {
		return Request{}, fmt.Errorf("error converting length to int: %v", err)
	}
	_, err = reader.Discard(5)
	if err != nil {
		return Request{}, err
	}
	message := make([]byte, messageLength)
	_, err = reader.Read(message)
	if err != nil {
		return Request{}, fmt.Errorf("error reading message: %v", err)
	}
	messageStr := string(message)

	return parseRequest(messageStr), nil
}

func CreateMessage(service string, dataFields []DataField) []byte {
	var pairs []string
	for _, df := range dataFields {
		pair := fmt.Sprintf("%s-%s", df.Key, df.Value)
		pairs = append(pairs, pair)
	}
	data := strings.Join(pairs, "_")
	length := len(service + data)
	lengthStr := fmt.Sprintf("%05d", length)
	result := lengthStr + service + data
	return []byte(result)
}

func parseRequest(data string) Request {
	serviceName := data[:5]

	if serviceName == "sinit" {
		var fields []DataField
		status := DataField{
			Key:   "status",
			Value: data[5:7],
		}
		fields = append(fields, status)
		return Request{
			Service: data[7:],
			Data:    fields,
		}
	} else if data[5:7] == "OK" || data[5:7] == "NK" {
		dataFields := getDataFields(data[7:])
		return Request{
			Service: data[7:],
			Data:    dataFields,
		}
	}

	dataFields := getDataFields(data[5:])
	return Request{
		Service: serviceName,
		Data:    dataFields,
	}
}

func getDataFields(data string) []DataField {
	pairs := strings.Split(data, "_")
	var dataFields []DataField

	for _, pair := range pairs {
		kv := strings.Split(pair, "-")
		if len(kv) == 2 {
			dataFields = append(dataFields, DataField{Key: kv[0], Value: kv[1]})
		} else {
			fmt.Println("Invalid pair:", pair)
		}
	}

	return dataFields
}
