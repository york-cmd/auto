package portscan

import (
	"auto/collector/cdn"
	"log"
)

func Run() {
	log.Println("start port scan ...")
	if len(cdn.NoCDNIPs) != 0 {
		getIPlistFile()
		masscanRun()
		nmapRun()
		getServices()
	} else {
		getYesCDNWebServices()
	}
	log.Println("port scan complete .")
	printResults()
}
