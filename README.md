# json-to-go




```go

func main() {

	jsonStr := `{"name": "John", "time": "2023-01-01T12:34:56","age":1, "city": "New York"}`
	typeName := "Person"

	res, _ := jsonToGo(jsonStr, typeName, false, false, true)
	fmt.Println(res)
}
 
//  output
type Person struct {
        Name string`json:"name,omitempty"`
        Time time.Time`json:"time,omitempty"`
        Age float64`json:"age,omitempty"`
        City string`json:"city,omitempty"`
}
```