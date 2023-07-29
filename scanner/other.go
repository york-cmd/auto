package scanner

import (
	"auto/utils"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"time"
)

var (
	temp             = `<script class='web-vulns'>webVulns.push(%s)</script>`
	templateFilePath = "template.html"
	websiteNum       = 1
)

func convertTimeToTimestamp(timeStr string) int64 {
	parsedTime, err := time.Parse(time.RFC3339Nano, timeStr)
	if err != nil {
		return 0
	}
	unixTimestampMs := parsedTime.UnixNano() / int64(time.Millisecond)
	return unixTimestampMs
}

// checkFileNotModified 当 xray.log 超过 1min 无变化时kill xray
func checkFileNotModified() {
	filePath := fmt.Sprintf("%s/tmp/vuln/xray.log", utils.Pwd())
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		utils.CheckError("checkFileNotModified()", err)
		return
	}
	defer watcher.Close()

	exit := make(chan struct{})

	go func() {
		var lastChangeTime time.Time
		durationThreshold := time.Minute
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					lastChangeTime = time.Now()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				utils.CheckError("checkFileNotModified()", err)
			case <-time.After(durationThreshold):
				if time.Since(lastChangeTime) >= durationThreshold {
					close(exit)
					return
				}
			}
		}
	}()
	err = watcher.Add(filePath)
	if err != nil {
		utils.CheckError("checkFileNotModified()", err)
	}
	<-exit
}
func printResults() {
	XrayResult()
	data := [][]string{
		{"nuclei", strconv.Itoa(len(nuclei))},
		{"xscan", strconv.Itoa(len(xscan))},
		{"xray", strconv.Itoa(xrayNum)},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"工具", "数量"})
	table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	table.AppendBulk(data)
	table.Render()
	tableData := fmt.Sprintf(`
|  工具   | 数量  |
|  ----  | ----  |
| website  | %v |
| nuclei  | %v |
| xscan  | %v |
| xray  | %v |
`, websiteNum, len(nuclei), len(xscan), xrayNum)
	utils.SendMessageToServerChan("漏扫结束, 请查看漏洞情报 !", tableData)
}
