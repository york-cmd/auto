package urlvalid

import (
	"auto/commands"
	"auto/utils"
	"fmt"
)

func httpxRun() {
	result, err := utils.Command(commands.Config.Other.Httpx)
	if err != nil {
		fmt.Println(commands.Config.Other.Httpx)
		fmt.Println(result)
		utils.CheckError("httpxRun()", err)
	}
}
