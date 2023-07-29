package waf

import (
	"auto/collector/result"
	"auto/models"
	"auto/utils"
	"encoding/json"
	"io"
	"os"
)

func wafw00fResult() {
	path := "tmp/wafw00f/waf.json"
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		utils.CheckError("wafw00fResult()", err)
	}
	jsonData, err := io.ReadAll(file)
	if err != nil {
		utils.CheckError("wafw00fResult()", err)
	}
	var wafData []models.Wafw00f
	err = json.Unmarshal(jsonData, &wafData)
	if err != nil {
		utils.CheckError("wafw00fResult()", err)
	}
	for _, data := range wafData {
		result.WebInfo[data.URL] = models.WebInfo{Title: result.WebInfo[data.URL].Title, Fingerprint: result.WebInfo[data.URL].Fingerprint, Waf: data.Firewall}
	}
}
