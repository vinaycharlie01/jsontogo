package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
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

func formatNumber(str string) string {
	// Replace digits with their corresponding strings using the numbers map
	numbers := map[string]string{
		"0": "Zero_", "1": "One_", "2": "Two_", "3": "Three_",
		"4": "Four_", "5": "Five_", "6": "Six_", "7": "Seven_",
		"8": "Eight_", "9": "Nine_",
	}

	for _, char := range str {
		if isDigit(string(char)) {
			str = strings.Replace(str, string(char), numbers[string(char)], 1)
		}
	}

	return str
}

func isDigit(s string) bool {
	// Check if a string represents a digit
	re := regexp.MustCompile(`^\d$`)
	return re.MatchString(s)
}

func chekDatatype() {

}

func main() {
	data := `{"a":123,"b":12.3,"c":"123","d":"12.3","e":true}`
	var resultMap map[string]interface{}
	if err := json.Unmarshal([]byte(data), &resultMap); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert numeric values to the desired types
	for _, value := range resultMap {
		switch value.(type) {
		case int:
			fmt.Println("Hello")
		}
	}

	// fmt.Printf("%+v\n", resultMap)

}
