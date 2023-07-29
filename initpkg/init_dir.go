package initpkg

import "auto/utils"

func initCollectorDir() {
	utils.Dir("tmp/subdomain")
	utils.Dir("tmp/subdomain/oneforall")
	utils.Dir("tmp/portscan")
	utils.Dir("tmp/portscan/nmap")
	utils.Dir("tmp/TideFinger")
	utils.Dir("tmp/wafw00f")
	utils.Dir("result")
	utils.Dir("result/subdomains")
	utils.Dir("result/webinfo")
	utils.Dir("result/services")
	utils.Dir("result/httpx")
}
func initScannerDir() {
	utils.Dir("result")
	utils.Dir("result/vuln")
	utils.Dir("result/vuln/tools")
	utils.Dir("tmp")
	utils.Dir("tmp/vuln")
}
func initSecurityDir() {
	utils.Dir("tmp/subdomain")
	utils.Dir("tmp/subdomain/oneforall")
	utils.Dir("tmp/portscan")
	utils.Dir("tmp/portscan/nmap")
	utils.Dir("tmp/TideFinger")
	utils.Dir("tmp/wafw00f")
	utils.Dir("tmp/vuln")
	utils.Dir("result")
	utils.Dir("result/subdomains")
	utils.Dir("result/webinfo")
	utils.Dir("result/services")
	utils.Dir("result/httpx")
	utils.Dir("result/vuln")
	utils.Dir("result/vuln/tools")
}
