package collector

import (
	"auto/collector/cdn"
	"auto/collector/fingerprint"
	"auto/collector/portscan"
	"auto/collector/result"
	"auto/collector/subdomain"
	"auto/collector/urlvalid"
	"auto/collector/waf"
	"auto/utils"
	"fmt"
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
	utils.SendMessageToServerChan("信息收集完成, 请查看站点信息 !", fmt.Sprintf("站点数量: %v", len(result.WebInfo)))
	log.Println("information collection ended .")
}
