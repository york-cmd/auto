package models

import "encoding/xml"

type MasscanResult struct {
	IP    string       `json:"ip"`
	Ports []PortDetail `json:"ports"`
}

type PortDetail struct {
	Port   int    `json:"port"`
	Proto  string `json:"proto"`
	Status string `json:"status"`
	Reason string `json:"reason"`
	TTL    int    `json:"ttl"`
}

type Nmap struct {
	XMLName xml.Name `xml:"nmaprun"`
	Host    host     `xml:"host"`
}

type host struct {
	Addr  Address `xml:"address"`
	Ports ports   `xml:"ports"`
}
type Address struct {
	IP string `xml:"addr,attr"`
}
type ports struct {
	Port []port `xml:"port"`
}
type port struct {
	Portid  string  `xml:"portid,attr"`
	Service service `xml:"service"`
}
type service struct {
	Name string `xml:"name,attr"`
}
type PortService struct {
	Port    string
	Service string
}
type ProtocolPort struct {
	Protocol string
	Port     string
}
