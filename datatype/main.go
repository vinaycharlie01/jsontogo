package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func checkJSONDataType(data []byte) (string, error) {
	var parsed interface{}

	err := json.Unmarshal(data, &parsed)
	if err != nil {
		return "", err
	}

	dataType := reflect.TypeOf(parsed).String()
	return dataType, nil
}

func main() {
	// Example JSON data
	jsonData := []byte(`{"name": "John", "age": 30, "city": "New York"}`)

	dataType, err := checkJSONDataType(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("JSON data type:", dataType)
}
