package scanner

import (
	"auto/utils"
	"encoding/json"
	"fmt"
)

var (
	htmlResultPath = "result/vuln"
	templatePath   = "tools/template/template.html"
)

func getResult() {
	XrayResult()
	nucleiResult()
	xscanResult()
	nucleiToXray()
	xscanToXray()
}
func getHtmlResult() {
	getResult()
	for key, datas := range results {
		filename := fmt.Sprintf("%v/%v.html", htmlResultPath, key)
		utils.CopyFile(filename, templatePath)
		var str []string
		for _, data := range datas {
			jsonData, err := json.Marshal(data)
			if err != nil {
				utils.CheckError("wafw00fResult()", err)
			}
			str = append(str, fmt.Sprintf(temp, string(jsonData)))
		}
		utils.AppendToFile(filename, str)
	}
}
