package subdomain

import (
	"auto/utils"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"strconv"
)

// goSubRun goResult 协程运行
func goSubRun(f func(string), domain string) {
	defer wg.Done()
	f(domain)
}
func goResult(f func()) {
	defer wg.Done()
	f()
}

// subdomainRun 子域名收集工具运行
func subdomainRun(domain string) {
	wg.Add(3)
	go goSubRun(oneforallRun, domain)
	go goSubRun(subfinderRun, domain)
	go goSubRun(ksubEnumRun, domain)
	wg.Wait()
}

// GetResults 提取工具的子域名
func GetResults(alterxUse bool) {
	wg.Add(3)
	go goResult(oneResult)
	go goResult(subResult)
	go goResult(ksubResult)
	wg.Wait()
	numOneForAll, numSubfinder, numKsubdomain = len(oneforallResults), len(subfinderResults), len(ksubdomainResults)
	var tempResults []string
	tempResults = append(tempResults, append(oneforallResults, append(subfinderResults, ksubdomainResults...)...)...)
	tempResults = utils.Deduplication(tempResults)
	numThreeTools = len(tempResults)
	if alterxUse {
		utils.SliceWriter("tmp/subdomain/temp.txt", tempResults)
		alterxRun()
		altResult()
		numAlterx = len(alterxResults)
		tempResults = append(tempResults, alterxResults...)
		tempResults = utils.Deduplication(tempResults)
		numFourTools = len(tempResults)
		utils.SliceWriter("tmp/subdomain/all.txt", tempResults)
	} else {
		utils.SliceWriter("tmp/subdomain/all.txt", tempResults)
	}
	ksubVerifyRun()
	verResult()
	log.Println("collection of subdomain ends .")
	data := [][]string{
		{"oneforall", strconv.Itoa(numOneForAll)},
		{"subfinder", strconv.Itoa(numSubfinder)},
		{"ksubdomain enum", strconv.Itoa(numKsubdomain)},
		{"merge deduplication", strconv.Itoa(numThreeTools)},
	}
	if alterxUse {
		tmp := [][]string{
			{"alterx", strconv.Itoa(numAlterx)},
			{"merge deduplication", strconv.Itoa(numFourTools)},
		}
		data = append(data, tmp...)
	}
	data = append(data, []string{"ksubdomain verify", strconv.Itoa(len(Subdomains))})
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"工具", "数量"})
	table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	table.AppendBulk(data)
	table.Render()
}
