package initpkg

import (
	"auto/utils"
	"flag"
	"fmt"
	"log"
	"os"
)

func initFlags() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Commands:")
		fmt.Fprintln(os.Stderr, "  collector   信息收集")
		fmt.Fprintln(os.Stderr, "  scanner     漏洞扫描")
		fmt.Fprintln(os.Stderr, "  security    一条龙")
		fmt.Fprintln(os.Stderr)
		flag.PrintDefaults()
	}

	// 解析命令行参数
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	command := args[0]

	switch command {
	case "collector":
		subCollectorCommand(args)
	case "scanner":
		subScannerCommand(args)
	case "security":
		subSecurityCommand(args)
	}
}

// getTargetSlice 从命令行输入中目标切片 ( 单个目标 / 文件路径 )
func getTargetSlice(target, targets string) []string {
	var domains []string
	if targets != "" {
		slice, err := utils.ReadTextFileToSlice(targets)
		if err != nil {
			log.Fatalf("%v", err)
		}
		domains = append(domains, slice...)
		return domains
	}
	domains = append(domains, target)
	return domains
}
