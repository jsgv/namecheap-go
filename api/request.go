package api

import (
	"encoding/xml"
)

type ApiResponse struct {
	XMLName           xml.Name `xml:"ApiResponse" json:"-"`
	Status            string   `xml:"Status,attr"`
	Xmlns             string   `xml:"xmlns,attr"`
	RequestedCommand  string   `xml:"RequestedCommand"`
	Server            string   `xml:"Server"`
	GMTTimeDifference string   `xml:"GMTTimeDifference"`
	ExecutionTime     string   `xml:"ExecutionTime"`
	Errors            struct {
		XMLName xml.Name `xml:"Errors" json:"-"`
		Error   struct {
			XMLName xml.Name `xml:"Error" json:"-"`
			Number  int      `xml:"Number,attr" json:",omitempty"`
			Error   string   `xml:",chardata" json:",omitempty"`
		}
	}
	Warnings struct {
		XMLName xml.Name `xml:"Warnings" json:"-"`
		Warning struct {
			XMLName xml.Name `xml:"Warning" json:"-"`
			Number  int      `xml:"Number,attr" json:",omitempty"`
			Warning string   `xml:",chardata" json:",omitempty"`
		}
	}
}

type Paging struct {
	TotalItems  int `xml:"TotalItems"`
	CurrentPage int `xml:"CurrentPage"`
	PageSize    int `xml:"PageSize"`
}
