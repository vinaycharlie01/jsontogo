# json-to-go



### jsonToStruct
Golang code generator for creating struct from json.

### INSTALLATION
go get github.com/vinaycharlie01/jsontogo


## Introduction

The JSON to Go Struct Converter is a Go code designed to generate Go struct definitions from JSON input. It aims to simplify the process of creating Go structs corresponding to JSON data structures

## Features

- Automatic generation of Go struct definitions from JSON input.
- Configurable options for struct naming, flattening, and field tags.
- Handling of nested structures, arrays, and basic data types.
- Deduplication of struct types based on field names and types.
- Support for "omitempty" field tags based on input data.

```go

 
//  output
package main

type EventData1 []struct {
	Metadata Metadata `json:"metadata,omitempty"`
	Analysis Analysis `json:"analysis,omitempty"`
}
type Metadata struct {
	ElotSort              string  `json:"elot_sort,omitempty"`
	Latitude              float64 `json:"latitude,omitempty"`
	TimeZone              string  `json:"time_zone,omitempty"`
	CountyFips            string  `json:"county_fips,omitempty"`
	CountyName            string  `json:"county_name,omitempty"`
	CarrierRoute          string  `json:"carrier_route,omitempty"`
	UtcOffset             int     `json:"utc_offset,omitempty"`
	RecordType            string  `json:"record_type,omitempty"`
	ElotSequence          string  `json:"elot_sequence,omitempty"`
	Longitude             float64 `json:"longitude,omitempty"`
	ZipType               string  `json:"zip_type,omitempty"`
	CongressionalDistrict string  `json:"congressional_district,omitempty"`
	Rdi                   string  `json:"rdi,omitempty"`
	Precision             string  `json:"precision,omitempty"`
	Dst                   bool    `json:"dst,omitempty"`
}
type Analysis struct {
	DpvMatchCode int    `json:"dpv_match_code,omitempty"`
	DpvFootnotes int    `json:"dpv_footnotes,omitempty"`
	DpvCmra      string `json:"dpv_cmra,omitempty"`
	DpvVacant    string `json:"dpv_vacant,omitempty"`
	Active       int    `json:"active,omitempty"`
}
### Done
```
