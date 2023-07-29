package cdn

import (
	"auto/commands"
	"github.com/olekukonko/tablewriter"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
)

// 1. 返回 ip 列表进行端口扫描
// 2. 返回 ip 和 域名的对应关系
// 3. 返回 存在 cdn 的域名列表, 拼接常见 web 端口进行验活

var (
	NoCDNIPs      []string
	YesCDNDomains []string
	IPDomainsMap  = make(map[string][]string)
	wg            sync.WaitGroup
	rw            sync.RWMutex
)

// detectCDN cdn 识别
// 1. 获取域名 ipv4 的解析记录
// 2. ipv4 数量 = 1 => NoCDNIPs
// 3. ipv4 数量 > 1 domain => YesCDNDomains
func detectCDN(domain string) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return
	}
	// 获取 ipv4 切片
	var ipv4s []string
	for _, ip := range ips {
		if ip.To4() != nil {
			ipv4s = append(ipv4s, ip.String())
		}
	}
	if len(ipv4s) == 1 {
		rw.Lock()
		IPDomainsMap[ipv4s[0]] = append(IPDomainsMap[ipv4s[0]], domain)
		rw.Unlock()
		return
	}
	if len(ipv4s) > 1 {
		rw.Lock()
		YesCDNDomains = append(YesCDNDomains, domain)
		rw.Unlock()
	}
}

// printResults 打印结果
func printResults() {
	data := [][]string{
		{"yesCDN domains", strconv.Itoa(len(YesCDNDomains))},
		{"noCDN ips", strconv.Itoa(len(NoCDNIPs))},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"工具", "数量"})
	table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	table.AppendBulk(data)
	table.Render()
}
func Run(subdomains []string) {
	log.Println("start cdn identification ...")
	limit := make(chan struct{}, commands.Config.Goroutine.Cdn)
	for _, domain := range subdomains {
		wg.Add(1)
		limit <- struct{}{}
		go func(domain string) {
			detectCDN(domain)
			<-limit
			wg.Done()
		}(domain)
	}
	wg.Wait()
	for ip, _ := range IPDomainsMap {
		NoCDNIPs = append(NoCDNIPs, ip)
	}
	log.Println("cdn identification completed .")
	printResults()
}
