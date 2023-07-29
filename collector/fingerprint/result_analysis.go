package fingerprint

import (
	"auto/collector/result"
	"auto/models"
	"auto/utils"
	"strings"
)

func tideFingerResult() {
	var urlFingerMap = make(map[string]string)
	filename := "tmp/TideFinger/TideFinger.txt"
	lines, err := utils.ReadTextFileToSlice(filename)
	if err != nil {
		utils.CheckError("tideFingerResult()", err)
	}
	for _, line := range lines {
		if strings.Contains(line, "[") {
			split := strings.Split(line, ",")
			url := strings.TrimSpace(split[4])
			finger := strings.TrimSpace(split[3])
			urlFingerMap[url] = finger
		}
	}
	for url, data := range result.WebInfo {
		result.WebInfo[url] = models.WebInfo{Title: data.Title, Fingerprint: urlFingerMap[url]}
	}
}
