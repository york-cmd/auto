package scanner

import (
	"auto/models"
	"auto/utils"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
)

var (
	xrayHtmlResultPath   = "result/vuln/xray.html"
	nucleiJsonResultPath = "result/vuln/tools/nuclei.json"
	xscanJsonResultPath  = "result/vuln/tools/xscan.json"
)

var (
	xrayNum int
	nuclei  []models.Nuclei
	xscan   []models.Xscan
	results = make(map[string][]models.Xray)
)

func XrayResult() {
	if !utils.FileExists(xrayHtmlResultPath) {
		return
	}
	content, err := os.ReadFile(xrayHtmlResultPath)
	if err != nil {
		utils.CheckError("XrayResult()", err)
	}
	pattern := `<script class='web-vulns'>webVulns\.push\(.+?\)</script>`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(string(content), -1)
	xrayNum = len(matches)
}

func nucleiResult() {
	if !utils.FileExists(nucleiJsonResultPath) {
		return
	}
	file, err := os.Open(nucleiJsonResultPath)
	if err != nil {
		utils.CheckError("nucleiResult()", err)
	}
	jsonData, err := io.ReadAll(file)
	err = json.Unmarshal(jsonData, &nuclei)
	if err != nil {
		utils.CheckError("nucleiResult()", err)

	}
}
func xscanResult() {
	if !utils.FileExists(xscanJsonResultPath) {
		return
	}
	slice, err := utils.ReadTextFileToSlice(xscanJsonResultPath)
	if err != nil {
		utils.CheckError("xscanResult()", err)
	}
	for _, data := range slice {
		var payload models.Xscan
		json.Unmarshal([]byte(data), &payload)
		xscan = append(xscan, payload)
	}
}
func nucleiToXray() {
	for _, data := range nuclei {
		tmp := models.Xray{
			CreateTime: convertTimeToTimestamp(data.Timestamp),
			Detail: models.Detail{
				Addr: data.Host,
				Snapshot: [][]string{
					{data.Request},
					{data.Response},
				},
				Extra: models.Extra{
					Links: data.Info.Reference,
					Level: data.Info.Severity,
				},
			},
			Plugin: data.TemplateID,
			Target: models.Target{URL: data.Host},
		}
		key := fmt.Sprintf("nuclei-%v", data.Info.Severity)
		results[key] = append(results[key], tmp)
	}
}
func xscanToXray() {
	for _, data := range xscan {
		tmp := models.Xray{
			Detail: models.Detail{
				Snapshot: [][]string{
					{data.Req},
					{},
				},
				Payload: data.SuggestPayload,
			},
			Plugin: fmt.Sprintf("xscan-%v", data.Desc),
			Target: models.Target{URL: data.URL},
		}
		results["xscan"] = append(results["xscan"], tmp)
	}
}
