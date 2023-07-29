package portscan

import (
	"auto/commands"
	"auto/utils"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func masscanRun() {
	result, err := utils.Command(commands.Config.Portscan.Masscan)
	if err != nil {
		fmt.Println(commands.Config.Portscan.Masscan)
		fmt.Println(result)
		utils.CheckError("masscanRun()", err)
	}
	masscanResult()
	firewallJudgment()
}
func nmapRun() {
	getNmapCode()
	limit := make(chan struct{}, commands.Config.Goroutine.Nmap)
	for _, code := range nmapCodes {
		wg.Add(1)
		limit <- struct{}{}
		go func(code string) {
			defer wg.Done()
			result, err := utils.Command(code)
			if err != nil {
				fmt.Println(code)
				fmt.Println(result)
				utils.CheckError("nmapRun()", err)
			}
			<-limit
		}(code)
	}
	wg.Wait()
	close(limit)
	nmapResult()
}
