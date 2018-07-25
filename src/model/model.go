package model

import "fmt"

type ParserResponse struct {
	OS         string `json:"OS"`
	DeviceType string `json:"DeviceType"`
	Browser    string `json:"Browser"`
	Country    string `json:"Country"`
	Domain     string `json:"Domain"`
}

type ParserError struct {
	ErrMsg string `json:"error"`
}

func (p *ParserError) Error() string {
	return fmt.Sprintf("Error %s", p.ErrMsg)
}
