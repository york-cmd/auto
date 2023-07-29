package models

type WebInfo struct {
	Title       string
	Fingerprint string
	Waf         string
}
type Wafw00f struct {
	URL          string `json:"url"`
	Detected     bool   `json:"detected"`
	Firewall     string `json:"firewall"`
	Manufacturer string `json:"manufacturer"`
}
type JsonResult struct {
	Url         string `json:"Url"`
	Title       string `json:"Title"`
	Fingerprint string `json:"Fingerprint"`
	Waf         string `json:"Waf"`
}
