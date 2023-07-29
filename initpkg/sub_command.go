package initpkg

import (
	"auto/collector"
	"auto/scanner"
	"auto/security"
	"flag"
	"fmt"
	"os"
)

// subCollectorCommand 信息收集子命令
func subCollectorCommand(args []string) {
	collectorCommand := flag.NewFlagSet("collector", flag.ExitOnError)
	collectorCommand.Usage = func() {
		fmt.Fprintln(os.Stderr, "collector Commands:")
		fmt.Fprintln(os.Stderr, "    --target <target>  单个目标")
		fmt.Fprintln(os.Stderr, "    --targets <file>   目标文件")
		fmt.Fprintln(os.Stderr, "    --alterx <boolean> 是否使用 alterx 生成子域名列表 ( 默认关闭 )")
		fmt.Fprintln(os.Stderr)
		collectorCommand.PrintDefaults()
	}
	if len(args[1:]) == 0 {
		collectorCommand.Usage()
		os.Exit(0)
	}
	var target, targetFile string
	var alterxUse bool
	collectorCommand.StringVar(&target, "target", "", "")
	collectorCommand.StringVar(&targetFile, "targets", "", "")
	collectorCommand.BoolVar(&alterxUse, "alterx", false, "")
	collectorCommand.Parse(args[1:])
	domains := getTargetSlice(target, targetFile)
	initCollectorDir()
	collector.Run(domains, alterxUse)
}

// subScannerCommand 漏洞扫描子命令
func subScannerCommand(args []string) {
	scannerCommand := flag.NewFlagSet("scanner", flag.ExitOnError)
	scannerCommand.Usage = func() {
		fmt.Fprintln(os.Stderr, "scanner Commands:")
		fmt.Fprintln(os.Stderr, "    --target <target>  单个目标")
		fmt.Fprintln(os.Stderr, "    --targets <file>   目标文件")
		fmt.Fprintln(os.Stderr)
		scannerCommand.PrintDefaults()
	}
	if len(args[1:]) == 0 {
		scannerCommand.Usage()
		os.Exit(0)
	}
	var target, targetFile string
	scannerCommand.StringVar(&target, "target", "", "")
	scannerCommand.StringVar(&targetFile, "targets", "", "")
	scannerCommand.Parse(args[1:])
	initScannerDir()
	if target == "" {
		scanner.MultipleRun(targetFile)
	} else {
		scanner.SingleRun(target)
	}
}

// subSecurityCommand 信息收集+漏洞扫描 一条龙
func subSecurityCommand(args []string) {
	securityCommand := flag.NewFlagSet("security", flag.ExitOnError)
	securityCommand.Usage = func() {
		fmt.Fprintln(os.Stderr, "security Commands:")
		fmt.Fprintln(os.Stderr, "    --target <target>  单个目标")
		fmt.Fprintln(os.Stderr, "    --targets <file>   目标文件")
		fmt.Fprintln(os.Stderr, "    --alterx <boolean> 是否使用 alterx 生成子域名列表 ( 默认关闭 )")
		fmt.Fprintln(os.Stderr)
		securityCommand.PrintDefaults()
	}
	if len(args[1:]) == 0 {
		securityCommand.Usage()
		os.Exit(0)
	}
	var target, targetFile string
	var alterxUse bool
	securityCommand.StringVar(&target, "target", "", "")
	securityCommand.StringVar(&targetFile, "targets", "", "")
	securityCommand.BoolVar(&alterxUse, "alterx", false, "")
	securityCommand.Parse(args[1:])
	domains := getTargetSlice(target, targetFile)
	initSecurityDir()
	security.Run(domains, alterxUse)
}
