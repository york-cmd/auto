package urlvalid

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

// printResults 打印结果
func printResults(allUrl, okUrl int) {
	data := [][]string{
		{"active urls", strconv.Itoa(allUrl)},
		{"websites", strconv.Itoa(okUrl)},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"状态", "数量"})
	table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	table.AppendBulk(data)
	table.Render()
}
