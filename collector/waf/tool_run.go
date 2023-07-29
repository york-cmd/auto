package waf

import (
	"auto/commands"
	"auto/utils"
	"fmt"
)

func wafw00fRun() {
	result, err := utils.Command(commands.Config.Other.Wafw00f)
	if err != nil {
		fmt.Println(commands.Config.Other.Wafw00f)
		fmt.Println(result)
		utils.CheckError("wafw00fRun()", err)
	}
}
