package subdomain

import (
	"auto/utils"
	"log"
	"strings"
)

// 各工具结果切片 + 保存文件名
var (
	ksubdomainResults  []string
	ksubdomainFileName = "tmp/subdomain/ksubdomain.txt"
)
var (
	oneforallResults  []string
	oneforallFilePath = "tmp/subdomain/oneforall"
)
var (
	subfinderResults  []string
	subfinderFileName = "tmp/subdomain/subfinder.txt"
)
var (
	alterxResults  []string
	alterxFileName = "tmp/subdomain/alterx.txt"
)
var (
	Subdomains        []string
	subdomainFileName = "result/subdomains/subdomains.txt"
)

// oneResult oneforall 子域名提取
func oneResult() {
	files, err := utils.GetSuffixFiles(oneforallFilePath, ".csv")
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, filename := range files {
		oneforallResults = append(oneforallResults, utils.GetCsvColumn(filename, 6)...)
	}
}

// subResult subfinder 子域名提取
func subResult() {
	slice, err := utils.ReadTextFileToSlice(subfinderFileName)
	if err != nil {
		log.Fatalf("%v", err)
	}
	subfinderResults = append(subfinderResults, slice...)
}

// ksubResult ksubdomain 子域名提取
func ksubResult() {
	var ipDomains = make(map[string][]string)
	lines, err := utils.ReadTextFileToSlice(ksubdomainFileName)
	if err != nil {
		log.Fatalf("%v", err)
	}
	// 分割 => 提取 ip 和 子域名, 当单个 ip 的子域名数量 > 15 就不再接受此 ip 的子域名
	for _, line := range lines {
		parts := strings.Split(line, "=>")
		domain := parts[0]
		ip := parts[len(parts)-1]
		if len(ipDomains[ip]) < 15 {
			ipDomains[ip] = append(ipDomains[ip], domain)
			ksubdomainResults = append(ksubdomainResults, domain)
		}
	}
}

// altResult alterx 子域名提取
func altResult() {
	slice, err := utils.ReadTextFileToSlice(alterxFileName)
	if err != nil {
		utils.CheckError("altResult()", err)
	}
	alterxResults = append(alterxResults, slice...)
}

// verResult ksubdomain 验证 alterx 的子域名列表存活的子域名
func verResult() {
	slice, err := utils.ReadTextFileToSlice(subdomainFileName)
	if err != nil {
		utils.CheckError("verResult()", err)
	}
	Subdomains = append(Subdomains, slice...)
}
