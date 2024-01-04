package main

type EventData1 []struct {
	Analysis Analysis `json:"analysis,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
}
type Analysis struct {
	Active       int    `json:"active,omitempty"`
	DpvCmra      string `json:"dpv_cmra,omitempty"`
	DpvFootnotes int    `json:"dpv_footnotes,omitempty"`
	DpvMatchCode int    `json:"dpv_match_code,omitempty"`
	DpvVacant    string `json:"dpv_vacant,omitempty"`
}
type Metadata struct {
	CarrierRoute          string  `json:"carrier_route,omitempty"`
	CongressionalDistrict string  `json:"congressional_district,omitempty"`
	CountyFips            string  `json:"county_fips,omitempty"`
	CountyName            string  `json:"county_name,omitempty"`
	Dst                   bool    `json:"dst,omitempty"`
	ElotSequence          string  `json:"elot_sequence,omitempty"`
	ElotSort              string  `json:"elot_sort,omitempty"`
	Latitude              float64 `json:"latitude,omitempty"`
	Longitude             float64 `json:"longitude,omitempty"`
	Precision             string  `json:"precision,omitempty"`
	Rdi                   string  `json:"rdi,omitempty"`
	RecordType            string  `json:"record_type,omitempty"`
	TimeZone              string  `json:"time_zone,omitempty"`
	UtcOffset             int     `json:"utc_offset,omitempty"`
	ZipType               string  `json:"zip_type,omitempty"`
}
