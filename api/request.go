package api

import (
	"encoding/xml"
)

type ApiResponse struct {
	XMLName           xml.Name `xml:"ApiResponse" json:"-"`
	Status            string   `xml:"Status,attr"`
	RequestedCommand  string   `xml:"RequestedCommand"`
	Server            string   `xml:"Server"`
	GMTTimeDifference string   `xml:"GMTTimeDifference"`
	ExecutionTime     string   `xml:"ExecutionTime"`
	Errors            struct {
		XMLName xml.Name `xml:"Errors" json:"-"`
		Error   []struct {
			XMLName xml.Name `xml:"Error" json:"-"`
			Number  string   `xml:"Number,attr"`
			Error   string   `xml:",chardata"`
		}
	}
	// TODO
	// Warnings struct {
	// 	XMLName xml.Name `xml:"Warnings" json:"-"`
	// }
}
