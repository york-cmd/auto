package portscan

import (
	"auto/collector/cdn"
	"auto/commands"
	"auto/utils"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"strings"
)

var (
	YesFirewallIPS []string // 存在防火墙的 IP 地址
	nmapCodes      []string // nmap 执行命令
)

func getIPlistFile() {
	utils.SliceWriter("tmp/portscan/iplist.txt", cdn.NoCDNIPs)
}

// firewallJudgment 过滤存在防火墙的 IP 地址
func firewallJudgment() {
	for ip, ports := range IPPortMap {
		if len(ports) > 50 {
			YesFirewallIPS = append(YesFirewallIPS, ip)
			delete(IPPortMap, ip)
		}
	}
}
func getNmapCode() {
	for ip, ports := range IPPortMap {
		portStr := strings.Join(ports, ",")
		code := fmt.Sprintf(commands.Config.Portscan.Nmap, ip, portStr, ip)
		nmapCodes = append(nmapCodes, code)
	}
}

// printResults 打印结果
func printResults() {
	data := [][]string{}
	for service, ipPorts := range serviceIPPorts {
		data = append(data, []string{service, strconv.Itoa(len(ipPorts))})
	}
	data = append(data, []string{"urls", strconv.Itoa(len(urls))})
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"状态", "数量"})
	table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	table.AppendBulk(data)
	table.Render()
}
