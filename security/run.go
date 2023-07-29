package security

import (
	"auto/collector"
	"auto/scanner"
)

func Run(domains []string, alterxUse bool) {
	// 信息收集
	collector.Run(domains, alterxUse)
	// 漏洞扫描
	scanner.MultipleRun("result/httpx/urls.txt")
}
