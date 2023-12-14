package clients

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func ReadData(conn net.Conn) (Response, error) {
	reader := bufio.NewReader(conn)
	length, err := reader.Peek(5)
	if err != nil {
		return Response{}, fmt.Errorf("error reading length: %v", err)
	}
	lengthStr := string(length)
	messageLength, err := strconv.Atoi(lengthStr)
	if err != nil {
		return Response{}, fmt.Errorf("error converting length to int: %v", err)
	}
	reader.Discard(5)
	message := make([]byte, messageLength)
	_, err = reader.Read(message)
	if err != nil {
		return Response{}, fmt.Errorf("error reading message: %v", err)
	}
	messageStr := string(message)

	return parseResponse(messageStr), nil
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

func parseResponse(data string) Response {
	serviceName := data[:5]
	serviceStatus := data[5:7]
	dataFields := getDataFields(data[7:])

	return Response{
		Service: serviceName,
		Status:  serviceStatus,
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
