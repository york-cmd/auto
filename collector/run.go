package collector

import (
	"auto/collector/cdn"
	"auto/collector/fingerprint"
	"auto/collector/portscan"
	"auto/collector/result"
	"auto/collector/subdomain"
	"auto/collector/urlvalid"
	"auto/collector/waf"
	"log"
)

func Run(domains []string, alterxUse bool) {
	log.Println("start collecting information ...")
	subdomain.Run(domains, alterxUse)
	cdn.Run(subdomain.Subdomains)
	portscan.Run()
	urlvalid.Run()
	fingerprint.Run()
	waf.Run()
	result.Save()
	log.Println("information collection ended .")
}
