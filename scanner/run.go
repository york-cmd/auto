package scanner

import (
	"auto/utils"
	"log"
)

func SingleRun(target string) {
	log.Println("start vulnerability scanning with xscan ...")
	xscanSingleRun(target)
	log.Println("xscan vulnerability scan ended .")
	log.Println("start vulnerability scanning with nuclei ...")
	nucleiSingleRun(target)
	log.Println("nuclei vulnerability scan ended .")
	crawlergoToXrayRun([]string{target})
	getHtmlResult()
	printResults()
}
func MultipleRun(filename string) {
	log.Println("start vulnerability scanning with xscan ...")
	xscanMultipleRun(filename)
	log.Println("xscan vulnerability scan ended .")
	log.Println("start vulnerability scanning with nuclei ...")
	nucleiMultipleRun(filename)
	log.Println("nuclei vulnerability scan ended .")
	targets, _ := utils.ReadTextFileToSlice(filename)
	websiteNum = len(targets)
	crawlergoToXrayRun(targets)
	getHtmlResult()
	printResults()
}
