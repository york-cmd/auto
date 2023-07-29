package subdomain

import (
	"log"
	"sync"
)

var (
	wg sync.WaitGroup
)
var (
	numOneForAll  int
	numSubfinder  int
	numKsubdomain int
	numThreeTools int
	numAlterx     int
	numFourTools  int
)

// Run 子域名收集主函数
func Run(domains []string, alterxUse bool) {
	log.Println("start subdomain collection ...")
	for _, domain := range domains {
		subdomainRun(domain)
	}
	GetResults(alterxUse)
}
