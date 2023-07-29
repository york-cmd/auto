package result

import (
	"auto/models"
	"auto/utils"
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
)

var (
	WebInfo = make(map[string]models.WebInfo)
	tmp     []models.JsonResult
)

func Save() {
	log.Println("start saving the information collection results ...")
	getJsonStructResult()
	saveToCSV()
	saveToJSON()
	log.Println("the result is saved .")
}

func getJsonStructResult() {
	for url, data := range WebInfo {
		tmp = append(tmp, models.JsonResult{
			Url:         url,
			Title:       data.Title,
			Fingerprint: data.Fingerprint,
			Waf:         data.Waf,
		})
	}
}

func saveToCSV() {
	filename := "result/webinfo/webinfo.csv"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("os.Create error : %v", err)
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(file)
	writer.Write([]string{"url", "title", "fingerprint", "waf"})
	for _, info := range tmp {
		writer.Write([]string{info.Url, info.Title, info.Fingerprint, info.Waf})
	}
	writer.Flush()
}

func saveToJSON() {
	filename := "result/webinfo/webinfo.json"
	file, err := os.Create(filename)
	if err != nil {
		utils.CheckError("saveToJSON()", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	encoder.Encode(tmp)
}
