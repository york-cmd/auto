package fingerprint

import (
	"auto/commands"
	"auto/utils"
	"fmt"
)

func tideFingerRun() {
	result, err := utils.Command(commands.Config.Other.TideFinger)
	if err != nil {
		fmt.Println(commands.Config.Other.TideFinger)
		fmt.Println(result)
		utils.CheckError("tideFingerRun()", err)
	}
}
