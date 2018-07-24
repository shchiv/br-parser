package model

type ParserResponse struct {
	OS         string `json:"OS"`
	DeviceType string `json:"DeviceType"`
	Browser    string `json:"Browser"`
	Country    string `json:"Country"`
	Domain     string `json:"Domain"`
}
