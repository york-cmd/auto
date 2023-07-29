package subdomain

import (
	"auto/commands"
	"auto/utils"
	"fmt"
)

// oneforallRun ksubEnumRun subfinderRun alterxRun ksubVerifyRun
// 执行命令进行子域名收集
func oneforallRun(domain string) {
	code := fmt.Sprintf(commands.Config.Subdomain.OneForAll, domain)
	result, err := utils.Command(code)
	if err != nil {
		fmt.Println(code)
		fmt.Println(result)
		utils.CheckError("oneforallRun()", err)
	}
}
func ksubEnumRun(domain string) {
	code := fmt.Sprintf(commands.Config.Subdomain.KsubEnum, domain)
	result, err := utils.Command(code)
	if err != nil {
		fmt.Println(code)
		fmt.Println(result)
		utils.CheckError("ksubEnumRun()", err)
	}
}
func subfinderRun(domain string) {
	code := fmt.Sprintf(commands.Config.Subdomain.Subfinder, domain)
	result, err := utils.Command(code)
	if err != nil {
		fmt.Println(code)
		fmt.Println(result)
		utils.CheckError("subfinderRun()", err)
	}
}
func alterxRun() {
	result, err := utils.Command(commands.Config.Subdomain.Alterx)
	if err != nil {
		fmt.Println(commands.Config.Subdomain.Alterx)
		fmt.Println(result)
		utils.CheckError("alterxRun()", err)
	}
}
func ksubVerifyRun() {
	result, err := utils.Command(commands.Config.Subdomain.KsubVerify)
	if err != nil {
		fmt.Println(commands.Config.Subdomain.KsubVerify)
		fmt.Println(result)
		utils.CheckError("ksubVerifyRun()", err)
	}
}
