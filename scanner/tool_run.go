package scanner

import (
	"auto/commands"
	"auto/utils"
	"fmt"
	"log"
	"sync"
)

var (
	wg  sync.WaitGroup
	pwd = utils.Pwd()
)

func nucleiMultipleRun(filename string) {
	code := fmt.Sprintf(commands.Config.Vulnscan.NucleiMultiple, filename)
	result, err := utils.Command(code)
	if err != nil {
		fmt.Println(code)
		fmt.Println(result)
		utils.CheckError("nucleiMultipleRun()", err)
	}
}
func nucleiSingleRun(target string) {
	code := fmt.Sprintf(commands.Config.Vulnscan.NucleiSingle, target)
	result, err := utils.Command(code)
	if err != nil {
		fmt.Println(code)
		fmt.Println(result)
		utils.CheckError("nucleiSingleRun()", err)
	}
}
func XrayRun() {
	code := fmt.Sprintf(commands.Config.Vulnscan.XrayListen, pwd, pwd)
	dir := fmt.Sprintf("%s/tools/vuln/xray", pwd)
	result, err := utils.CommandByDir(code, dir)
	if err != nil {
		fmt.Println(code)
		fmt.Println(result)
		utils.CheckError("XrayRun()", err)
	}
}
func xscanSingleRun(target string) {
	code := fmt.Sprintf(commands.Config.Vulnscan.XscanSingle, target, pwd)
	dir := fmt.Sprintf("%s/tools/vuln/xscan", pwd)
	result, err := utils.CommandByDir(code, dir)
	if err != nil {
		fmt.Println(code)
		fmt.Println(result)
		utils.CheckError("xscanSingleRun()", err)
	}
}
func xscanMultipleRun(filename string) {
	filePath := fmt.Sprintf("%s/%s", pwd, filename)
	code := fmt.Sprintf(commands.Config.Vulnscan.XscanMultiple, filePath, pwd)
	dir := fmt.Sprintf("%s/tools/vuln/xscan", pwd)
	result, err := utils.CommandByDir(code, dir)
	if err != nil {
		fmt.Println(code)
		fmt.Println(result)
		utils.CheckError("xscanMultipleRun()", err)
	}
}
func crawlergoRun(targets []string) {
	var codes []string
	for _, target := range targets {
		codes = append(codes, fmt.Sprintf(commands.Config.Vulnscan.Crawlergo, target))
	}
	limit := make(chan struct{}, commands.Config.Goroutine.Crawlergo)
	for _, code := range codes {
		wg.Add(1)
		limit <- struct{}{}
		go func(code string) {
			defer wg.Done()
			result, err := utils.Command(code)
			if err != nil {
				fmt.Println(code)
				fmt.Println(result)
				utils.KillProcess("xray")
				utils.CheckError("crawlergoRun()", err)
			}
			<-limit
		}(code)
	}
	wg.Wait()
	close(limit)
}
func crawlergoToXrayRun(targets []string) {
	log.Println("start xray listen ...")
	go XrayRun()
	log.Println("start crawling with Crawlergo ...")
	crawlergoRun(targets)
	log.Println("end of crawling . ")
	log.Println("wait for the xray vulnerability scan to complete .")
	checkFileNotModified()
	utils.KillProcess("xray")
	log.Println("xray vulnerability scan ended .")
}
