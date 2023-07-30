package scanner

import (
	"auto/commands"
	"auto/utils"
	"log"
)

func SingleRun(target string) {
	log.Println("start a vulnerability scan ...")
	if commands.Config.ScannerTools.Xscan {
		log.Println("start vulnerability scanning with xscan ...")
		xscanSingleRun(target)
		log.Println("xscan vulnerability scan ended .")
	}
	if commands.Config.ScannerTools.Nuclei {
		log.Println("start vulnerability scanning with nuclei ...")
		nucleiSingleRun(target)
		log.Println("nuclei vulnerability scan ended .")
	}
	if commands.Config.ScannerTools.Xray {
		crawlergoToXrayRun([]string{target})
	}
	getHtmlResult()
	printResults()
	log.Println("vulnerability scan ended .")
}
func MultipleRun(filename string) {
	log.Println("start a vulnerability scan ...")
	if commands.Config.ScannerTools.Xscan {
		log.Println("start vulnerability scanning with xscan ...")
		xscanMultipleRun(filename)
		log.Println("xscan vulnerability scan ended .")
	}
	if commands.Config.ScannerTools.Nuclei {
		log.Println("start vulnerability scanning with nuclei ...")
		nucleiMultipleRun(filename)
		log.Println("nuclei vulnerability scan ended .")
	}
	targets, _ := utils.ReadTextFileToSlice(filename)
	websiteNum = len(targets)
	if commands.Config.ScannerTools.Xray {
		crawlergoToXrayRun(targets)
	}
	getHtmlResult()
	printResults()
	log.Println("vulnerability scan ended .")
}
