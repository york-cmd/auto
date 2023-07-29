package urlvalid

import (
	"auto/collector/result"
	"auto/models"
	"auto/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

var Urls []string

func httpxResult() {
	file, err := os.Open("result/httpx/urlInfo.json")
	if err != nil {
		utils.CheckError("httpxResult()", err)
	}
	defer file.Close()
	var dataList []models.HttpxResult
	decoder := json.NewDecoder(file)
	for decoder.More() {
		var data models.HttpxResult
		if err := decoder.Decode(&data); err != nil {
			log.Fatalf("%v", err)
		}
		dataList = append(dataList, data)
	}
	var tmpMap = make(map[string]models.HttpxResult)
	for _, data := range dataList {
		if data.StatusCode == 200 {
			key := fmt.Sprintf("%v-%v-%v", data.Title, data.Host, data.Hash.BodyMD5)
			if existingData, flag := tmpMap[key]; !flag {
				tmpMap[key] = data
			} else {
				// https 在 TideFinger 识别更准确
				if !strings.Contains(existingData.URL, "https") && strings.Contains(data.URL, "https") {
					tmpMap[key] = data
				}
			}
		}
	}

	for _, data := range tmpMap {
		result.WebInfo[data.URL] = models.WebInfo{
			Title: data.Title,
		}
	}
	for url, _ := range result.WebInfo {
		Urls = append(Urls, url)
	}
	utils.SliceWriter("result/httpx/urls.txt", Urls)
	printResults(len(dataList), len(result.WebInfo))
}
