package portscan

import (
	"auto/models"
	"auto/utils"
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	IPPortMap         = make(map[string][]string)
	IPPortServicesMap = make(map[string][]models.PortService)
)

// masscanResult masscan ip port 提取
func masscanResult() {
	path := "tmp/portscan/masscan.json"
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		utils.CheckError("masscanResult()", err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		utils.CheckError("masscanResult()", err)
	}
	// 替换逗号和右方括号的组合, 否则会报错
	jsonStr := string(data)
	jsonStr = strings.Replace(jsonStr, ",\n]", "]\n", -1)
	var results []models.MasscanResult
	err = json.Unmarshal([]byte(jsonStr), &results)
	if err != nil {
		utils.CheckError("masscanResult()", err)
	}
	for _, result := range results {
		IPPortMap[result.IP] = append(IPPortMap[result.IP], strconv.Itoa(result.Ports[0].Port))
	}
}

func nmapResult() {
	/***
	1. 遍历 tmp/portscan/nmap 获取 ip.xml 文件
	2. 遍历 解析 xml 文件, 提取 port - service
	*/
	files, err := utils.GetSuffixFiles("tmp/portscan/nmap", ".xml")
	if err != nil {
		utils.CheckError("nmapResult()", err)
	}
	for _, file := range files {
		getPortService(file)
	}
}
func getPortService(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		utils.CheckError("getPortService()", err)
	}
	defer file.Close()
	var tempData models.Nmap
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&tempData)
	if err != nil {
		utils.CheckError("getPortService()", err)
	}
	var PortServices []models.PortService
	for _, port := range tempData.Host.Ports.Port {
		PortServices = append(PortServices, models.PortService{Port: port.Portid, Service: port.Service.Name})
	}
	IPPortServicesMap[tempData.Host.Addr.IP] = PortServices
}
