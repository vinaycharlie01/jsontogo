# json-to-go

### its genarate struct field like below.


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

```