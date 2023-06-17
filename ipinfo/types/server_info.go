package types

type ServerInfo struct {
	Ip      string      `json:"-"`
	Score   int64       `json:"-"`
	Country string      `json:"country"`
	ASN     string      `json:"asn"`
	Status  *JSONStatus `json:"dc_status"`
}
